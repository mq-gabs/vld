package schema

import (
	"fmt"
	"slices"
)

type sliceSchema[T comparable] struct {
	baseSchema[[]T]
}

func Slice[T comparable]() *sliceSchema[T] {
	return &sliceSchema[T]{
		baseSchema: newBaseSchema[[]T](),
	}
}

func (ss *sliceSchema[T]) Custom(fn Validator[[]T]) *sliceSchema[T] {
	ss.appendValidator(fn)

	return ss
}

func (ss *sliceSchema[T]) MinLength(minLen int) *sliceSchema[T] {
	ss.appendValidator(func(a []T) error {
		if len(a) < minLen {
			return fmt.Errorf("required min length: %v", minLen)
		}

		return nil
	})

	return ss
}

func (ss *sliceSchema[T]) MaxLength(maxLen int) *sliceSchema[T] {
	ss.appendValidator(func(a []T) error {
		if len(a) > maxLen {
			return fmt.Errorf("required max length: %v", maxLen)
		}

		return nil
	})

	return ss
}

func (ss *sliceSchema[T]) Contains(target T) *sliceSchema[T] {
	ss.appendValidator(func(t []T) error {
		if !slices.Contains(t, target) {
			return fmt.Errorf("slice must contain value: %v", target)
		}

		return nil
	})

	return ss
}
