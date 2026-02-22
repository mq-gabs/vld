package schema

import (
	"testing"
)

func Test_Map(t *testing.T) {
	m := map[string]string{
		"name": "John doe",
		"city": "London",
	}

	validSchema := Map[string, string]().LengthMax(10)
	invalidSchema := Map[string, string]().LengthMin(10)

	if err := validSchema.Validate(m); err != nil {
		t.Error(err)
	}
	if err := invalidSchema.Validate(m); err == nil {
		t.Error(errExpectedError)
	}
}

func Test_MapChild(t *testing.T) {
	m := map[string]string{
		"name": "John Doe",
		"city": "Tokyo",
	}

	validStrSchema := String().LengthMax(10)
	invalidStrSchema := String().LengthMin(10)

	validSchema := Map[string, string]().Child(validStrSchema)
	invalidSchema := Map[string, string]().Child(invalidStrSchema)

	if err := validSchema.Validate(m); err != nil {
		t.Error(err)
	}
	if err := invalidSchema.Validate(m); err == nil {
		t.Error(errExpectedError)
	}
}
