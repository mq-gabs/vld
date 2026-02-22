package schema

import (
	"testing"
)

func Test_Slice(t *testing.T) {
	nums := []int{1, 2, 3, 4}

	validSchema := Slice[int]().LengthMin(2).LengthMax(8)
	invalidSchema := Slice[int]().LengthMin(6).LengthMax(10)

	if err := validSchema.Validate(nums); err != nil {
		t.Error(err)
	}
	if err := invalidSchema.Validate(nums); err == nil {
		t.Error(errExpectedError)
	}
}
