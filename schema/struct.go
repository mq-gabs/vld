package schema

import "errors"

type SchemaStruct[T any] struct {
	baseSchema[T]
	tupleSet TupleSet[T]
}

func Struct[T any](fn TupleSet[T]) *SchemaStruct[T] {
	return &SchemaStruct[T]{
		baseSchema: newBaseSchema[T](),
		tupleSet:   fn,
	}
}

func (ss *SchemaStruct[T]) Validate(v any) error {
	typedV, ok := v.(*T)
	if !ok {
		return errors.New("invalid type")
	}

	tuples := ss.tupleSet(typedV)

	var err error
	for _, t := range tuples {
		if e := t.Validate(); e != nil {
			errors.Join(err, e)
		}
	}

	return err
}
