package schema

import (
	"testing"
)

func Test_Number(t *testing.T) {
	n := 10

	validSchema := Number[int]().Max(10).NonZero()
	invalidSchema := Number[int]().NonZero().Negative()

	if err := validSchema.Validate(n); err != nil {
		t.Error(err)
	}
	if err := invalidSchema.Validate(n); err == nil {
		t.Error(errExpectedError)
	}
}

func Test_NumberOptional(t *testing.T) {
	s := Number[int]().Min(10).Optional()

	err := s.Validate(0)
	if err != nil {
		t.Fatalf("expected no error, but got %v", err)
	}
}
