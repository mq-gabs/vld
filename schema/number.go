package schema

import (
	"errors"
	"fmt"
)

type NumberType interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64
}

type numberSchema[T NumberType] struct {
	baseSchema[T]
}

func Number[T NumberType]() *numberSchema[T] {
	return &numberSchema[T]{
		baseSchema: newBaseSchema[T](),
	}
}

func (ns *numberSchema[T]) Custom(fn Validator[T]) *numberSchema[T] {
	ns.appendValidator(fn)

	return ns
}

func (ns *numberSchema[T]) Min(min T) *numberSchema[T] {
	ns.appendValidator(func(i T) error {
		if i < min {
			return fmt.Errorf("required min value: %v", min)
		}

		return nil
	})

	return ns
}

func (is *numberSchema[T]) Max(max T) *numberSchema[T] {
	is.appendValidator(func(i T) error {
		if i > max {
			return fmt.Errorf("required max value: %v", i)
		}

		return nil
	})

	return is
}

func (is *numberSchema[T]) Equals(target T) *numberSchema[T] {
	is.appendValidator(func(i T) error {
		if i != target {
			return fmt.Errorf("value must be equal to: %v", target)
		}

		return nil
	})

	return is
}

func (is *numberSchema[T]) NonZero() *numberSchema[T] {
	is.appendValidator(func(i T) error {
		if i == 0 {
			return errors.New("value must be non zero")
		}

		return nil
	})

	return is
}

func (is *numberSchema[T]) Positive() *numberSchema[T] {
	is.appendValidator(func(i T) error {
		if i < 0 {
			return errors.New("value must be positive")
		}

		return nil
	})

	return is
}

func (is *numberSchema[T]) Negative() *numberSchema[T] {
	is.appendValidator(func(i T) error {
		if i > 0 {
			return errors.New("value must be negative")
		}

		return nil
	})

	return is
}
