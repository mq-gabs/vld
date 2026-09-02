package validator_test

import (
	"errors"
	"testing"

	"github.com/mq-gabs/vld/validator"
	"github.com/mq-gabs/vld/validator/vstring"
)

func Test_When_String(t *testing.T) {
	tests := []struct {
		name     string
		validate func(string) error
		input    string
		valid    bool
	}{
		{
			name:     "apply validator",
			validate: validator.When(true, vstring.LowerCase()),
			input:    "HELLOWORLD",
			valid:    false,
		},
		{
			name:     "skip validator",
			validate: validator.When(false, vstring.LowerCase()),
			input:    "HELLOWORLD",
			valid:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.validate(tt.input)

			if tt.valid && err != nil {
				t.Fatalf("expected valid but got error: %v", err)
			}

			if !tt.valid && err == nil {
				t.Fatal("expected invalid but got valid")
			}
		})
	}
}

func Test_WhenFunc_String(t *testing.T) {
	validate := validator.WhenFunc(
		func(value string) bool { return value != "" },
		vstring.LowerCase(),
	)

	if err := validate("HELLO"); err == nil {
		t.Fatal("expected conditional validator to run")
	}

	if err := validate(""); err != nil {
		t.Fatalf("expected conditional validator to be skipped: %v", err)
	}
}

func Test_OneOf_String(t *testing.T) {
	validate := validator.OneOf(
		vstring.StartsWith("id-"),
		vstring.Email(),
	)

	if err := validate("id-123"); err != nil {
		t.Fatalf("expected first validator to satisfy OneOf: %v", err)
	}

	if err := validate("person@example.com"); err != nil {
		t.Fatalf("expected second validator to satisfy OneOf: %v", err)
	}

	if err := validate("invalid"); err == nil {
		t.Fatal("expected OneOf to fail when all validators fail")
	}
}

func Test_OneOf_PreservesValidatorErrors(t *testing.T) {
	firstErr := errors.New("first error")
	secondErr := errors.New("second error")
	validate := validator.OneOf(
		func(string) error { return firstErr },
		func(string) error { return secondErr },
	)

	err := validate("value")
	if err == nil {
		t.Fatal("expected OneOf to return an error")
	}
	if !errors.Is(err, firstErr) || !errors.Is(err, secondErr) {
		t.Fatalf("expected all validator errors to be preserved: %v", err)
	}
}
