package vnumber

import (
	"errors"
	"fmt"
	"slices"
)

var (
	ErrMinValue        = errors.New("value must be greater than or equal to minimum")
	ErrMaxValue        = errors.New("value must be less than or equal to maximum")
	ErrNotPositive     = errors.New("value must be positive")
	ErrNotNegative     = errors.New("value must be negative")
	ErrNotZero         = errors.New("value must be zero")
	ErrIsZero          = errors.New("value must not be zero")
	ErrNotEven         = errors.New("value must be even")
	ErrNotOdd          = errors.New("value must be odd")
	ErrNotInRange      = errors.New("value is not in the allowed range")
	ErrNumberNotInList = errors.New("value is not in the allowed list")
	ErrInvalidDecimal  = errors.New("value has too many decimal places")
	ErrNotMultiple     = errors.New("value must be a multiple of the specified number")
)

// Numeric is a constraint for numeric types
type Numeric interface {
	~int | ~int32 | ~int64 | ~float32 | ~float64
}

// NumberValidator is a function type for validating numbers
type NumberValidator[T Numeric] func(T) error

// Number groups other validators in a single function identifying the value by its name
func Number[T Numeric](name string, validators ...NumberValidator[T]) NumberValidator[T] {
	return func(value T) error {
		var err error

		for _, validate := range validators {
			err = errors.Join(err, validate(value))
		}

		if err != nil {
			return fmt.Errorf("name=%s;errors=%w", name, err)
		}

		return nil
	}
}

// MinValue validates that the number is greater than or equal to the minimum value
func MinValue[T Numeric](min T) NumberValidator[T] {
	return func(value T) error {
		if value < min {
			return ErrMinValue
		}
		return nil
	}
}

// MaxValue validates that the number is less than or equal to the maximum value
func MaxValue[T Numeric](max T) NumberValidator[T] {
	return func(value T) error {
		if value > max {
			return ErrMaxValue
		}
		return nil
	}
}

// InRange validates that the number is within the specified range (inclusive)
func InRange[T Numeric](min, max T) NumberValidator[T] {
	return func(value T) error {
		if value < min || value > max {
			return ErrNotInRange
		}
		return nil
	}
}

// IsPositive validates that the number is greater than zero
func IsPositive[T Numeric]() NumberValidator[T] {
	return func(value T) error {
		if value <= 0 {
			return ErrNotPositive
		}
		return nil
	}
}

// IsNegative validates that the number is less than zero
func IsNegative[T Numeric]() NumberValidator[T] {
	return func(value T) error {
		if value >= 0 {
			return ErrNotNegative
		}
		return nil
	}
}

// IsZero validates that the number equals zero
func IsZero[T Numeric]() NumberValidator[T] {
	return func(value T) error {
		if value != 0 {
			return ErrNotZero
		}
		return nil
	}
}

// IsNonZero validates that the number does not equal zero
func IsNonZero[T Numeric]() NumberValidator[T] {
	return func(value T) error {
		if value == 0 {
			return ErrIsZero
		}
		return nil
	}
}

// IsEven validates that the number is even (only for integer types)
func IsEven[T Numeric]() NumberValidator[T] {
	return func(value T) error {
		// Convert to int64 for the modulo operation
		intVal := int64(value)
		if intVal%2 != 0 {
			return ErrNotEven
		}
		return nil
	}
}

// IsOdd validates that the number is odd (only for integer types)
func IsOdd[T Numeric]() NumberValidator[T] {
	return func(value T) error {
		// Convert to int64 for the modulo operation
		intVal := int64(value)
		if intVal%2 == 0 {
			return ErrNotOdd
		}
		return nil
	}
}

// NumberOneOf validates that the number is one of the allowed values
func NumberOneOf[T Numeric](allowed ...T) NumberValidator[T] {
	return func(value T) error {
		if slices.Contains(allowed, value) {
			return nil
		}
		return ErrNumberNotInList
	}
}
