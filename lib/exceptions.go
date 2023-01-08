package lib

import "github.com/go-playground/validator/v10"

type BadRequest struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
}

type NotFound struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
}

type InternalServerError struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
}

type Unauthorized struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
}

type BadValidation struct {
	Message    string     `json:"message"`
	StatusCode int        `json:"status_code"`
	Errors     []ApiError `json:"errors"`
}

type ApiError struct {
	Param   string `json:"param"`
	Message string `json:"message"`
}

func NewBadValidation(errors []ApiError) BadValidation {
	return BadValidation{
		Message:    "Bad validation",
		StatusCode: 400,
		Errors:     errors,
	}
}

func MsgForTag(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	case "email":
		return "Invalid email"
	case "datetime":
		return "Invalid date"
	case "min":
		return "This field must be at least " + fe.Param() + " characters"
	case "max":
		return "This field must be at most " + fe.Param() + " characters"
	}

	return fe.Error() // default error
}
