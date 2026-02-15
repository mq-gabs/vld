package schema

import (
	"errors"
	"fmt"
)

type SchemaStructType interface {
	SchemaJSON() map[string]any
}

type structSchema[T SchemaStructType] struct {
	baseSchema[T]
	schemas map[string]Schema[any]
}

func Struct[T SchemaStructType]() *structSchema[T] {
	return &structSchema[T]{
		baseSchema: newBaseSchema[T](),
		schemas:    make(map[string]Schema[any]),
	}
}

func (ss *structSchema[T]) Field(key string, schema Schema[any]) *structSchema[T] {
	ss.schemas[key] = schema

	return ss
}

func (ss *structSchema[T]) Validate(structValue any) error {
	structValueTyped, ok := structValue.(SchemaStructType)
	if !ok {
		return errors.New("struct must implement SchemaJSON")
	}

	json := structValueTyped.SchemaJSON()

	var err error
	for key, value := range json {
		schema, ok := ss.schemas[key]
		if !ok {
			return fmt.Errorf("no schema set for field: %v", key)
		}

		schemaErr := schema.Validate(value)
		if schemaErr != nil {
			err = errors.Join(err, fmt.Errorf("[%v]: %v", key, schemaErr.Error()))
		}
	}

	return err
}
