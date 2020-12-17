package utils

import (
	"net/http"

	"github.com/gorilla/mux"
)

func DecodeParams(r *http.Request, path string) string {
	params := mux.Vars(r)
	return params[path]
}
