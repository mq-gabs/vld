package schema

import (
	"errors"
	"fmt"
	"regexp"
	"slices"
)

var (
	regexEmail = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	regexUUID  = regexp.MustCompile(`^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[1-5][a-fA-F0-9]{3}-[89abAB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$`)
	regexURL   = regexp.MustCompile(`^(https?|ftp):\/\/[^\s/$.?#].[^\s]*$`)
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

func (ss *stringSchema) UUID() *stringSchema {
	ss.appendValidator(func(s string) error {
		if !regexUUID.MatchString(s) {
			return errors.New("must be valid UUID")
		}

		return nil
	})

	return ss
}

func (ss *stringSchema) Email() *stringSchema {
	ss.appendValidator(func(s string) error {
		if !regexEmail.MatchString(s) {
			return errors.New("must be valid email")
		}

		return nil
	})

	return ss
}

func (ss *stringSchema) URL() *stringSchema {
	ss.appendValidator(func(s string) error {
		if !regexURL.MatchString(s) {
			return errors.New("must be valid URL")
		}

		return nil
	})

	return ss
}

func (ss *stringSchema) Enum(enum []string) *stringSchema {
	if len(enum) == 0 {
		ss.appendValidator(func(s string) error {
			return errors.New("invalid setting, enum must not be empty")
		})

		return ss
	}

	ss.appendValidator(func(s string) error {
		if !slices.Contains(enum, s) {
			return fmt.Errorf("must be one of %v", enum)
		}

		return nil
	})

	return ss
}
