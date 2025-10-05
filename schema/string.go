package schema

import (
	"fmt"
)

type stringSchema struct {
	baseSchema[string]
}

func String() *stringSchema {
	return &stringSchema{
		baseSchema: newBaseSchema[string](),
	}
}

func (ss *stringSchema) Custom(fn Validator[string]) *stringSchema {
	ss.appendValidator(fn)

	return ss
}

func (ss *stringSchema) MinLength(minSize int) *stringSchema {
	ss.appendValidator(func(value string) error {
		if len(value) < minSize {
			return fmt.Errorf("required min length: %v", minSize)
		}

		return nil
	})

	return ss
}

func (ss *stringSchema) MaxLength(maxSize int) *stringSchema {
	ss.appendValidator(func(value string) error {
		if len(value) > maxSize {
			return fmt.Errorf("required max length: %v", maxSize)
		}

		return nil
	})

	return ss
}
