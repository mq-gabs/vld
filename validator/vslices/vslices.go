package vslices

import (
	"errors"
	"fmt"

	"github.com/mq-gabs/vld/validator"
)

var (
	ErrInvalidItem = errors.New("invalid item")
)

type SliceValidator[T any] func([]T) error

func Slice[T any](name string, validators ...SliceValidator[T]) SliceValidator[T] {
	return func(t []T) error {
		var err error

		for _, validate := range validators {
			err = errors.Join(err, validate(t))
		}

		if err != nil {
			return fmt.Errorf("name=%s;errors=%w", name, err)
		}

		return nil
	}
}

func MinLen[T any](min int) SliceValidator[T] {
	if min < 0 {
		panic("minimum length cannot be negative")
	}

	return func(t []T) error {
		if len(t) < min {
			return validator.ErrMinLen
		}

		return nil
	}
}

func MaxLen[T any](max int) SliceValidator[T] {
	return func(t []T) error {
		if len(t) > max {
			return validator.ErrMaxLen
		}

		return nil
	}
}

func Each[T any](validator validator.GenericValidator[T]) SliceValidator[T] {
	return func(t []T) error {
		for _, v := range t {
			if err := validator(v); err != nil {
				return errors.Join(ErrInvalidItem, err)
			}
		}

		return nil
	}
}
