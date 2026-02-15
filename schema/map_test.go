package schema

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Map(t *testing.T) {
	m := map[string]string{
		"name": "John doe",
		"city": "London",
	}

	validSchema := Map[string, string]().LengthMax(10)
	invalidSchema := Map[string, string]().LengthMin(10)

	err1 := validSchema.Validate(m)
	t.Log(err1)
	err2 := invalidSchema.Validate(m)
	t.Log(err2)

	assert.Nil(t, err1)
	assert.NotNil(t, err2)
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

	err1 := validSchema.Validate(m)
	t.Log(err1)
	err2 := invalidSchema.Validate(m)
	t.Log(err2)

	assert.Nil(t, err1)
	assert.NotNil(t, err2)
}
