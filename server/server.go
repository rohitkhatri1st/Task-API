/*
	The server package binds all the struct and interfaces of various aspects such as router, database, logging etc.
	StartServer and StopServer functions are exposed to call them via main package or via command line to start/stop
	the execution.

	The server itself listing on some address and port (localhost:8000 (default)) via go routine and will be blocked until
	StopServer function is called via some function or command line. Before stoping server all the resources and connections are closed.
*/

package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/rohitkhatri1st/Task-API/api"
	"github.com/rohitkhatri1st/Task-API/app"
	"github.com/rohitkhatri1st/Task-API/app/validator"
	"github.com/rohitkhatri1st/Task-API/server/config"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
	"gorm.io/gorm"
)

// Server object encapsulates api, business logic (app),router, storage layer and loggers
type Server struct {
	httpServer *http.Server
	Router     *chi.Mux
	Log        *zerolog.Logger
	WebhookLog *zerolog.Logger
	Config     *config.Config
	PSql       *gorm.DB
	API        *api.API
}

// NewServer returns a new Server object
func NewServer() *Server {
	c := config.GetConfig()
	r := chi.NewMux()
	server := &Server{
		httpServer: &http.Server{},
		Config:     c,
		Router:     r,
	}

	server.InitLoggers()
	appLogger := server.Log.With().Str("type", "app").Logger()
	apiLogger := server.Log.With().Str("type", "api").Logger()
	server.API = api.NewAPI(&api.Options{
		MainRouter: r,
		Logger:     &apiLogger,
		Validator:  validator.NewValidation(),
	})

	server.API.App = app.NewApp(&app.Options{Logger: &appLogger, Config: &c.APPConfig})

	app.InitService(server.API.App)

	return server
}

func (s *Server) InitLoggers() {
	cw := zerolog.ConsoleWriter{Out: os.Stdout}

	zlog := zerolog.New(cw).With().Timestamp().Stack().Caller().Logger()
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	s.Log = &zlog
}

func (s *Server) StartServer() {

	s.httpServer = &http.Server{
		Handler:      s.Router,
		Addr:         fmt.Sprintf("%s:%s", s.Config.ServerConfig.ListenAddr, s.Config.ServerConfig.Port),
		ReadTimeout:  s.Config.ServerConfig.ReadTimeout * time.Second,
		WriteTimeout: s.Config.ServerConfig.WriteTimeout * time.Second,
	}

	s.Log.Info().Msgf("Staring server at %s:%s", s.Config.ServerConfig.ListenAddr, s.Config.ServerConfig.Port)
	go func() {
		err := s.httpServer.ListenAndServe()
		if err != nil {
			s.Log.Error().Err(err).Msg("")
			return
		}
	}()
}

func (s *Server) StopServer() {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	s.Log.Debug().Msg("Shutting Down Server")
	s.httpServer.Shutdown(ctx)
	s.Log.Debug().Msg("HTTP Server Shut Down")
}
