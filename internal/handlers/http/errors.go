package http

import (
	"fmt"

	"github.com/danielalmeidafarias/go-clean/internal/errors"
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber"
)

var HttpStatusCode map[errors.ErrorStatus]int = map[errors.ErrorStatus]int{
	errors.BadRequest:         400,
	errors.NotFound:           404,
	errors.Conflict:           409,
	errors.FailedPrecondition: 412,
	errors.Internal:           500,
}

func invalidRequestBody(c *fiber.Ctx) {
	c.Status(400).JSON(fiber.Map{"error": "invalid request body"})
}

func validationError(c *fiber.Ctx, err error) error {
	validationErrors := make(map[string]string)

	if errs, ok := err.(validator.ValidationErrors); ok {
		for _, fieldErr := range errs {
			field := fieldErr.Field()
			tag := fieldErr.Tag()

			validationErrors[field] = fmt.Sprintf("field %s validation failed: %s", field, tag)
		}
	} else {
		validationErrors["error"] = "invalid input"
	}

	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"errors": validationErrors})
}

func useCaseError(c *fiber.Ctx, err *errors.Error) {
	c.Status(HttpStatusCode[err.Code]).JSON(fiber.Map{"error": err.Message})
}
