package vmap

import (
	"testing"

	"github.com/mq-gabs/vld/validator/vnumber"
	"github.com/mq-gabs/vld/validator/vstring"
)

func Test_Map_GroupedValidator(t *testing.T) {
	validator := Map("config",
		Required[string, string](),
		MinLen[string, string](1),
		MaxLen[string, string](5),
	)

	tests := []struct {
		name  string
		value map[string]string
		valid bool
	}{
		{
			name:  "valid map",
			value: map[string]string{"key": "value"},
			valid: true,
		},
		{
			name:  "empty map",
			value: map[string]string{},
			valid: false,
		},
		{
			name:  "map with 5 items",
			value: map[string]string{"a": "1", "b": "2", "c": "3", "d": "4", "e": "5"},
			valid: true,
		},
		{
			name:  "map with 6 items (exceeds max)",
			value: map[string]string{"a": "1", "b": "2", "c": "3", "d": "4", "e": "5", "f": "6"},
			valid: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validator(tt.value)

			if tt.valid && err != nil {
				t.Fatalf("expected valid but got error: %v", err)
			}

			if !tt.valid && err == nil {
				t.Fatalf("expected invalid but got no error")
			}
		})
	}
}

func Test_Required(t *testing.T) {
	validator := Required[string, string]()

	tests := []struct {
		name  string
		value map[string]string
		valid bool
	}{
		{
			name:  "non-empty map",
			value: map[string]string{"key": "value"},
			valid: true,
		},
		{
			name:  "empty map",
			value: map[string]string{},
			valid: false,
		},
		{
			name:  "nil map is empty",
			value: nil,
			valid: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validator(tt.value)

			if tt.valid && err != nil {
				t.Fatalf("expected valid but got error: %v", err)
			}

			if !tt.valid && err == nil {
				t.Fatalf("expected invalid but got no error")
			}
		})
	}
}

func Test_MinLen(t *testing.T) {
	validator := MinLen[string, string](2)

	tests := []struct {
		name  string
		value map[string]string
		valid bool
	}{
		{
			name:  "map with 2 items",
			value: map[string]string{"a": "1", "b": "2"},
			valid: true,
		},
		{
			name:  "map with 3 items",
			value: map[string]string{"a": "1", "b": "2", "c": "3"},
			valid: true,
		},
		{
			name:  "map with 1 item",
			value: map[string]string{"a": "1"},
			valid: false,
		},
		{
			name:  "empty map",
			value: map[string]string{},
			valid: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validator(tt.value)

			if tt.valid && err != nil {
				t.Fatalf("expected valid but got error: %v", err)
			}

			if !tt.valid && err == nil {
				t.Fatalf("expected invalid but got no error")
			}
		})
	}
}

func Test_MaxLen(t *testing.T) {
	validator := MaxLen[string, string](3)

	tests := []struct {
		name  string
		value map[string]string
		valid bool
	}{
		{
			name:  "map with 1 item",
			value: map[string]string{"a": "1"},
			valid: true,
		},
		{
			name:  "map with 3 items",
			value: map[string]string{"a": "1", "b": "2", "c": "3"},
			valid: true,
		},
		{
			name:  "map with 4 items",
			value: map[string]string{"a": "1", "b": "2", "c": "3", "d": "4"},
			valid: false,
		},
		{
			name:  "empty map",
			value: map[string]string{},
			valid: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validator(tt.value)

			if tt.valid && err != nil {
				t.Fatalf("expected valid but got error: %v", err)
			}

			if !tt.valid && err == nil {
				t.Fatalf("expected invalid but got no error")
			}
		})
	}
}

func Test_ExactLen(t *testing.T) {
	validator := ExactLen[string, string](3)

	tests := []struct {
		name  string
		value map[string]string
		valid bool
	}{
		{
			name:  "map with exact length",
			value: map[string]string{"a": "1", "b": "2", "c": "3"},
			valid: true,
		},
		{
			name:  "map with fewer items",
			value: map[string]string{"a": "1", "b": "2"},
			valid: false,
		},
		{
			name:  "map with more items",
			value: map[string]string{"a": "1", "b": "2", "c": "3", "d": "4"},
			valid: false,
		},
		{
			name:  "empty map",
			value: map[string]string{},
			valid: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validator(tt.value)

			if tt.valid && err != nil {
				t.Fatalf("expected valid but got error: %v", err)
			}

			if !tt.valid && err == nil {
				t.Fatalf("expected invalid but got no error")
			}
		})
	}
}

