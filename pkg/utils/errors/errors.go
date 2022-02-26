package errors

import (
	"fmt"
	"strings"

	"github.com/efectn/library-management/pkg/database/ent"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

// Error represents an error that occurred while handling a request.
type Error struct {
	Code    int         `json:"code"`
	Message interface{} `json:"message"`
}

// Error makes it compatible with the `error` interface.
func (e *Error) Error() string {
	return fmt.Sprint(e.Message)
}

// NewErrors creates multiple new Error messages
func NewErrors(code int, messages ...interface{}) *Error {
	e := &Error{
		Code:    code,
		Message: utils.StatusMessage(code),
	}
	if len(messages) > 0 {
		e.Message = messages
	}
	return e
}

// HandleEntError is a method to handle Ent's errors.
func HandleEntErrors(err error) error {
	// Check not found error
	if ent.IsNotFound(err) {
		return NewErrors(fiber.StatusNotFound, "The field not found in the database. Please check your field!")
	}

	// Check constraint errors
	if ent.IsConstraintError(err) {
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			return NewErrors(fiber.StatusForbidden, "The unique field has created before. Please check your fields!")
		} else if strings.Contains(err.Error(), "add m2m edge for table") {
			return NewErrors(fiber.StatusInternalServerError, "Relations not found or incorrect. Please check relations!")
		}

		return NewErrors(fiber.StatusInternalServerError, err.Error())
	}

	// Check other errors
	if err != nil {
		return NewErrors(fiber.StatusInternalServerError, "An un-handled error occurred!", err)
	}

	return nil
}
