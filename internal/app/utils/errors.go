package utils

import (
	"errors"
	"net/http"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

type HTTPError struct {
	Code    int
	Message string
	Err     error
}

func (e HTTPError) Error() string {
	return e.Message
}

func NewHTTPError(code int, msg string, err error) HTTPError {
	return HTTPError{
		Code:    code,
		Message: msg,
		Err:     err,
	}
}

func BadRequest(msg string) HTTPError {
	return NewHTTPError(http.StatusBadRequest, msg, errors.New(msg))
}

func NotFound(msg string) HTTPError {
	return NewHTTPError(http.StatusNotFound, msg, errors.New(msg))
}
