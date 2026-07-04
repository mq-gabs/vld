package schema

import (
	"fmt"
	"slices"
)

type schemaSlice[T comparable] struct {
	baseSchema[[]T]
}

type SchemaSlice[T comparable] interface {
	Schema[[]T]

	Custom(Validator[[]T]) SchemaSlice[T]
	LengthMin(int) SchemaSlice[T]
	LengthMax(int) SchemaSlice[T]
	Contains(T) SchemaSlice[T]
}

func Slice[T comparable]() SchemaSlice[T] {
	return &schemaSlice[T]{
		baseSchema: newBaseSchema[[]T](),
	}
}

func (ss *schemaSlice[T]) Custom(fn Validator[[]T]) SchemaSlice[T] {
	ss.appendValidator(fn)

	return ss
}

func (ss *schemaSlice[T]) LengthMin(minLen int) SchemaSlice[T] {
	ss.appendValidator(func(a []T) error {
		if len(a) < minLen {
			return fmt.Errorf("required min length: %v", minLen)
		}

		return nil
	})

	return ss
}

func (ss *schemaSlice[T]) LengthMax(maxLen int) SchemaSlice[T] {
	ss.appendValidator(func(a []T) error {
		if len(a) > maxLen {
			return fmt.Errorf("required max length: %v", maxLen)
		}

		return nil
	})

	return ss
}

func (ss *schemaSlice[T]) Contains(target T) SchemaSlice[T] {
	ss.appendValidator(func(t []T) error {
		if !slices.Contains(t, target) {
			return fmt.Errorf("slice must contain value: %v", target)
		}

		return nil
	})

	return ss
}
