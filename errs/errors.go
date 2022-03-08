package errs

import "net/http"

type ApplicationError struct {
	Code    int    `json:",omitempty"`
	Message string `json:"message"`
}

func (e ApplicationError) AsMessage() *ApplicationError {
	return &ApplicationError{
		Message: e.Message,
	}
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

func NewValidationError(message string) *ApplicationError {
	return &ApplicationError{
		Message: message,
		Code:    http.StatusUnprocessableEntity, // 422
	}
}
