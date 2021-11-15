package utils

import (
	"github.com/go-playground/validator/v10"
)

type errorResponse struct {
	Name    string
	Tag     string
	Message string
}

func ValidateStruct(input interface{}) []*errorResponse {
	var errors []*errorResponse
	validate := validator.New()
	err := validate.Struct(input)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element errorResponse
			element.Name = err.Field()
			element.Tag = err.Tag()
			element.Message = err.Error()
			errors = append(errors, &element)
		}
	}

	return errors
}
