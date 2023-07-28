package utils

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func ValidateBody(obj interface{}) error {
	validate := validator.New()

	// Validate the struct
	err := validate.Struct(obj)
	if err != nil {
		// Validation error occurred
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Printf("%s is required\n", err.Field())
		}
		return err
	}
	return nil
}
