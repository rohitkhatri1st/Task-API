package app

import (
	"encoding/json"
	"net/http"

	"github.com/rohitkhatri1st/Task-API/model"
	"github.com/rs/zerolog"
)

type Utility interface {
	CreateRequestObject(request http.Request, requestBodyBytes []byte) *model.SentRequest
	CreateResponseObject(response http.Response, responseBodyBytes []byte) *model.ReceivedResponse
	Logs(request http.Request, requestBody interface{}, location string, errs ...error)
}

type UtilityImpl struct {
	App    *App
	Logger *zerolog.Logger
}
type UtilityOpts struct {
	App    *App
	Logger *zerolog.Logger
}

func InitUtility(opts *UtilityOpts) Utility {
	ut := UtilityImpl{
		App:    opts.App,
		Logger: opts.Logger,
	}
	return &ut
}

func (ui *UtilityImpl) CreateResponseObject(response http.Response, responseBodyBytes []byte) *model.ReceivedResponse {
	var body interface{}
	json.Unmarshal(responseBodyBytes, &body)
	responseObject := model.ReceivedResponse{
		Status:        response.Status,
		StatusCode:    response.StatusCode,
		Proto:         response.Proto,
		ProtoMajor:    response.ProtoMajor,
		ProtoMinor:    response.ProtoMinor,
		Header:        response.Header,
		Body:          body,
		ContentLength: response.ContentLength,
	}
	return &responseObject
}

func (ui *UtilityImpl) CreateRequestObject(request http.Request, requestBodyBytes []byte) *model.SentRequest {
	var body interface{}
	json.Unmarshal(requestBodyBytes, &body)
	header := request.Header.Clone()
	delete(header, "Authorization")
	requestObject := model.SentRequest{
		Method:        request.Method,
		URL:           request.URL,
		Proto:         request.Proto,
		ProtoMajor:    request.ProtoMajor,
		ProtoMinor:    request.ProtoMinor,
		Header:        header,
		Body:          body,
		ContentLength: request.ContentLength,
		Form:          request.Form,
		PostForm:      request.PostForm,
	}
	return &requestObject
}

func (ui *UtilityImpl) Logs(request http.Request, requestBody interface{}, location string, errs ...error) {
	requestBodyBytes, err := json.Marshal(requestBody)
	if err != nil {
		errs = append(errs, err)
	}
	reqObject := ui.CreateRequestObject(request, requestBodyBytes)
	if len(errs) > 0 {
		var err error
		for _, v := range errs {
			if v != nil {
				err = v
			}
		}
		if err != nil {
			ui.Logger.Err(err).
				Interface("request", reqObject).
				Errs("errors", errs).
				Interface("location", location).Msg("An error occurred")
		}
	}
}
