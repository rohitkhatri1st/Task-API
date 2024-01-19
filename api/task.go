package api

import (
	"encoding/json"
	"net/http"

	"github.com/rohitkhatri1st/Task-API/model"
)

func (a *API) healthCheck(w http.ResponseWriter, r *http.Request) {
	asd := model.Task{
		Name: "Hello",
	}
	abc, _ := json.Marshal(asd)
	w.WriteHeader(http.StatusOK)
	w.Write(abc)
}

func (a *API) taskError(w http.ResponseWriter, r *http.Request) {
	asd := model.Task{
		Name: "Hello",
	}
	abc, _ := json.Marshal(asd)
	w.WriteHeader(http.StatusBadRequest)

	w.Write(abc)
}
