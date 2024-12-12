package utils

import "github.com/go-playground/validator/v10"

// Response defines the structure of API response.
type Response struct {
	Status  bool        `json:"status"`  // Status indicates the success or failure of the API call.
	Message string      `json:"message"` // Message provides a descriptive message about the API response.
	Data    interface{} `json:"data"`    // Data contains the payload of the API response.
}

// APIResponse constructs and returns an API response.
func APIResponse(status bool, message string, data interface{}) Response {
	return Response{
		Status:  status,
		Message: message,
		Data:    data,
	}
}

// FormatValidationError formats validation errors into a slice of strings.
func FormatValidationError(err error) []string {
	var errors []string

	for _, e := range err.(validator.ValidationErrors) {
		errors = append(errors, e.Error())
	}

	return errors
}
