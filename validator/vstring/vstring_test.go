package vstring

import (
	"errors"
	"strings"
	"testing"

	"github.com/mq-gabs/vld/validator"
)

func Test_String(t *testing.T) {
	validate := String("name",
		Required(),
		MaxLen(10),
		MinLen(3),
	)

	tests := []struct {
		value string
		valid bool
	}{
		{
			value: "john",
			valid: true,
		},
		{
			value: "johnas",
			valid: true,
		},
		{
			value: "mi",
			valid: false,
		},
		{
			value: "thisisverylongname",
			valid: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.value, func(t *testing.T) {
			err := validate(tt.value)

			if tt.valid && err != nil {
				t.Fatalf("expected valid but it's invalid: %v", err)
			}

			if !tt.valid && err == nil {
				t.Fatalf("expected invalid but it's valid")
			}
		})
	}
}

func Test_StringRequired(t *testing.T) {
	validate := Required()

	tests := []struct {
		name  string
		value string
		valid bool
	}{
		{
			name:  "non-empty string",
			value: "hello",
			valid: true,
		},
		{
			name:  "single character",
			value: "a",
			valid: true,
		},
		{
			name:  "empty string",
			value: "",
			valid: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validate(tt.value)

			if tt.valid && err != nil {
				t.Fatalf("expected valid but got error: %v", err)
			}

			if !tt.valid && err == nil {
				t.Fatalf("expected error but got none")
			}

			if !tt.valid && !errors.Is(err, validator.ErrValueIsRequired) {
				t.Fatalf("expected ErrValueIsRequired but got: %v", err)
			}
		})
	}
}

func Test_MinLen(t *testing.T) {
	validate := MinLen(5)

	tests := []struct {
		name  string
		value string
		valid bool
	}{
		{
			name:  "exactly min length",
			value: "hello",
			valid: true,
		},
		{
			name:  "longer than min",
			value: "hello world",
			valid: true,
		},
		{
			name:  "shorter than min",
			value: "hi",
			valid: false,
		},
		{
			name:  "empty string",
			value: "",
			valid: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validate(tt.value)

			if tt.valid && err != nil {
				t.Fatalf("expected valid but got error: %v", err)
			}

			if !tt.valid && err == nil {
				t.Fatalf("expected error but got none")
			}

			if !tt.valid && !errors.Is(err, validator.ErrMinLen) {
				t.Fatalf("expected ErrMinLen but got: %v", err)
			}
		})
	}
}

func Test_MaxLen(t *testing.T) {
	validate := MaxLen(10)

	tests := []struct {
		name  string
		value string
		valid bool
	}{
		{
			name:  "exactly max length",
			value: "0123456789",
			valid: true,
		},
		{
			name:  "shorter than max",
			value: "hello",
			valid: true,
		},
		{
			name:  "longer than max",
			value: "this is a very long string",
			valid: false,
		},
		{
			name:  "empty string",
			value: "",
			valid: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validate(tt.value)

			if tt.valid && err != nil {
				t.Fatalf("expected valid but got error: %v", err)
			}

			if !tt.valid && err == nil {
				t.Fatalf("expected error but got none")
			}

			if !tt.valid && !errors.Is(err, validator.ErrMaxLen) {
				t.Fatalf("expected ErrMaxLen but got: %v", err)
			}
		})
	}
}

func Test_Email(t *testing.T) {
	validate := Email()

	tests := []struct {
		name  string
		value string
		valid bool
	}{
		{
			name:  "valid email",
			value: "user@example.com",
			valid: true,
		},
		{
			name:  "valid email with dots",
			value: "user.name@example.co.uk",
			valid: true,
		},
		{
			name:  "valid email with numbers and special chars",
			value: "user+tag@example.com",
			valid: true,
		},
		{
			name:  "missing @",
			value: "userexample.com",
			valid: false,
		},
		{
			name:  "missing domain",
			value: "user@",
			valid: false,
		},
		{
			name:  "missing local part",
			value: "@example.com",
			valid: false,
		},
		{
			name:  "missing tld",
			value: "user@example",
			valid: false,
		},
		{
			name:  "empty string",
			value: "",
			valid: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validate(tt.value)

			if tt.valid && err != nil {
				t.Fatalf("expected valid but got error: %v", err)
			}

			if !tt.valid && err == nil {
				t.Fatalf("expected error but got none")
			}

			if !tt.valid && !errors.Is(err, ErrInvalidEmail) {
				t.Fatalf("expected ErrInvalidEmail but got: %v", err)
			}
		})
	}
}

