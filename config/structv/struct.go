package structv

import (
	"errors"
	"fmt"

	"github.com/mq-gabs/vld/config/validate"
)

type Pair struct {
	validate func() error
}

func NewPair[T any](
	name string,
	value T,
	validator validate.Validator[T],
) Pair {
	return Pair{
		validate: func() error {
			if err := validator.Validate(value); err != nil {
				return fmt.Errorf("%s: %w", name, err)
			}

			return nil
		},
	}
}

type ConfigStruct[T any] struct {
	Nullable bool

	Validate func(T) []Pair
}

type structValidator[T any] struct {
	nullable bool
	validate func(T) []Pair
}

func (sv *structValidator[T]) Validate(v *T) error {
	if v == nil {
		if sv.nullable {
			return nil
		}

		return errors.New("value cannot be nil")
	}

	var err error

	for _, pair := range sv.validate(*v) {
		if pair.validate == nil {
			continue
		}

		err = errors.Join(err, pair.validate())
	}

	return err
}

func (cs *ConfigStruct[T]) Build() validate.Validator[*T] {
	return &structValidator[T]{
		nullable: cs.Nullable,
		validate: cs.Validate,
	}
}
