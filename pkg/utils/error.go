package utils

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2/log"
)

type ErrorResponse struct {
	Error       bool
	FailedField string
	Tag         string
	Value       interface{}
}

type XValidator struct {
}

type GlobalErrorHandlerResp struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func (v XValidator) Validate(data interface{}) []ErrorResponse {
	var validationErrors []ErrorResponse
	errs := validator.New().Struct(data)
	if errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			var elem ErrorResponse

			log.Info(err)
			elem.FailedField = err.Field()
			elem.Tag = err.Error()
			elem.Value = err.Value()
			elem.Error = true

			validationErrors = append(validationErrors, elem)
		}
	}

	return validationErrors
}
