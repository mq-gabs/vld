package schema

import (
	"errors"
)

type schemaStruct[T any] struct {
	baseSchema[T]
	tupleSet     TupleSet[T]
	isStructZero IsZero[*T]
}

type SchemaStruct[T any] interface {
	Schema[T]

	Clone() SchemaStruct[T]
	Optional() SchemaStruct[T]
}

func Struct[T any](fn TupleSet[T]) SchemaStruct[T] {
	ss := &schemaStruct[T]{
		baseSchema: newBaseSchema[T](),
		tupleSet:   fn,
	}

	ss.isStructZero = func(t *T) bool {
		return t == nil
	}

	return ss
}

func (ss *schemaStruct[T]) Optional() SchemaStruct[T] {
	ss.optional = true

	return ss
}

func (ss *schemaStruct[T]) Clone() SchemaStruct[T] {
	return &schemaStruct[T]{
		baseSchema: ss.baseSchema.clone(),
		tupleSet:   ss.tupleSet,
	}
}

func (ss *schemaStruct[T]) Validate(v any) error {
	typedV, ok := v.(*T)
	if !ok {
		return errors.New("invalid type")
	}

	if ss.optional {
		if ss.isStructZero == nil {
			return errors.New("[internal] is struct zero function not set")
		}

		if ss.isStructZero(typedV) {
			return nil
		}
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
