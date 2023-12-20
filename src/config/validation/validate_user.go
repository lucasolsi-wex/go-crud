package validation

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en_US"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en2 "github.com/go-playground/validator/v10/translations/en"
	"github.com/lucasolsi-wex/go-crud/src/config/custom_errors"
	"github.com/lucasolsi-wex/go-crud/src/model/request"
)

var (
	Validate         = validator.New()
	error_translator ut.Translator
)

func init() {
	if val, ok := binding.Validator.Engine().(*validator.Validate); ok {
		en := en_US.New()
		un := ut.New(en, en)
		error_translator, _ = un.GetTranslator("en")
		en2.RegisterDefaultTranslations(val, error_translator)
	}
}

func ValidateUserError(
	validationError error) *custom_errors.CustomErr {
	var jsonError *json.UnmarshalTypeError
	var jsonValidationError validator.ValidationErrors

	if errors.As(validationError, &jsonError) {
		return custom_errors.NewBadRequestError("Invalid type")
	} else if errors.As(validationError, &jsonValidationError) {
		errorCauses := []custom_errors.Causes{}
		for _, e := range validationError.(validator.ValidationErrors) {
			cause := custom_errors.Causes{
				Message: e.Translate(error_translator),
				Field:   e.Field(),
			}
			errorCauses = append(errorCauses, cause)
		}

		return custom_errors.NewUserValidationFieldsError("Invalid fields!",
			errorCauses)
	} else {
		return custom_errors.NewBadRequestError("Error while converting fields")
	}
}

func ValidateFirstAndLastName(request request.UserRequest) *custom_errors.CustomErr {
	if len(request.FirstName) == 0 || len(request.LastName) == 0 {
		errorCauses := []custom_errors.Causes{}
		cause := custom_errors.Causes{
			Field:   "firstName/lastName",
			Message: "User first/last names required",
		}
		errorCauses = append(errorCauses, cause)
		return custom_errors.NewUserValidationFieldsError("Invalid fields", errorCauses)
	}
	return nil
}

func ValidateNameUniqueness(alreadyExists bool) *custom_errors.CustomErr {
	if alreadyExists {
		errorCauses := []custom_errors.Causes{}
		cause := custom_errors.Causes{
			Field:   "firstName/lastName",
			Message: "User with the same first and last name already exists",
		}
		errorCauses = append(errorCauses, cause)
		return custom_errors.NewUserValidationFieldsError("Invalid fields", errorCauses)
	}
	return nil
}
