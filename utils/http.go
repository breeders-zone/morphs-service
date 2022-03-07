package utils

import (
	"fmt"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/manyminds/api2go"
)

func Validate(input interface{}) []api2go.Error {
	var valErrors []api2go.Error
	if err := validator.New().Struct(input); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element api2go.Error
			element.Detail = "err"
			element.Title = "Unprocessable Entity"
			element.Status = fmt.Sprint(fiber.StatusUnprocessableEntity)
			element.Source = &api2go.ErrorSource{Pointer: "/data/" + err.Field()}
			valErrors = append(valErrors, element)
		}
	}

	return valErrors
}