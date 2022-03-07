package errs

import "net/http"

type ApplicationError struct {
	Code    int
	Message string
}

func NewNotFoundError(message string) *ApplicationError {
	return &ApplicationError{
		Message: message,
		Code:    http.StatusNotFound,
	}
}

func NewUnexpectedError(message string) *ApplicationError {
	return &ApplicationError{
		Message: message,
		Code:    http.StatusInternalServerError,
	}
}
