package app

import (
	"net/http"

	"github.com/rohitkhatri1st/Task-API/server/config"
	"github.com/rs/zerolog"
)

type App struct {
	Logger *zerolog.Logger
	Config *config.APPConfig

	// List of services this app is implementing
	Task    Task
	Utility Utility
}

// Options contains arguments required to create a new app instance
type Options struct {
	Logger     *zerolog.Logger
	Config     *config.APPConfig
	HttpClient http.Client
}

func NewApp(opts *Options) *App {
	return &App{
		Logger: opts.Logger,
		Config: opts.Config,
	}
}
