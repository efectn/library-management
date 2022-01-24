package utils

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
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

func ParseBody(c *fiber.Ctx, body interface{}) error {
	if err := c.BodyParser(body); err != nil {
		return ReturnErrorMessage(c, err.Error())
	}

	return nil
}
