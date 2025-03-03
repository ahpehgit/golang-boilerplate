package validators

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

func AddTranslations(validate *validator.Validate, translation *ut.Translator) {
	_ = validate.RegisterTranslation("email-format", *translation, func(ut ut.Translator) error {
		return ut.Add("email-format", "{0} is not in the correct format", true) // see universal-translator for details
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("email-format", fe.Field())
		return t
	})

	_ = validate.RegisterTranslation("script-injection", *translation, func(ut ut.Translator) error {
		return ut.Add("script-injection", "{0} is not in a legitimate string format", true) // see universal-translator for details
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("script-injection", fe.Field())
		return t
	})

	//register more validation here if required
}

func TranslateErrors(err error, translation *ut.Translator) (errs []string) {
	if err == nil {
		return nil
	}

	validatorErrs := err.(validator.ValidationErrors)
	for _, e := range validatorErrs {
		translatedErr := e.Translate(*translation)
		errs = append(errs, translatedErr)
	}

	return
}
