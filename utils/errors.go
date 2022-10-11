package utils

import (
	"errors"
	"net/http"
)

var ErrResourceNotFound = errors.New("Cannot Find Specified Resource")

func FindError(w http.ResponseWriter, err error, statusCode int) {
	w.WriteHeader(statusCode)
	w.Write([]byte(err.Error()))
}
