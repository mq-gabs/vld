package validator

import "errors"

var (
	ErrValueIsRequired = errors.New("value is required")
	ErrMinLen          = errors.New("required min length")
	ErrMaxLen          = errors.New("required max length")
)

type GenericValidator[T any] func(T) error