func Test_InRange(t *testing.T) {
	validator := InRange[string, string](2, 4)

	tests := []struct {
		name  string
		value map[string]string
		valid bool
	}{
		{
			name:  "map with 2 items (min)",
			value: map[string]string{"a": "1", "b": "2"},
			valid: true,
		},
		{
			name:  "map with 3 items",
			value: map[string]string{"a": "1", "b": "2", "c": "3"},
			valid: true,
		},
		{
			name:  "map with 4 items (max)",
			value: map[string]string{"a": "1", "b": "2", "c": "3", "d": "4"},
			valid: true,
		},
		{
			name:  "map with 1 item (below range)",
			value: map[string]string{"a": "1"},
			valid: false,
		},
		{
			name:  "map with 5 items (above range)",
			value: map[string]string{"a": "1", "b": "2", "c": "3", "d": "4", "e": "5"},
			valid: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validator(tt.value)

			if tt.valid && err != nil {
				t.Fatalf("expected valid but got error: %v", err)
			}

			if !tt.valid && err == nil {
				t.Fatalf("expected invalid but got no error")
			}
		})
	}
}

func Test_ContainsKey(t *testing.T) {
	validator := ContainsKey[string, string]("name")

	tests := []struct {
		name  string
		value map[string]string
		valid bool
	}{
		{
			name:  "map contains key",
			value: map[string]string{"name": "John", "age": "30"},
			valid: true,
		},
		{
			name:  "map does not contain key",
			value: map[string]string{"age": "30", "city": "NY"},
			valid: false,
		},
		{
			name:  "empty map",
			value: map[string]string{},
			valid: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validator(tt.value)

			if tt.valid && err != nil {
				t.Fatalf("expected valid but got error: %v", err)
			}

			if !tt.valid && err == nil {
				t.Fatalf("expected invalid but got no error")
			}
		})
	}
}

func Test_NotContainsKey(t *testing.T) {
	validator := NotContainsKey[string, string]("password")

	tests := []struct {
		name  string
		value map[string]string
		valid bool
	}{
		{
			name:  "map does not contain forbidden key",
			value: map[string]string{"name": "John", "age": "30"},
			valid: true,
		},
		{
			name:  "map contains forbidden key",
			value: map[string]string{"name": "John", "password": "secret"},
			valid: false,
		},
		{
			name:  "empty map",
			value: map[string]string{},
			valid: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validator(tt.value)

			if tt.valid && err != nil {
				t.Fatalf("expected valid but got error: %v", err)
			}

			if !tt.valid && err == nil {
				t.Fatalf("expected invalid but got no error")
			}
		})
	}
}

func Test_HasKeys(t *testing.T) {
	validator := HasKeys[string, string]("name", "age", "email")

	tests := []struct {
		name  string
		value map[string]string
		valid bool
	}{
		{
			name:  "map has all required keys",
			value: map[string]string{"name": "John", "age": "30", "email": "john@example.com", "city": "NY"},
			valid: true,
		},
		{
			name:  "map missing one key",
			value: map[string]string{"name": "John", "age": "30"},
			valid: false,
		},
		{
			name:  "map missing multiple keys",
			value: map[string]string{"name": "John"},
			valid: false,
		},
		{
			name:  "empty map",
			value: map[string]string{},
			valid: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validator(tt.value)

			if tt.valid && err != nil {
				t.Fatalf("expected valid but got error: %v", err)
			}

			if !tt.valid && err == nil {
				t.Fatalf("expected invalid but got no error")
			}
		})
	}
}

func Test_NotHasKeys(t *testing.T) {
	validator := NotHasKeys[string, string]("password", "secret", "token")

	tests := []struct {
		name  string
		value map[string]string
		valid bool
	}{
		{
			name:  "map does not contain forbidden keys",
			value: map[string]string{"name": "John", "age": "30", "email": "john@example.com"},
			valid: true,
		},
		{
			name:  "map contains one forbidden key",
			value: map[string]string{"name": "John", "password": "secret"},
			valid: false,
		},
		{
			name:  "map contains multiple forbidden keys",
			value: map[string]string{"name": "John", "password": "pass", "token": "xyz"},
			valid: false,
		},
		{
			name:  "empty map",
			value: map[string]string{},
			valid: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validator(tt.value)

			if tt.valid && err != nil {
				t.Fatalf("expected valid but got error: %v", err)
			}

			if !tt.valid && err == nil {
				t.Fatalf("expected invalid but got no error")
			}
		})
	}
}

