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

	validSchema := Map[string, string]().MaxLength(10)
	invalidSchema := Map[string, string]().MinLength(10)

	err1 := validSchema.Validate(m)
	t.Log(err1)
	err2 := invalidSchema.Validate(m)
	t.Log(err2)

	assert.Nil(t, err1)
	assert.NotNil(t, err2)
}
