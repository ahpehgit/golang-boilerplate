package validators

import (
	"net/mail"

	"github.com/go-playground/validator/v10"
)

func ValidateEmailAddressFormat(field validator.FieldLevel) bool {
	// eg. 12345678-1234-5678--ABCD-ABCDEFGHIJKL
	if field.Field().String() == "" {
		return true
	}

	_, err := mail.ParseAddress(field.Field().String())
	return err == nil
}
