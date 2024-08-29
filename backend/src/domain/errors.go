package domain

import (
	"errors"
	"net/http"
)

var (
	ErrorNotFound     = errors.New("not found resource")
	ErrorFailedInsert = errors.New("failed insert")
)

func GetStatusCode(err error) int {
	switch err {
	case ErrorNotFound:
		return http.StatusNotFound
	case ErrorFailedInsert:
		return http.StatusInternalServerError
	default:
		return http.StatusInternalServerError
	}
}
