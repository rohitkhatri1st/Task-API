package api

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/rohitkhatri1st/Task-API/model"
)

func (a *API) createTask(w http.ResponseWriter, r *http.Request) {
	var task *model.Task
	if err := a.DecodeJSONBody(r, &task); err != nil {
		a.SendJsonResponse(w, http.StatusBadRequest, nil, err)
		return
	}
	errorsList := []error{}
	defer func() {
		a.App.Utility.Logs(*r, task, "createTask", errorsList...)
	}()

	if errs := a.Validator.Validate(task); errs != nil {
		errorsList = append(errorsList, errs...)
		a.SendJsonResponse(w, http.StatusInternalServerError, nil, errs...)
		return
	}

	task.ID = 0
	task, err := a.App.Task.CreateTask(task)
	if err != nil {
		errorsList = append(errorsList, err)
		a.SendJsonResponse(w, http.StatusInternalServerError, nil, err)
		return
	}
	a.SendJsonResponse(w, http.StatusCreated, task)
}

func (a *API) updateTask(w http.ResponseWriter, r *http.Request) {
	var task *model.Task
	if err := a.DecodeJSONBody(r, &task); err != nil {
		a.SendJsonResponse(w, http.StatusBadRequest, nil, err)
		return
	}
	errorsList := []error{}
	defer func() {
		a.App.Utility.Logs(*r, task, "createTask", errorsList...)
	}()

	if errs := a.Validator.Validate(task); errs != nil {
		errorsList = append(errorsList, errs...)
		a.SendJsonResponse(w, http.StatusBadRequest, nil, errs...)
		return
	}

	if task.ID == 0 {
		err := errors.New("id is a required field")
		errorsList = append(errorsList, err)
		a.SendJsonResponse(w, http.StatusBadRequest, nil, err)
		return
	}
	task, err := a.App.Task.UpdateTask(task)
	if err != nil {
		errorsList = append(errorsList, err)
		a.SendJsonResponse(w, http.StatusInternalServerError, nil, err)
		return
	}
	a.SendJsonResponse(w, http.StatusOK, task)
}

func (a *API) getAllTasks(w http.ResponseWriter, r *http.Request) {
	errorsList := []error{}
	defer func() {
		a.App.Utility.Logs(*r, nil, "createTask", errorsList...)
	}()

	page := a.GetPageValue(r)
	tasks, err := a.App.Task.GetAllTasks(uint(page))
	if err != nil {
		errorsList = append(errorsList, err)
		a.SendJsonResponse(w, http.StatusBadRequest, nil, err)
		return
	}
	a.SendJsonResponse(w, http.StatusOK, tasks)

}

func (a *API) getTaskById(w http.ResponseWriter, r *http.Request) {
	errorsList := []error{}
	defer func() {
		a.App.Utility.Logs(*r, nil, "createTask", errorsList...)
	}()

	taskIdString := chi.URLParam(r, "taskId")
	taskId, err := strconv.Atoi(taskIdString)
	if err != nil {
		errorsList = append(errorsList, err)
		a.SendJsonResponse(w, http.StatusBadRequest, nil, err)
		return
	}
	task, err := a.App.Task.GetTask(uint(taskId))
	if err != nil {
		errorsList = append(errorsList, err)
		a.SendJsonResponse(w, http.StatusBadRequest, nil, err)
		return
	}
	a.SendJsonResponse(w, http.StatusOK, task)
}

func (a *API) deleteTask(w http.ResponseWriter, r *http.Request) {
	errorsList := []error{}
	defer func() {
		a.App.Utility.Logs(*r, nil, "createTask", errorsList...)
	}()

	taskIdString := chi.URLParam(r, "taskId")
	taskId, err := strconv.Atoi(taskIdString)
	if err != nil {
		errorsList = append(errorsList, err)
		a.SendJsonResponse(w, http.StatusBadRequest, nil, err)
		return
	}
	err = a.App.Task.DeleteTask(uint(taskId))
	if err != nil {
		errorsList = append(errorsList, err)
		a.SendJsonResponse(w, http.StatusBadRequest, nil, err)
		return
	}
	a.SendJsonResponse(w, http.StatusOK, true)
}
