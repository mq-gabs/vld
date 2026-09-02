package vslices

import (
	"errors"
	"fmt"
	"slices"

	"github.com/mq-gabs/vld/internal/utils"
	"github.com/mq-gabs/vld/validator"
)

var (
	ErrSliceMinLen   = errors.New("slice must have minimum length")
	ErrSliceMaxLen   = errors.New("slice must have maximum length")
	ErrSliceExactLen = errors.New("slice must have exact length")
	ErrSliceRequired = errors.New("slice is required and cannot be empty")
	ErrInvalidItem   = errors.New("invalid item")
)

// SliceValidator is a function type for validating slices
type SliceValidator[T any] func([]T) error

// Slice groups other validators in a single function identifying the value by its name
func Slice[T any](name string, validators ...SliceValidator[T]) SliceValidator[T] {
	return func(t []T) error {
		err := utils.NewErrorGroup(utils.WithSeparator(","))

		for _, validate := range validators {
			err.Join(validate(t))
		}

		if !err.IsNil() {
			return err
		}

		return nil
	}
}

// Required validates that the slice is not empty
func Required[T any]() SliceValidator[T] {
	return func(t []T) error {
		if len(t) == 0 {
			return ErrSliceRequired
		}
		return nil
	}
}

// MinLen validates that the slice has at least the specified number of items
func MinLen[T any](min int) SliceValidator[T] {
	if min < 0 {
		panic("minimum length cannot be negative")
	}

	return func(t []T) error {
		if len(t) < min {
			return ErrSliceMinLen
		}

		return nil
	}
}

// MaxLen validates that the slice has at most the specified number of items
func MaxLen[T any](max int) SliceValidator[T] {
	if max < 0 {
		panic("maximum length cannot be negative")
	}

	return func(t []T) error {
		if len(t) > max {
			return ErrSliceMaxLen
		}

		return nil
	}
}

// ExactLen validates that the slice has exactly the specified number of items
func ExactLen[T any](length int) SliceValidator[T] {
	if length < 0 {
		panic("exact length cannot be negative")
	}

	return func(t []T) error {
		if len(t) != length {
			return ErrSliceExactLen
		}

		return nil
	}
}

// InRange validates that the slice size is within the specified range (inclusive)
func InRange[T any](min, max int) SliceValidator[T] {
	if min < 0 || max < 0 {
		panic("range bounds cannot be negative")
	}
	if min > max {
		panic("minimum cannot be greater than maximum")
	}

	return func(t []T) error {
		size := len(t)
		if size < min || size > max {
			return fmt.Errorf("slice size %d is not in range [%d, %d]", size, min, max)
		}
		return nil
	}
}

// Each validates each item in the slice using the provided validator
func Each[T any](validate validator.GenericValidator[T]) SliceValidator[T] {
	return func(t []T) error {
		err := utils.NewErrorGroup(utils.WithSeparator(","))
		for i, v := range t {
			if itemErr := validate(v); itemErr != nil {
				err.Join(fmt.Errorf("index %d: %w", i, itemErr))
			}
		}
		if !err.IsNil() {
			return errors.Join(ErrInvalidItem, err)
		}
		return nil
	}
}

// Contains validates that the slice contains the specified item
func Contains[T comparable](item T) SliceValidator[T] {
	return func(t []T) error {
		if slices.Contains(t, item) {
			return nil
		}
		return fmt.Errorf("slice does not contain %v", item)
	}
}

// NotContains validates that the slice does not contain the specified item
func NotContains[T comparable](item T) SliceValidator[T] {
	return func(t []T) error {
		if slices.Contains(t, item) {
			return fmt.Errorf("slice contains forbidden item %v", item)
		}
		return nil
	}
}

// NoDuplicates validates that the slice has no duplicate items
func NoDuplicates[T comparable]() SliceValidator[T] {
	return func(t []T) error {
		seen := make(map[T]bool)
		for i, v := range t {
			if seen[v] {
				return fmt.Errorf("duplicate item at index %d: %v", i, v)
			}
			seen[v] = true
		}
		return nil
	}
}
