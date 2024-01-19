package api

import (
	"encoding/json"
	"net/http"

	"github.com/rohitkhatri1st/Task-API/model"
)

// InitRoutes initializes all the endpoints
func (a *API) InitRoutes() {
	a.Router.APIRoot.Get("/health-check", a.healthCheck)

	a.Router.APIRoot.Post("/task", a.createTask)
	a.Router.APIRoot.Put("/task", a.healthCheck)
	a.Router.APIRoot.Get("/tasks", a.healthCheck)
	a.Router.APIRoot.Get("/task/{taskId}", a.healthCheck)
	a.Router.APIRoot.Delete("/task", a.healthCheck)

	a.Router.APIRoot.Get("/task-error", a.taskError)
}

func (a *API) healthCheck(w http.ResponseWriter, r *http.Request) {
	asd := model.Task{
		Title: "Hello",
	}
	abc, _ := json.Marshal(asd)
	w.WriteHeader(http.StatusOK)
	w.Write(abc)
}
