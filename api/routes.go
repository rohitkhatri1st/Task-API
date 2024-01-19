package api

// InitRoutes initializes all the endpoints
func (a *API) InitRoutes() {
	a.Router.APIRoot.Get("/task", a.healthCheck)
	a.Router.APIRoot.Get("/task-error", a.taskError)
}
