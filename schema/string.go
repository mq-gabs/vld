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

type schemaString struct {
	baseSchema[string]
}

type SchemaString interface {
	Schema[string]

	Custom(Validator[string]) SchemaString
	LengthMin(int) SchemaString
	LengthMax(int) SchemaString
	UUID() SchemaString
	Email() SchemaString
	URL() SchemaString
	Enum([]string) SchemaString
}

func String() SchemaString {
	return &schemaString{
		baseSchema: newBaseSchema[string](),
	}
}

func (ss *schemaString) Custom(fn Validator[string]) SchemaString {
	ss.appendValidator(fn)

	return ss
}

func (ss *schemaString) LengthMin(minSize int) SchemaString {
	ss.appendValidator(func(value string) error {
		if len(value) < minSize {
			return fmt.Errorf("required min length: %v", minSize)
		}

		return nil
	})

	return ss
}

func (ss *schemaString) LengthMax(maxSize int) SchemaString {
	ss.appendValidator(func(value string) error {
		if len(value) > maxSize {
			return fmt.Errorf("required max length: %v", maxSize)
		}

		return nil
	})

	return ss
}

func (ss *schemaString) UUID() SchemaString {
	ss.appendValidator(func(s string) error {
		if !regexUUID.MatchString(s) {
			return errors.New("must be valid UUID")
		}

		return nil
	})

	return ss
}

func (ss *schemaString) Email() SchemaString {
	ss.appendValidator(func(s string) error {
		if !regexEmail.MatchString(s) {
			return errors.New("must be valid email")
		}

		return nil
	})

	return ss
}

func (ss *schemaString) URL() SchemaString {
	ss.appendValidator(func(s string) error {
		if !regexURL.MatchString(s) {
			return errors.New("must be valid URL")
		}

		return nil
	})

	return ss
}

func (ss *schemaString) Enum(enum []string) SchemaString {
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