func Test_UUID(t *testing.T) {
	validate := UUIDv4()

	tests := []struct {
		name  string
		value string
		valid bool
	}{
		{
			name:  "valid UUID v4",
			value: "550e8400-e29b-41d4-a716-446655440000",
			valid: true,
		},
		{
			name:  "valid UUID v4 uppercase",
			value: "550E8400-E29B-41D4-A716-446655440000",
			valid: true,
		},
		{
			name:  "valid UUID v1",
			value: "6ba7b810-9dad-11d1-80b4-00c04fd430c8",
			valid: true,
		},
		{
			name:  "invalid UUID - missing hyphen",
			value: "550e8400e29b41d4a716446655440000",
			valid: false,
		},
		{
			name:  "invalid UUID - wrong format",
			value: "550e8400-e29b-21d4-a716-446655440000a",
			valid: false,
		},
		{
			name:  "invalid UUID - empty",
			value: "",
			valid: false,
		},
		{
			name:  "invalid UUID - too short",
			value: "550e8400-e29b-41d4",
			valid: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validate(tt.value)

			if tt.valid && err != nil {
				t.Fatalf("expected valid but got error: %v", err)
			}

			if !tt.valid && err == nil {
				t.Fatalf("expected error but got none")
			}

			if !tt.valid && !errors.Is(err, ErrInvalidUUID) {
				t.Fatalf("expected ErrInvalidUUID but got: %v", err)
			}
		})
	}
}

func Test_Pattern(t *testing.T) {
	// Test for numeric pattern
	numericValidator := Pattern(`^\d+$`)

	tests := []struct {
		name    string
		pattern string
		value   string
		valid   bool
	}{
		{
			name:    "numeric pattern - all digits",
			pattern: `^\d+$`,
			value:   "12345",
			valid:   true,
		},
		{
			name:    "numeric pattern - with letters",
			pattern: `^\d+$`,
			value:   "123abc",
			valid:   false,
		},
		{
			name:    "hex color pattern",
			pattern: `^#[0-9A-Fa-f]{6}$`,
			value:   "#FF5733",
			valid:   true,
		},
		{
			name:    "hex color pattern - invalid",
			pattern: `^#[0-9A-Fa-f]{6}$`,
			value:   "#GGGGGG",
			valid:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			validator := Pattern(tt.pattern)
			err := validator(tt.value)

			if tt.valid && err != nil {
				t.Fatalf("expected valid but got error: %v", err)
			}

			if !tt.valid && err == nil {
				t.Fatalf("expected error but got none")
			}

			if !tt.valid && !errors.Is(err, ErrNoMatch) {
				t.Fatalf("expected ErrNoMatch but got: %v", err)
			}
		})
	}

	// Quick test for the prepared validator
	err := numericValidator("789")
	if err != nil {
		t.Fatalf("expected numeric pattern to accept '789' but got error: %v", err)
	}
}

func Test_Alphanumeric(t *testing.T) {
	validate := Alphanumeric()

	tests := []struct {
		name  string
		value string
		valid bool
	}{
		{
			name:  "only letters",
			value: "abcDEF",
			valid: true,
		},
		{
			name:  "only digits",
			value: "123456",
			valid: true,
		},
		{
			name:  "letters and digits",
			value: "abc123DEF456",
			valid: true,
		},
		{
			name:  "with space",
			value: "abc 123",
			valid: false,
		},
		{
			name:  "with special char",
			value: "abc@123",
			valid: false,
		},
		{
			name:  "empty string",
			value: "",
			valid: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validate(tt.value)

			if tt.valid && err != nil {
				t.Fatalf("expected valid but got error: %v", err)
			}

			if !tt.valid && err == nil {
				t.Fatalf("expected error but got none")
			}

			if !tt.valid && !errors.Is(err, ErrNotAlphaNum) {
				t.Fatalf("expected ErrNotAlphaNum but got: %v", err)
			}
		})
	}
}

