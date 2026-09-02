package validator

import "errors"

var (
	ErrValueIsRequired = errors.New("value is required")
	ErrMinLen          = errors.New("required min length")
	ErrMaxLen          = errors.New("required max length")
)

func When[T any](condition bool, validate func(T) error) func(T) error {
	if condition {
		return validate
	}

	return func(t T) error { return nil }
}
