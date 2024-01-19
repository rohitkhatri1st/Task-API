package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/rohitkhatri1st/Task-API/model"
)

func (a *API) taskError(w http.ResponseWriter, r *http.Request) {
	asd := model.Task{
		Name: "Hello",
	}
	abc, _ := json.Marshal(asd)
	w.WriteHeader(http.StatusBadRequest)

	w.Write(abc)
}

func (a *API) createTask(w http.ResponseWriter, r *http.Request) {
	var task model.Task
	if err := a.DecodeJSONBody(r, &task); err != nil {
		fmt.Println(err)
	}
	fmt.Println(task)
	a.App.Task.CreateTask(&task)
	w.WriteHeader(http.StatusBadRequest)

	w.Write([]byte("asd"))
}
