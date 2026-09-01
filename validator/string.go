package validator

import (
	"errors"
	"fmt"
	"regexp"
	"slices"
	"strings"
	"unicode"
)

var (
	ErrInvalidEmail  = errors.New("invalid email format")
	ErrInvalidUUID   = errors.New("invalid uuid format")
	ErrNoMatch       = errors.New("value does not match required pattern")
	ErrNotAlphaNum   = errors.New("value must be alphanumeric")
	ErrInvalidStart  = errors.New("value must start with required prefix")
	ErrInvalidEnd    = errors.New("value must end with required suffix")
	ErrNotAllowed    = errors.New("value contains forbidden characters")
	ErrNotInList     = errors.New("value is not in the allowed list")
	ErrHasWhitespace = errors.New("value must not contain whitespace")
	ErrNotLowerCase  = errors.New("value must be lowercase")
	ErrNotUpperCase  = errors.New("value must be uppercase")
)

type StringValidator func(string) error

// String groups other validators in a single function identifing the value by its name
func String(name string, validators ...StringValidator) StringValidator {
	return func(value string) error {
		var err error

		for _, validate := range validators {
			err = errors.Join(err, validate(value))
		}

		if err != nil {
			return fmt.Errorf("name=%s;errors=%w", name, err)
		}

		return nil
	}
}

// StringRequired validates empty string
func StringRequired() StringValidator {
	return func(s string) error {
		if s == "" {
			return ErrValueIsRequired
		}

		return nil
	}
}

// MinLen validates minimum string length
func MinLen(min int) StringValidator {
	return func(s string) error {
		if len(s) < min {
			return ErrMinLen
		}

		return nil
	}
}

// MaxLen validates maximum string length
func MaxLen(max int) StringValidator {
	return func(s string) error {
		if len(s) > max {
			return ErrMaxLen
		}

		return nil
	}
}

// Email validates email pattern
func Email() StringValidator {
	emailPattern := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return func(s string) error {
		if !emailPattern.MatchString(s) {
			return ErrInvalidEmail
		}
		return nil
	}
}

// UUIDv4 validates v4 pattern: xxxxxxxx-xxxx-4xxx-yxxx-xxxxxxxxxxxx
func UUIDv4() StringValidator {
	uuidPattern := regexp.MustCompile(`^[0-9a-f]{8}-[0-9a-f]{4}-4[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$|^[0-9a-f]{8}-[0-9a-f]{4}-[1-5][0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$`)
	return func(s string) error {
		if !uuidPattern.MatchString(strings.ToLower(s)) {
			return ErrInvalidUUID
		}
		return nil
	}
}

// Pattern validates that the string matches a given regex pattern
func Pattern(pattern string) StringValidator {
	regex := regexp.MustCompile(pattern)
	return func(s string) error {
		if !regex.MatchString(s) {
			return ErrNoMatch
		}
		return nil
	}
}

// Alphanumeric validates that the string contains only alphanumeric characters
func Alphanumeric() StringValidator {
	return func(s string) error {
		for _, r := range s {
			if !unicode.IsLetter(r) && !unicode.IsDigit(r) {
				return ErrNotAlphaNum
			}
		}
		return nil
	}
}

// StartsWith validates that the string starts with the given prefix
func StartsWith(prefix string) StringValidator {
	return func(s string) error {
		if !strings.HasPrefix(s, prefix) {
			return ErrInvalidStart
		}
		return nil
	}
}

// EndsWith validates that the string ends with the given suffix
func EndsWith(suffix string) StringValidator {
	return func(s string) error {
		if !strings.HasSuffix(s, suffix) {
			return ErrInvalidEnd
		}
		return nil
	}
}

// NotContains validates that the string does not contain any of the forbidden characters
func NotContains(forbidden string) StringValidator {
	return func(s string) error {
		if strings.ContainsAny(s, forbidden) {
			return ErrNotAllowed
		}
		return nil
	}
}

// OneOf validates that the string is one of the allowed values
func OneOf(allowed ...string) StringValidator {
	return func(s string) error {
		if slices.Contains(allowed, s) {
			return nil
		}
		return ErrNotInList
	}
}

// NoWhitespace validates that the string contains no whitespace characters
func NoWhitespace() StringValidator {
	return func(s string) error {
		for _, r := range s {
			if unicode.IsSpace(r) {
				return ErrHasWhitespace
			}
		}
		return nil
	}
}

// LowerCase validates that the string contains only lowercase letters (or non-letters)
func LowerCase() StringValidator {
	return func(s string) error {
		for _, r := range s {
			if unicode.IsLetter(r) && unicode.IsUpper(r) {
				return ErrNotLowerCase
			}
		}
		return nil
	}
}

// UpperCase validates that the string contains only uppercase letters (or non-letters)
func UpperCase() StringValidator {
	return func(s string) error {
		for _, r := range s {
			if unicode.IsLetter(r) && unicode.IsLower(r) {
				return ErrNotUpperCase
			}
		}
		return nil
	}
}
