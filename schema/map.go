package schema

import (
	"errors"
	"fmt"
)

type mapSchema[T comparable, U any] struct {
	baseSchema[map[T]U]
}

func Map[T comparable, U any]() *mapSchema[T, U] {
	return &mapSchema[T, U]{
		baseSchema: newBaseSchema[map[T]U](),
	}
}

func (ms *mapSchema[T, U]) Custom(fn Validator[map[T]U]) *mapSchema[T, U] {
	ms.appendValidator(fn)

	return ms
}

func (ms *mapSchema[T, U]) MaxLength(max int) *mapSchema[T, U] {
	ms.appendValidator(func(m map[T]U) error {
		if len(m) > max {
			return fmt.Errorf("required max length: %v", max)
		}

		return nil
	})

	return ms
}

func (ms *mapSchema[T, U]) MinLength(min int) *mapSchema[T, U] {
	ms.appendValidator(func(m map[T]U) error {
		if len(m) < min {
			return fmt.Errorf("required min length: %v", min)
		}

		return nil
	})

	return ms
}

func (ms *mapSchema[T, U]) Child(schema Schema[U]) *mapSchema[T, U] {
	ms.appendValidator(func(m map[T]U) error {
		var err error
		for key, value := range m {
			schemaErr := schema.Validate(value)
			if schemaErr != nil {
				err = errors.Join(err, fmt.Errorf("[%v]: %v", key, schemaErr.Error()))
			}
		}

		return err
	})

	return ms
}
