package schema

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Number(t *testing.T) {
	n := 10

	validSchema := Number[int]().Max(10).NonZero()
	invalidSchema := Number[int]().NonZero().Negative()

	err1 := validSchema.Validate(n)
	t.Log(err1)
	err2 := invalidSchema.Validate(n)
	t.Log(err2)

	assert.Nil(t, err1)
	assert.NotNil(t, err2)
}
