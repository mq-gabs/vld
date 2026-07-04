package schema

import (
	"errors"
	"testing"
)

func Test_Schema_Clone_String(t *testing.T) {
	base := String().LengthMax(4)
	extended := base.Clone().Custom(func(s string) error {
		if s == "abcd" {
			return errors.New("invalid")
		}

		return nil
	})

	tests := []struct {
		value string
		scm   SchemaString
		valid bool
	}{
		{"abcd", base, true},
		{"abcde", base, false},
		{"abcd", extended, false},
		{"xyzk", extended, true},
	}

	for _, tt := range tests {
		t.Run(tt.value, func(t *testing.T) {
			err := tt.scm.Validate(tt.value)

			if !(err == nil == tt.valid) {
				t.Errorf("expected-valid=%v got=%v", tt.valid, err)
			}
		})
	}
}
