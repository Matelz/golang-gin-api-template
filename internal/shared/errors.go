package shared

import "fmt"

type ValidationError struct {
	Field   string
	Message string
}

func (v ValidationError) Error() string {
	return fmt.Sprintf("field '%s' %s", v.Field, v.Message)
}

func NewValidationError(field, message string) ValidationError {
	return ValidationError{
		Field:   field,
		Message: message,
	}
}

func ErrFieldRequired(field string) ValidationError {
	return NewValidationError(field, "is required")
}

func ErrInvalidField(field string) ValidationError {
	return NewValidationError(field, "must be a valid phone number")
}
