package validator_test

import (
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
