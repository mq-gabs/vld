package schema

import (
	"errors"
	"slices"
)

type Validator[T any] func(T) error
type Parser[T any] func(T) T
type IsZero[T any] func(T) bool

type Schema[T any] interface {
	Validate(any) error
}

var initialValidate = func() error {
	return errors.New("not implemented")
}

type baseSchema[T any] struct {
	validators []Validator[T]
	parsers    []Parser[T]
	isZero     IsZero[T]
	optional   bool
}

func newBaseSchema[T any]() baseSchema[T] {
	return baseSchema[T]{}
}

func (bs *baseSchema[T]) Validate(value any) error {
	typedValue, ok := value.(T)
	if !ok {
		return errors.New("invalid type")
	}

	if bs.optional {
		if bs.isZero == nil {
			return errors.New("[internal] zero function not set")
		}

		if bs.isZero(typedValue) {
			return nil
		}
	}

	if len(bs.validators) == 0 {
		return initialValidate()
	}

	var err error
	for _, valid := range bs.validators {
		err = errors.Join(err, valid(typedValue))
	}

	return err
}

func (bs *baseSchema[T]) appendValidator(newValidator Validator[T]) {
	bs.validators = append(bs.validators, newValidator)
}

func (bs *baseSchema[T]) clone() baseSchema[T] {
	return baseSchema[T]{
		validators: slices.Clone(bs.validators),
		parsers:    slices.Clone(bs.parsers),
		isZero:     bs.isZero,
		optional:   bs.optional,
	}
}
