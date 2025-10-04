package schema

import (
	"errors"
	"fmt"
)

type intSchema struct {
	baseSchema[int]
}

func Int() *intSchema {
	return &intSchema{
		baseSchema: newBaseSchema[int](),
	}
}

func (is *intSchema) Min(min int) *intSchema {
	is.appendValidator(func(i int) error {
		if i < min {
			return fmt.Errorf("required min value: %v", min)
		}

		return nil
	})

	return is
}

func (is *intSchema) Max(max int) *intSchema {
	is.appendValidator(func(i int) error {
		if i > max {
			return fmt.Errorf("required max value: %v", i)
		}

		return nil
	})

	return is
}

func (is *intSchema) Equals(target int) *intSchema {
	is.appendValidator(func(i int) error {
		if i != target {
			return fmt.Errorf("value must be equal to: %v", target)
		}

		return nil
	})

	return is
}

func (is *intSchema) NonZero() *intSchema {
	is.appendValidator(func(i int) error {
		if i == 0 {
			return errors.New("value must be non zero")
		}

		return nil
	})

	return is
}

func (is *intSchema) Positive() *intSchema {
	is.appendValidator(func(i int) error {
		if i < 0 {
			return errors.New("value must be positive")
		}

		return nil
	})

	return is
}

func (is *intSchema) Negative() *intSchema {
	is.appendValidator(func(i int) error {
		if i > 0 {
			return errors.New("value must be negative")
		}

		return nil
	})

	return is
}
