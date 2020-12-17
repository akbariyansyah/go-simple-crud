package utils

import (
	"encoding/json"
	"go-simple-crud/model"
	"log"
	"net/http"
)

func SendErrorResponse(msg string, code int, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	byteData, err := json.Marshal(model.Response{
		Message: msg,
		Data:    nil,
	})
	if err != nil {
		log.Panicln(err.Error())
	}
	w.Write(byteData)
}
func SendSuccessResponse(msg string, code int, data interface{}, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	byteData, err := json.Marshal(model.Response{
		Message: msg,
		Data:    data,
	})
	if err != nil {
		log.Panicln(err.Error())
	}
	w.Write(byteData)
}
