package models

import "net/http"

type CustomErr struct {
	Message string   `json:"message,omitempty"`
	Err     string   `json:"error,omitempty"`
	Causes  []Causes `json:"causes,omitempty"`
	Code    int      `json:"-"`
}

type Causes struct {
	Field   string `json:"field,omitempty"`
	Message string `json:"message,omitempty"`
}

func (r *CustomErr) Error() string {
	return r.Message
}

func NewBadRequestError(message string) *CustomErr {
	return &CustomErr{
		Message: message,
		Err:     "bad_rest",
		Code:    http.StatusBadRequest,
	}
}

func NewUserValidationFieldsError(message string, causes []Causes) *CustomErr {
	return &CustomErr{
		Message: message,
		Err:     "validation_error",
		Causes:  causes,
		Code:    http.StatusBadRequest,
	}
}

func NewInternalServerError(message string) *CustomErr {
	return &CustomErr{
		Message: message,
		Err:     "internal_server_error",
		Code:    http.StatusInternalServerError,
	}
}

func NewUserNotFoundError(message string) *CustomErr {
	return &CustomErr{
		Message: message,
		Err:     "not_found",
		Code:    http.StatusNotFound,
	}
}
