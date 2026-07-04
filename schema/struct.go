package schema

import (
	"errors"
)

type schemaStruct[T any] struct {
	baseSchema[T]
	tupleSet TupleSet[T]
}

type SchemaStruct[T any] interface {
	Schema[T]
}

func Struct[T any](fn TupleSet[T]) SchemaStruct[T] {
	return &schemaStruct[T]{
		baseSchema: newBaseSchema[T](),
		tupleSet:   fn,
	}
}

func (ss *schemaStruct[T]) Validate(v any) error {
	typedV, ok := v.(*T)
	if !ok {
		return errors.New("invalid type")
	}

	b := newTupleBuilder()

	ss.tupleSet(b, typedV)

	var err error
	for _, t := range b.tuples {
		if e := t.Validate(); e != nil {
			err = errors.Join(err, e)
		}
	}

	return err
}
