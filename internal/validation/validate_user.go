package validation

import (
	"encoding/json"
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/lucasolsi-wex/go-crud/internal/models"
)

func ValidateUserError(
	validationError error) *models.CustomErr {
	var jsonError *json.UnmarshalTypeError
	var jsonValidationError validator.ValidationErrors

	if errors.As(validationError, &jsonError) {
		return models.NewBadRequestError("Invalid type")
	} else if errors.As(validationError, &jsonValidationError) {
		errorCauses := []models.Causes{}
		for _, e := range validationError.(validator.ValidationErrors) {
			cause := models.Causes{
				Message: e.Error(),
				Field:   e.Field(),
			}
			errorCauses = append(errorCauses, cause)
		}

		return models.NewUserValidationFieldsError("Invalid fields!",
			errorCauses)
	} else {
		return models.NewBadRequestError("Error while converting fields")
	}
}

func ValidateFirstAndLastName(request models.UserRequest) *models.CustomErr {
	if len(request.FirstName) == 0 || len(request.LastName) == 0 {
		errorCauses := []models.Causes{}
		cause := models.Causes{
			Field:   "firstName/lastName",
			Message: "User first/last names required",
		}
		errorCauses = append(errorCauses, cause)
		return models.NewUserValidationFieldsError("Invalid fields", errorCauses)
	}
	return nil
}

func ValidateNameUniqueness(alreadyExists bool) *models.CustomErr {
	if alreadyExists {
		errorCauses := []models.Causes{}
		cause := models.Causes{
			Field:   "firstName/lastName",
			Message: "User with the same first and last name already exists",
		}
		errorCauses = append(errorCauses, cause)
		return models.NewUserValidationFieldsError("Invalid fields", errorCauses)
	}
	return nil
}
