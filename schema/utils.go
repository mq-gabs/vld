package schema

import "errors"

type TupleSet[T any] func(*T) []Tuple[any]

type Tuple[T any] struct {
	fields []T
	schema Schema[T]
}

func (t Tuple[T]) Validate() error {
	var err error
	for _, f := range t.fields {
		if e := t.schema.Validate(f); e != nil {
			errors.Join(err, e)
		}
	}

	return err
}

func T(f any, s Schema[any]) Tuple[any] {
	return Tuple[any]{fields: []any{f}, schema: s}
}

func M(f []any, s Schema[any]) Tuple[any] {
	return Tuple[any]{fields: f, schema: s}
}
