package interfaces

import "net/http"

// Error
type Error struct {
	Err     string `json:"err"`
	Message string `json:"message"`
	Code    int    `json:"code"`
}

// NewBadRequestError
func NewBadRequestError(message string) *Error {
	return &Error{
		Err:     "Bad Request",
		Message: message,
		Code:    http.StatusBadRequest,
	}
}

// NewInternalServerError
func NewInternalServerError(message string) *Error {
	return &Error{
		Err:     "Internal service error",
		Message: message,
		Code:    http.StatusInternalServerError,
	}
}

// NewNotFoundError
func NewNotFoundError(message string) *Error {
	return &Error{
		Err:     "Resource not found",
		Message: message,
		Code:    http.StatusNotFound,
	}
}
