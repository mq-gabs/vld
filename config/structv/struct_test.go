package structv

import (
	"errors"
	"strings"
	"testing"

	"github.com/mq-gabs/vld/config/validate"
)

type mockValidator[T any] struct {
	validate func(T) error
}

func (m mockValidator[T]) Validate(v T) error {
	return m.validate(v)
}

type User struct {
	Name string
	Age  int
}

func TestStructValidatorNullable(t *testing.T) {
	cfg := ConfigStruct[User]{
		Nullable: true,
		Validate: func(User) []Pair {
			return nil
		},
	}

	validator := cfg.Build()

	err := validator.Validate(nil)

	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
}

func TestStructValidatorNotNullable(t *testing.T) {
	cfg := ConfigStruct[User]{
		Nullable: false,
		Validate: func(User) []Pair {
			return nil
		},
	}

	validator := cfg.Build()

	err := validator.Validate(nil)

	if err == nil {
		t.Fatal("expected error but got nil")
	}
}

func TestStructValidatorSuccess(t *testing.T) {
	nameValidator := mockValidator[string]{
		validate: func(v string) error {
			if len(v) < 3 {
				return errors.New("invalid name")
			}

			return nil
		},
	}

	ageValidator := mockValidator[int]{
		validate: func(v int) error {
			if v < 18 {
				return errors.New("invalid age")
			}

			return nil
		},
	}

	cfg := ConfigStruct[User]{
		Validate: func(u User) []Pair {
			return []Pair{
				NewPair("name", u.Name, nameValidator),
				NewPair("age", u.Age, ageValidator),
			}
		},
	}

	validator := cfg.Build()

	err := validator.Validate(&User{
		Name: "John",
		Age:  25,
	})

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestStructValidatorSingleFailure(t *testing.T) {
	nameValidator := mockValidator[string]{
		validate: func(v string) error {
			if len(v) < 3 {
				return errors.New("too short")
			}

			return nil
		},
	}

	ageValidator := mockValidator[int]{
		validate: func(v int) error {
			return nil
		},
	}

	cfg := ConfigStruct[User]{
		Validate: func(u User) []Pair {
			return []Pair{
				NewPair("name", u.Name, nameValidator),
				NewPair("age", u.Age, ageValidator),
			}
		},
	}

	validator := cfg.Build()

	err := validator.Validate(&User{
		Name: "Jo",
		Age:  30,
	})

	if err == nil {
		t.Fatal("expected error but got nil")
	}

	if !strings.Contains(err.Error(), "name") {
		t.Fatalf("expected field name in error, got %v", err)
	}
}

func TestStructValidatorMultipleFailures(t *testing.T) {
	nameValidator := mockValidator[string]{
		validate: func(v string) error {
			if len(v) < 3 {
				return errors.New("too short")
			}

			return nil
		},
	}

	ageValidator := mockValidator[int]{
		validate: func(v int) error {
			if v < 18 {
				return errors.New("too young")
			}

			return nil
		},
	}

	cfg := ConfigStruct[User]{
		Validate: func(u User) []Pair {
			return []Pair{
				NewPair("name", u.Name, nameValidator),
				NewPair("age", u.Age, ageValidator),
			}
		},
	}

	validator := cfg.Build()

	err := validator.Validate(&User{
		Name: "Jo",
		Age:  15,
	})

	if err == nil {
		t.Fatal("expected error but got nil")
	}

	msg := err.Error()

	if !strings.Contains(msg, "name") {
		t.Fatalf("expected name error, got %v", err)
	}

	if !strings.Contains(msg, "age") {
		t.Fatalf("expected age error, got %v", err)
	}
}

func TestStructValidatorNilPairValidation(t *testing.T) {
	cfg := ConfigStruct[User]{
		Validate: func(User) []Pair {
			return []Pair{
				{},
			}
		},
	}

	validator := cfg.Build()

	err := validator.Validate(&User{
		Name: "John",
		Age:  30,
	})

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestNewPairPassesValueToValidator(t *testing.T) {
	called := false

	validator := mockValidator[string]{
		validate: func(v string) error {
			called = true

			if v != "john" {
				t.Fatalf("expected john, got %s", v)
			}

			return nil
		},
	}

	pair := NewPair("name", "john", validator)

	err := pair.validate()

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !called {
		t.Fatal("expected validator to be called")
	}
}

var _ validate.Validator[*User] = (*structValidator[User])(nil)
