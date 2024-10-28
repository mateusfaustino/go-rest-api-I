package configurations

import "net/http"

// ErrorConfig represents the error object.
// @Summary Error information
// @Description Structure for describing why the error occurred
type ErrorConfig struct {
	// Error message.
	Message string `json:"message" example:"error trying to process request"`

	// Error description.
	Err string `json:"error" example:"internal_server_error"`

	// Error code.
	Code int `json:"code" example:"500"`

	// Error causes.
	Causes []Causes `json:"causes"`
}

// Causes represents the error causes.
// @Summary Error Causes
// @Description Structure representing the causes of an error.
type Causes struct {
	// Field associated with the error cause.
	// @json
	// @jsonTag field
	Field string `json:"field" example:"name"`

	// Error message describing the cause.
	// @json
	// @jsonTag message
	Message string `json:"message" example:"name is required"`
}

func (r *ErrorConfig) Error() string {
	return r.Message
}

func NewBadRequestError(message string) *ErrorConfig {
	return &ErrorConfig{
		Message: message,
		Err:     "bad_request",
		Code:    http.StatusBadRequest,
	}
}

func NewUnauthorizedRequestError(message string) *ErrorConfig {
	return &ErrorConfig{
		Message: message,
		Err:     "unauthorized",
		Code:    http.StatusUnauthorized,
	}
}

func NewBadRequestValidationError(message string, causes []Causes) *ErrorConfig {
	return &ErrorConfig{
		Message: message,
		Err:     "bad_request",
		Code:    http.StatusBadRequest,
		Causes:  causes,
	}
}

func NewInternalServerError(message string) *ErrorConfig {
	return &ErrorConfig{
		Message: message,
		Err:     "internal_server_error",
		Code:    http.StatusInternalServerError,
	}
}

func NewNotFoundError(message string) *ErrorConfig {
	return &ErrorConfig{
		Message: message,
		Err:     "not_found",
		Code:    http.StatusNotFound,
	}
}

func NewForbiddenError(message string) *ErrorConfig {
	return &ErrorConfig{
		Message: message,
		Err:     "forbidden",
		Code:    http.StatusForbidden,
	}
}