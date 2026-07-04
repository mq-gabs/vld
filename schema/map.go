package schema

import (
	"errors"
	"fmt"
)

type schemaMap[T comparable, U any] struct {
	baseSchema[map[T]U]
}

type SchemaMap[T comparable, U any] interface {
	Schema[map[T]U]

	Custom(Validator[map[T]U]) SchemaMap[T, U]
	LengthMax(int) SchemaMap[T, U]
	LengthMin(int) SchemaMap[T, U]
	Child(Schema[U]) SchemaMap[T, U]
}

func Map[T comparable, U any]() SchemaMap[T, U] {
	return &schemaMap[T, U]{
		baseSchema: newBaseSchema[map[T]U](),
	}
}

func (ms *schemaMap[T, U]) Custom(fn Validator[map[T]U]) SchemaMap[T, U] {
	ms.appendValidator(fn)

	return ms
}

func (ms *schemaMap[T, U]) LengthMax(max int) SchemaMap[T, U] {
	ms.appendValidator(func(m map[T]U) error {
		if len(m) > max {
			return fmt.Errorf("required max length: %v", max)
		}

		return nil
	})

	return ms
}

func (ms *schemaMap[T, U]) LengthMin(min int) SchemaMap[T, U] {
	ms.appendValidator(func(m map[T]U) error {
		if len(m) < min {
			return fmt.Errorf("required min length: %v", min)
		}

		return nil
	})

	return ms
}

func (ms *schemaMap[T, U]) Child(schema Schema[U]) SchemaMap[T, U] {
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
