package schema

import (
	"errors"
)

type TupleSet[T any] func(Builder, *T)

type Tuple[T any] struct {
	fields []T
	schema Schema[T]
}

func (t Tuple[T]) Validate() error {
	var err error
	for _, f := range t.fields {
		if e := t.schema.Validate(f); e != nil {
			err = errors.Join(err, e)
		}
	}

	return err
}

type tupleBuilder struct {
	tuples []Tuple[any]
}

func newTupleBuilder() *tupleBuilder {
	return &tupleBuilder{}
}

type Builder interface {
	F(f any, s Schema[any])
	Fs(f []any, s Schema[any])
}

func (b *tupleBuilder) F(f any, s Schema[any]) {
	b.tuples = append(b.tuples, Tuple[any]{
		fields: []any{f},
		schema: s,
	})
}

func (b *tupleBuilder) Fs(f []any, s Schema[any]) {
	b.tuples = append(b.tuples, Tuple[any]{f, s})
}
