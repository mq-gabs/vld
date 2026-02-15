package schema

import (
	"testing"
)

func Test_Slice(t *testing.T) {
	nums := []int{1, 2, 3, 4}

	validSchema := Slice[int]().MinLength(2).MaxLength(8)
	invalidSchema := Slice[int]().MinLength(6).MaxLength(10)

	if err := validSchema.Validate(nums); err != nil {
		t.Error(err)
	}
	if err := invalidSchema.Validate(nums); err != nil {
		t.Error(err)
	}
}
