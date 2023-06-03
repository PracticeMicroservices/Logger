package controllers

import (
	"logger/cmd/api/helpers"
	"logger/data/models"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
)

type Logger interface {
	WriteLog(w http.ResponseWriter, r *http.Request)
}

type loggerController struct {
	logModels models.Models
	json      *helpers.JsonResponse
}

type JSONPayload struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

func NewLoggerController(db *mongo.Client) Logger {
	return &loggerController{
		logModels: models.New(db),
		json:      &helpers.JsonResponse{},
	}
}

func (l *loggerController) WriteLog(w http.ResponseWriter, r *http.Request) {
	//read json
	var requestPayload JSONPayload
	_ = l.json.ReadJSON(w, r, &requestPayload)

	//insert data
	event := models.LogEntry{
		Name: requestPayload.Name,
		Data: requestPayload.Data,
	}

	err := l.logModels.LogEntry.Insert(event)
	if err != nil {
		_ = l.json.WriteJSONError(w, err)
		return
	}
	resp := &helpers.JsonResponse{
		Error:   false,
		Message: "logged",
	}
	_ = resp.WriteJSON(w, http.StatusOK, nil)
}