func Test_StartsWith(t *testing.T) {
	validator := StartsWith("hello")

	tests := []struct {
		name  string
		value string
		valid bool
	}{
		{
			name:  "exact match",
			value: "hello",
			valid: true,
		},
		{
			name:  "with suffix",
			value: "hello world",
			valid: true,
		},
		{
			name:  "different prefix",
			value: "goodbye world",
			valid: false,
		},
		{
			name:  "partially matching",
			value: "helloworld",
			valid: true,
		},
		{
			name:  "empty string",
			value: "",
			valid: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validator(tt.value)

			if tt.valid && err != nil {
				t.Fatalf("expected valid but got error: %v", err)
			}

			if !tt.valid && err == nil {
				t.Fatalf("expected error but got none")
			}

			if !tt.valid && !errors.Is(err, ErrInvalidStart) {
				t.Fatalf("expected ErrInvalidStart but got: %v", err)
			}
		})
	}
}

func Test_EndsWith(t *testing.T) {
	validator := EndsWith(".com")

	tests := []struct {
		name  string
		value string
		valid bool
	}{
		{
			name:  "exact match",
			value: ".com",
			valid: true,
		},
		{
			name:  "with prefix",
			value: "example.com",
			valid: true,
		},
		{
			name:  "different suffix",
			value: "example.org",
			valid: false,
		},
		{
			name:  "partial match",
			value: "website.co",
			valid: false,
		},
		{
			name:  "empty string",
			value: "",
			valid: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validator(tt.value)

			if tt.valid && err != nil {
				t.Fatalf("expected valid but got error: %v", err)
			}

			if !tt.valid && err == nil {
				t.Fatalf("expected error but got none")
			}

			if !tt.valid && !errors.Is(err, ErrInvalidEnd) {
				t.Fatalf("expected ErrInvalidEnd but got: %v", err)
			}
		})
	}
}

func Test_NotContains(t *testing.T) {
	validator := NotContains("<>&\"'")

	tests := []struct {
		name  string
		value string
		valid bool
	}{
		{
			name:  "clean string",
			value: "hello world",
			valid: true,
		},
		{
			name:  "with forbidden char <",
			value: "hello <world>",
			valid: false,
		},
		{
			name:  "with forbidden char &",
			value: "hello&world",
			valid: false,
		},
		{
			name:  "with forbidden char '",
			value: "it's",
			valid: false,
		},
		{
			name:  "empty string",
			value: "",
			valid: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validator(tt.value)

			if tt.valid && err != nil {
				t.Fatalf("expected valid but got error: %v", err)
			}

			if !tt.valid && err == nil {
				t.Fatalf("expected error but got none")
			}

			if !tt.valid && !errors.Is(err, ErrNotAllowed) {
				t.Fatalf("expected ErrNotAllowed but got: %v", err)
			}
		})
	}
}

func Test_OneOf(t *testing.T) {
	validator := OneOf("red", "green", "blue")

	tests := []struct {
		name  string
		value string
		valid bool
	}{
		{
			name:  "first option",
			value: "red",
			valid: true,
		},
		{
			name:  "middle option",
			value: "green",
			valid: true,
		},
		{
			name:  "last option",
			value: "blue",
			valid: true,
		},
		{
			name:  "not in list",
			value: "yellow",
			valid: false,
		},
		{
			name:  "case sensitive",
			value: "RED",
			valid: false,
		},
		{
			name:  "empty string",
			value: "",
			valid: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validator(tt.value)

			if tt.valid && err != nil {
				t.Fatalf("expected valid but got error: %v", err)
			}

			if !tt.valid && err == nil {
				t.Fatalf("expected error but got none")
			}

			if !tt.valid && !errors.Is(err, ErrNotInList) {
				t.Fatalf("expected ErrNotInList but got: %v", err)
			}
		})
	}
}