func Test_EachValue_String(t *testing.T) {
	// Wrapper to convert StringValidator to GenericValidator
	validator := EachValue[string, string](vstring.MinLen(3))

	tests := []struct {
		name  string
		value map[string]string
		valid bool
	}{
		{
			name:  "all values meet min length",
			value: map[string]string{"key1": "value1", "key2": "hello"},
			valid: true,
		},
		{
			name:  "one value too short",
			value: map[string]string{"key1": "value1", "key2": "hi"},
			valid: false,
		},
		{
			name:  "all values too short",
			value: map[string]string{"a": "ab", "b": "xy"},
			valid: false,
		},
		{
			name:  "empty map",
			value: map[string]string{},
			valid: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validator(tt.value)

			if tt.valid && err != nil {
				t.Fatalf("expected valid but got error: %v", err)
			}

			if !tt.valid && err == nil {
				t.Fatalf("expected invalid but got no error")
			}
		})
	}
}

func Test_EachValue_Number(t *testing.T) {
	// Wrapper to convert NumberValidator to GenericValidator
	validator := EachValue[string, int](vnumber.IsPositive[int]())

	tests := []struct {
		name  string
		value map[string]int
		valid bool
	}{
		{
			name:  "all values positive",
			value: map[string]int{"a": 1, "b": 2, "c": 10},
			valid: true,
		},
		{
			name:  "one value not positive",
			value: map[string]int{"a": 1, "b": 0, "c": 10},
			valid: false,
		},
		{
			name:  "negative value",
			value: map[string]int{"a": 1, "b": -5},
			valid: false,
		},
		{
			name:  "empty map",
			value: map[string]int{},
			valid: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validator(tt.value)

			if tt.valid && err != nil {
				t.Fatalf("expected valid but got error: %v", err)
			}

			if !tt.valid && err == nil {
				t.Fatalf("expected invalid but got no error")
			}
		})
	}
}

func Test_EachKey(t *testing.T) {
	// Wrapper to convert StringValidator to GenericValidator
	validator := EachKey[string, string](vstring.MinLen(2))

	tests := []struct {
		name  string
		value map[string]string
		valid bool
	}{
		{
			name:  "all keys meet min length",
			value: map[string]string{"key1": "v1", "key2": "v2"},
			valid: true,
		},
		{
			name:  "one key too short",
			value: map[string]string{"k": "v1", "key2": "v2"},
			valid: false,
		},
		{
			name:  "empty map",
			value: map[string]string{},
			valid: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validator(tt.value)

			if tt.valid && err != nil {
				t.Fatalf("expected valid but got error: %v", err)
			}

			if !tt.valid && err == nil {
				t.Fatalf("expected invalid but got no error")
			}
		})
	}
}

func Test_Each_Alias(t *testing.T) {
	// Test that Each is an alias for EachValue
	// Wrapper to convert StringValidator to GenericValidator
	validator := Each[string, string](vstring.MaxLen(10))

	tests := []struct {
		name  string
		value map[string]string
		valid bool
	}{
		{
			name:  "all values within max length",
			value: map[string]string{"key1": "value", "key2": "test"},
			valid: true,
		},
		{
			name:  "one value exceeds max length",
			value: map[string]string{"key1": "value", "key2": "this is a very long string"},
			valid: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validator(tt.value)

			if tt.valid && err != nil {
				t.Fatalf("expected valid but got error: %v", err)
			}

			if !tt.valid && err == nil {
				t.Fatalf("expected invalid but got no error")
			}
		})
	}
}

func Test_ComplexMap_Validation(t *testing.T) {
	// Complex validation: map must have 1-3 items, contain "id" and "name" keys,
	// and all values must be at least 2 characters
	// Wrapper to convert StringValidator to GenericValidator
	validator := Map("user",
		InRange[string, string](1, 3),
		HasKeys[string, string]("id", "name"),
		EachValue[string, string](vstring.MinLen(2)),
	)

	tests := []struct {
		name  string
		value map[string]string
		valid bool
	}{
		{
			name:  "valid user map",
			value: map[string]string{"id": "123", "name": "John"},
			valid: true,
		},
		{
			name:  "missing required key",
			value: map[string]string{"id": "123"},
			valid: false,
		},
		{
			name:  "value too short",
			value: map[string]string{"id": "1", "name": "John"},
			valid: false,
		},
		{
			name:  "too many items",
			value: map[string]string{"id": "123", "name": "John", "email": "john@example.com", "city": "NY"},
			valid: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validator(tt.value)

			if tt.valid && err != nil {
				t.Fatalf("expected valid but got error: %v", err)
			}

			if !tt.valid && err == nil {
				t.Fatalf("expected invalid but got no error")
			}
		})
	}
}
