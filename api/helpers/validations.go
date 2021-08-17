package helpers

import (
	dtos2 "github.com/cocolabo/go-gin-gorm/api/dtos"
	langs2 "github.com/cocolabo/go-gin-gorm/api/langs"
	"github.com/go-playground/validator/v10"
)

func GenerateValidationResponse(err error) (response dtos2.ValidationResponse) {
	response.Success = false

	var validations []dtos2.Validation

	validationErrors := err.(validator.ValidationErrors)

	for _, value := range validationErrors {
		field, rule := value.Field(), value.Tag()

		validation := dtos2.Validation{Field: field, Message: langs2.GenerateValidationMessage(field, rule)}

		validations = append(validations, validation)
	}

	response.Validations = validations

	return response
}
