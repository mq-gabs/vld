package schema

import (
	"errors"
	"testing"
)

func Test_Generic(t *testing.T) {
	n := 2
	name := "john"
	flag := true

	genericSchema := Generic().Custom(func(v any) error {
		switch v.(type) {
		case string, int, bool:
			return nil
		default:
			return errors.New("invalid type")
		}
	})

	if err := genericSchema.Validate(n); err != nil {
		t.Error(err)
	}
	if err := genericSchema.Validate(name); err != nil {
		t.Error(err)
	}
	if err := genericSchema.Validate(flag); err != nil {
		t.Error(err)
	}
}
