package slicev

import (
	"errors"

	"github.com/mq-gabs/vld/config/utils"
	"github.com/mq-gabs/vld/config/validate"
)

type ConfigSlice[T any] struct {
	MinLen utils.V[int]
	MaxLen utils.V[int]

	NonEmpty bool

	ValueValidator validate.Validator[T]

	Custom validate.Validate[[]T]
}

type sliceValidator[T any] struct {
	valueValidator validate.Validator[T]
	validations    []validate.Validate[[]T]
}

func (sv *sliceValidator[T]) append(v validate.Validate[[]T]) {
	if v == nil {
		return
	}
	sv.validations = append(sv.validations, v)
}

func (sv *sliceValidator[T]) Validate(s []T) error {
	var err error

	for _, v := range sv.validations {
		err = errors.Join(err, v(s))
	}

	if sv.valueValidator != nil {
		for _, v := range s {
			err = errors.Join(err, sv.valueValidator.Validate(v))
		}
	}

	return err
}

func (c ConfigSlice[T]) Build() validate.Validator[[]T] {
	sv := &sliceValidator[T]{
		valueValidator: c.ValueValidator,
	}

	if c.MinLen.IsSet() {
		sv.append(buildSliceMin[T](c.MinLen.Get()))
	}

	if c.MaxLen.IsSet() {
		sv.append(buildSliceMax[T](c.MaxLen.Get()))
	}

	if c.NonEmpty {
		sv.append(buildSliceNonEmpty[T]())
	}

	if c.Custom != nil {
		sv.append(c.Custom)
	}

	return sv
}

/* =========================
   BUILDERS (correct pattern)
========================= */

func buildSliceMin[T any](min int) validate.Validate[[]T] {
	return func(s []T) error {
		if len(s) < min {
			return errors.New("slice too small")
		}
		return nil
	}
}

func buildSliceMax[T any](max int) validate.Validate[[]T] {
	return func(s []T) error {
		if len(s) > max {
			return errors.New("slice too large")
		}
		return nil
	}
}

func buildSliceNonEmpty[T any]() validate.Validate[[]T] {
	return func(s []T) error {
		if len(s) == 0 {
			return errors.New("slice must not be empty")
		}
		return nil
	}
}
