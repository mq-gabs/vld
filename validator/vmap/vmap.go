package vmap

import (
	"errors"
	"fmt"

	"github.com/mq-gabs/vld/internal/utils"
	"github.com/mq-gabs/vld/validator"
)

var (
	ErrMapMinLen        = errors.New("map must have min length")
	ErrMapMaxLen        = errors.New("map must have max length")
	ErrMapExactLen      = errors.New("map must have exact length")
	ErrMapRequired      = errors.New("map is required and cannot be empty")
	ErrMapKeyNotFound   = errors.New("required key not found in map")
	ErrMapKeyNotAllowed = errors.New("key is not allowed in map")
	ErrMapKeyValidation = errors.New("validation failed for key value")
)

// MapValidator is a function type for validating maps
type MapValidator[K comparable, V any] func(map[K]V) error

// Map groups other validators in a single function identifying the value by its name
func Map[K comparable, V any](name string, validators ...MapValidator[K, V]) MapValidator[K, V] {
	return func(m map[K]V) error {
		err := utils.NewErrorGroup(utils.WithSeparator(","))

		for _, validate := range validators {
			err.Join(validate(m))
		}

		if !err.IsNil() {
			return err
		}

		return nil
	}
}

// Required validates that the map is not empty
func Required[K comparable, V any]() MapValidator[K, V] {
	return func(m map[K]V) error {
		if len(m) == 0 {
			return ErrMapRequired
		}
		return nil
	}
}

// MinLen validates that the map has at least the specified number of key-value pairs
func MinLen[K comparable, V any](min int) MapValidator[K, V] {
	return func(m map[K]V) error {
		if len(m) < min {
			return ErrMapMinLen
		}
		return nil
	}
}

// MaxLen validates that the map has at most the specified number of key-value pairs
func MaxLen[K comparable, V any](max int) MapValidator[K, V] {
	return func(m map[K]V) error {
		if len(m) > max {
			return ErrMapMaxLen
		}
		return nil
	}
}

// ExactLen validates that the map has exactly the specified number of key-value pairs
func ExactLen[K comparable, V any](length int) MapValidator[K, V] {
	return func(m map[K]V) error {
		if len(m) != length {
			return ErrMapExactLen
		}
		return nil
	}
}

// InRange validates that the map size is within the specified range (inclusive)
func InRange[K comparable, V any](min, max int) MapValidator[K, V] {
	return func(m map[K]V) error {
		size := len(m)
		if size < min || size > max {
			return fmt.Errorf("map size %d is not in range [%d, %d]", size, min, max)
		}
		return nil
	}
}

// ContainsKey validates that the map contains a specific key
func ContainsKey[K comparable, V any](key K) MapValidator[K, V] {
	return func(m map[K]V) error {
		if _, ok := m[key]; !ok {
			return fmt.Errorf("%w: %v", ErrMapKeyNotFound, key)
		}
		return nil
	}
}

// NotContainsKey validates that the map does not contain a specific key
func NotContainsKey[K comparable, V any](key K) MapValidator[K, V] {
	return func(m map[K]V) error {
		if _, ok := m[key]; ok {
			return fmt.Errorf("%w: %v", ErrMapKeyNotAllowed, key)
		}
		return nil
	}
}

// HasKeys validates that the map contains all specified keys
func HasKeys[K comparable, V any](keys ...K) MapValidator[K, V] {
	return func(m map[K]V) error {
		for _, key := range keys {
			if _, ok := m[key]; !ok {
				return fmt.Errorf("%w: %v", ErrMapKeyNotFound, key)
			}
		}
		return nil
	}
}

// NotHasKeys validates that the map does not contain any of the specified keys
func NotHasKeys[K comparable, V any](keys ...K) MapValidator[K, V] {
	return func(m map[K]V) error {
		for _, key := range keys {
			if _, ok := m[key]; ok {
				return fmt.Errorf("%w: %v", ErrMapKeyNotAllowed, key)
			}
		}
		return nil
	}
}

// EachValue validates each value in the map using the provided validator
func EachValue[K comparable, V any](validate validator.GenericValidator[V]) MapValidator[K, V] {
	return func(m map[K]V) error {
		err := utils.NewErrorGroup(utils.WithSeparator(","))
		for key, val := range m {
			if validationErr := validate(val); validationErr != nil {
				err.Join(fmt.Errorf("key %v: %w", key, validationErr))
			}
		}
		if !err.IsNil() {
			return err
		}
		return nil
	}
}

// EachKey validates each key in the map using the provided validator
func EachKey[K comparable, V any](validate validator.GenericValidator[K]) MapValidator[K, V] {
	return func(m map[K]V) error {
		err := utils.NewErrorGroup(utils.WithSeparator(","))
		for key := range m {
			if validationErr := validate(key); validationErr != nil {
				err.Join(fmt.Errorf("key %v: %w", key, validationErr))
			}
		}
		if !err.IsNil() {
			return err
		}
		return nil
	}
}

// Each validates each value in the map using the provided validator (alias for EachValue)
func Each[K comparable, V any](validate validator.GenericValidator[V]) MapValidator[K, V] {
	return EachValue[K, V](validate)
}
