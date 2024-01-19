package api

import (
	"net/http"
)

// InitRoutes initializes all the endpoints
func (a *API) InitRoutes() {
	a.Router.APIRoot.Get("/health-check", a.healthCheck)

	a.Router.APIRoot.Post("/task", a.createTask)
	a.Router.APIRoot.Put("/task", a.updateTask)
	a.Router.APIRoot.Get("/tasks", a.getAllTasks)
	a.Router.APIRoot.Get("/task/{taskId}", a.getTaskById)
	a.Router.APIRoot.Delete("/task/{taskId}", a.deleteTask)

}

func (a *API) healthCheck(w http.ResponseWriter, r *http.Request) {
	a.SendJsonResponse(w, http.StatusOK, true)
}
