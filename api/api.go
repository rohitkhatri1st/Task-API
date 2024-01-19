package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/rohitkhatri1st/Task-API/app"
	"github.com/rohitkhatri1st/Task-API/app/validator"
	"github.com/rs/zerolog"
)

type API struct {
	MainRouter *chi.Mux
	Router     *Router
	Logger     *zerolog.Logger
	Validator  *validator.Validator

	App *app.App
}

type Router struct {
	Root       *chi.Mux
	APIRoot    *chi.Mux
	StaticRoot *chi.Mux
	TestRoot   *chi.Mux
}

type Options struct {
	MainRouter *chi.Mux
	Logger     *zerolog.Logger
	Validator  *validator.Validator
}

func NewAPI(opts *Options) *API {
	api := API{
		MainRouter: opts.MainRouter,
		Router:     &Router{},
		Logger:     opts.Logger,
		Validator:  opts.Validator,
	}

	api.setupRoutes()
	return &api
}

func (a *API) setupRoutes() {
	a.Router.Root = a.MainRouter

	a.Router.APIRoot = chi.NewMux()
	a.Router.Root.Mount("/api", a.Router.APIRoot)
	// Declare More types of routes if needed
	a.InitRoutes()
}
