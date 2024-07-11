package utils

import (
	"github.com/gofiber/fiber/v2"
)

type APIError struct {
	Message    string `json:"message"`
	StatusCode int    `json:"-"`
}

func (err APIError) Error() string {
	return err.Message
}

func ThrowError(context *fiber.Ctx, code int, message string) error {
	return context.Status(code).JSON(APIError{
		Message: message,
	})
}
