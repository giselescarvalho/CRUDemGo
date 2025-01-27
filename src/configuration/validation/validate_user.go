package validation

import (
	"encoding/json"
	"errors"

	"github.com/gin-gonic/gin/binding"
	"github.com/giselescarvalho/CRUDemGo/src/configuration/apperror"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translation "github.com/go-playground/validator/v10/translations/en"
)

var (
	Validate = validator.New()
	transl   ut.Translator
)

func init() {
	// Casting
	if val, ok := binding.Validator.Engine().(*validator.Validate); ok {
		en := en.New()
		unt := ut.New(en, en)
		transl, _ = unt.GetTranslator("en")

		if err := en_translation.RegisterDefaultTranslations(val, transl); err != nil {
			panic("Failed to register translations: " + err.Error())
		}
	}
}

func ValidateUserError(validationErr error) *apperror.RestErr {
	var jsonErr *json.UnmarshalTypeError
	var jsonValidationError validator.ValidationErrors

	if errors.As(validationErr, &jsonErr) {
		return apperror.NewBadRequestError("Invalid field type")
	} else if errors.As(validationErr, &jsonValidationError) {
		errorsCauses := []apperror.Causes{}

		for _, e := range validationErr.(validator.ValidationErrors) {
			cause := apperror.Causes{
				Message: e.Translate(transl),
				Field:   e.Field(),
			}

			errorsCauses = append(errorsCauses, cause)
		}
		return apperror.NewBadRequestValidationError("Some fields are invalid", errorsCauses)
	} else {
		return apperror.NewBadRequestError("Error trying to convert fields")
	}
}