func Test_NoWhitespace(t *testing.T) {
	validator := NoWhitespace()

	tests := []struct {
		name  string
		value string
		valid bool
	}{
		{
			name:  "no whitespace",
			value: "helloworld",
			valid: true,
		},
		{
			name:  "with space",
			value: "hello world",
			valid: false,
		},
		{
			name:  "with tab",
			value: "hello\tworld",
			valid: false,
		},
		{
			name:  "with newline",
			value: "hello\nworld",
			valid: false,
		},
		{
			name:  "empty string",
			value: "",
			valid: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validator(tt.value)

			if tt.valid && err != nil {
				t.Fatalf("expected valid but got error: %v", err)
			}

			if !tt.valid && err == nil {
				t.Fatalf("expected error but got none")
			}

			if !tt.valid && !errors.Is(err, ErrHasWhitespace) {
				t.Fatalf("expected ErrHasWhitespace but got: %v", err)
			}
		})
	}
}

func Test_LowerCase(t *testing.T) {
	validator := LowerCase()

	tests := []struct {
		name  string
		value string
		valid bool
	}{
		{
			name:  "all lowercase",
			value: "hello",
			valid: true,
		},
		{
			name:  "lowercase with numbers",
			value: "hello123",
			valid: true,
		},
		{
			name:  "mixed case",
			value: "helloWorld",
			valid: false,
		},
		{
			name:  "all uppercase",
			value: "HELLO",
			valid: false,
		},
		{
			name:  "with special chars",
			value: "hello-world",
			valid: true,
		},
		{
			name:  "empty string",
			value: "",
			valid: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validator(tt.value)

			if tt.valid && err != nil {
				t.Fatalf("expected valid but got error: %v", err)
			}

			if !tt.valid && err == nil {
				t.Fatalf("expected error but got none")
			}

			if !tt.valid && !errors.Is(err, ErrNotLowerCase) {
				t.Fatalf("expected ErrNotLowerCase but got: %v", err)
			}
		})
	}
}

func Test_UpperCase(t *testing.T) {
	validator := UpperCase()

	tests := []struct {
		name  string
		value string
		valid bool
	}{
		{
			name:  "all uppercase",
			value: "HELLO",
			valid: true,
		},
		{
			name:  "uppercase with numbers",
			value: "HELLO123",
			valid: true,
		},
		{
			name:  "mixed case",
			value: "HelloWorld",
			valid: false,
		},
		{
			name:  "all lowercase",
			value: "hello",
			valid: false,
		},
		{
			name:  "with special chars",
			value: "HELLO-WORLD",
			valid: true,
		},
		{
			name:  "empty string",
			value: "",
			valid: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validator(tt.value)

			if tt.valid && err != nil {
				t.Fatalf("expected valid but got error: %v", err)
			}

			if !tt.valid && err == nil {
				t.Fatalf("expected error but got none")
			}

			if !tt.valid && !errors.Is(err, ErrNotUpperCase) {
				t.Fatalf("expected ErrNotUpperCase but got: %v", err)
			}
		})
	}
}

// Test combining multiple validators
func Test_StringWithMultipleValidators(t *testing.T) {
	validator := String("email_field",
		Required(),
		Email(),
		NoWhitespace(),
		MaxLen(100),
	)

	tests := []struct {
		name  string
		value string
		valid bool
	}{
		{
			name:  "valid email",
			value: "user@example.com",
			valid: true,
		},
		{
			name:  "empty string",
			value: "",
			valid: false,
		},
		{
			name:  "invalid email format",
			value: "not-an-email",
			valid: false,
		},
		{
			name:  "email with space",
			value: "user @example.com",
			valid: false,
		},
		{
			name:  "email too long",
			value: "verylongemailaddress" + strings.Repeat("a", 100) + "@example.com",
			valid: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validator(tt.value)

			if tt.valid && err != nil {
				t.Fatalf("expected valid but got error: %v", err)
			}

			if !tt.valid && err == nil {
				t.Fatalf("expected error but got none")
			}
		})
	}
}
