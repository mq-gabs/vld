package validator

import (
	"errors"

	"github.com/mq-gabs/vld/internal/utils"
)

var (
	ErrValueIsRequired = errors.New("value is required")
	ErrMinLen          = errors.New("required min length")
	ErrMaxLen          = errors.New("required max length")
)

// When applies validate only when condition is true.
func When[T any](condition bool, validate func(T) error) func(T) error {
	if condition {
		return validate
	}

	return func(_ T) error { return nil }
}

// WhenFunc applies validate only when condition returns true for the value.
func WhenFunc[T any](condition func(T) bool, validate func(T) error) func(T) error {
	return func(t T) error {
		if !condition(t) {
			return nil
		}

		return validate(t)
	}
}

// OneOf succeeds when at least one of the validators succeeds and otherwise
// returns the errors from all validators.
func OneOf[T any](validators ...func(T) error) func(T) error {
	return func(t T) error {
		errGroup := utils.NewErrorGroup()

		for _, validate := range validators {
			err := validate(t)
			if err == nil {
				return nil
			}

			errGroup.Join(err)
		}

		if !errGroup.IsNil() {
			return errGroup
		}

		return nil
	}
}
