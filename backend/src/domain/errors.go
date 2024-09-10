package domain

import (
	api "back/src/generated"
	"errors"
	"net/http"
)

var (
	ErrNoAffected = errors.New("no affected")

	ErrInternal = errors.New("error")
)

func getErrorCode(err error) int {
	switch err {
	default:
		return 0
	}
}

func getStatusCode(err error) int {
	switch err {
	default:
		return http.StatusInternalServerError
	}
}

func GetErrorResponse(err error) (int, api.Error) {
	return getStatusCode(err), api.Error{Code: int32(getErrorCode(err)), Message: err.Error()}
}
