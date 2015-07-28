package api

import (
	"net/http"
)

type apiError struct {
	code int
	msg  string
}

var (
	invalidReqErr      = newApiError(http.StatusBadRequest, "invalid request")
	undefinedMethodErr = newApiError(http.StatusMethodNotAllowed, "undefined method")
	listNotFoundErr    = newApiError(http.StatusNotFound, "list not found")
)

func newApiError(code int, msg string) *apiError {
	return &apiError{code, msg}
}

func (e *apiError) Error() string {
	return e.msg
}
