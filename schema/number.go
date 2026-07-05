package schema

import (
	"errors"
	"fmt"
)

type TypeNumber interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64
}

type schemaNumber[T TypeNumber] struct {
	baseSchema[T]
}

type SchemaNumber[T TypeNumber] interface {
	Schema[T]

	Clone() SchemaNumber[T]
	Optional() SchemaNumber[T]
	Custom(Validator[T]) SchemaNumber[T]
	Min(T) SchemaNumber[T]
	Max(T) SchemaNumber[T]
	Equals(T) SchemaNumber[T]
	NonZero() SchemaNumber[T]
	Positive() SchemaNumber[T]
	NonPositive() SchemaNumber[T]
	Negative() SchemaNumber[T]
	NonNegative() SchemaNumber[T]
}

func Number[T TypeNumber]() SchemaNumber[T] {
	bs := &schemaNumber[T]{
		baseSchema: newBaseSchema[T](),
	}

	bs.isZero = func(t T) bool {
		return t == 0
	}

	return bs
}

func (ns *schemaNumber[T]) Optional() SchemaNumber[T] {
	ns.optional = true
	return ns
}

func (ns *schemaNumber[T]) Clone() SchemaNumber[T] {
	return &schemaNumber[T]{
		baseSchema: ns.baseSchema.clone(),
	}
}

func (ns *schemaNumber[T]) Custom(fn Validator[T]) SchemaNumber[T] {
	ns.appendValidator(fn)

	return ns
}

func (ns *schemaNumber[T]) Min(min T) SchemaNumber[T] {
	ns.appendValidator(func(i T) error {
		if i < min {
			return fmt.Errorf("required min value: %v", min)
		}

		return nil
	})

	return ns
}

func (is *schemaNumber[T]) Max(max T) SchemaNumber[T] {
	is.appendValidator(func(i T) error {
		if i > max {
			return fmt.Errorf("required max value: %v", i)
		}

		return nil
	})

	return is
}

func (is *schemaNumber[T]) Equals(target T) SchemaNumber[T] {
	is.appendValidator(func(i T) error {
		if i != target {
			return fmt.Errorf("value must be equal to: %v", target)
		}

		return nil
	})

	return is
}

func (is *schemaNumber[T]) NonZero() SchemaNumber[T] {
	is.appendValidator(func(i T) error {
		if i == 0 {
			return errors.New("value must be non zero")
		}

		return nil
	})

	return is
}

func (is *schemaNumber[T]) NonPositive() SchemaNumber[T] {
	is.appendValidator(func(t T) error {
		if t > 0 {
			return errors.New("value must be non-positive")
		}

		return nil
	})

	return is
}

func (is *schemaNumber[T]) Positive() SchemaNumber[T] {
	is.appendValidator(func(i T) error {
		if i <= 0 {
			return errors.New("value must be positive")
		}

		return nil
	})

	return is
}

func (is *schemaNumber[T]) NonNegative() SchemaNumber[T] {
	is.appendValidator(func(t T) error {
		if t < 0 {
			return errors.New("value must be non-negative")
		}

		return nil
	})

	return is
}

func (is *schemaNumber[T]) Negative() SchemaNumber[T] {
	is.appendValidator(func(i T) error {
		if i >= 0 {
			return errors.New("value must be negative")
		}

		return nil
	})

	return is
}
