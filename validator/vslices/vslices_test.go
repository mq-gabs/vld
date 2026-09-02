package vslices

import (
	"testing"

	"github.com/mq-gabs/vld/validator"
	"github.com/mq-gabs/vld/validator/vnumber"
	"github.com/mq-gabs/vld/validator/vstring"
)

func Test_Slice_GroupedValidator(t *testing.T) {
	validator := Slice("items",
		Required[string](),
		MinLen[string](1),
		MaxLen[string](5),
	)

	tests := []struct {
		name  string
		value []string
		valid bool
	}{
		{
			name:  "valid slice",
			value: []string{"a", "b"},
			valid: true,
		},
		{
			name:  "empty slice",
			value: []string{},
			valid: false,
		},
		{
			name:  "slice with 5 items",
			value: []string{"a", "b", "c", "d", "e"},
			valid: true,
		},
		{
			name:  "slice with 6 items (exceeds max)",
			value: []string{"a", "b", "c", "d", "e", "f"},
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
	validator := Required[string]()

	tests := []struct {
		name  string
		value []string
		valid bool
	}{
		{
			name:  "non-empty slice",
			value: []string{"a"},
			valid: true,
		},
		{
			name:  "empty slice",
			value: []string{},
			valid: false,
		},
		{
			name:  "nil slice is empty",
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

func Test_MinLen_String(t *testing.T) {
	validator := MinLen[string](2)

	tests := []struct {
		name  string
		value []string
		valid bool
	}{
		{
			name:  "slice with 2 items",
			value: []string{"a", "b"},
			valid: true,
		},
		{
			name:  "slice with 3 items",
			value: []string{"a", "b", "c"},
			valid: true,
		},
		{
			name:  "slice with 1 item",
			value: []string{"a"},
			valid: false,
		},
		{
			name:  "empty slice",
			value: []string{},
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

func Test_MaxLen_Int(t *testing.T) {
	validator := MaxLen[int](3)

	tests := []struct {
		name  string
		value []int
		valid bool
	}{
		{
			name:  "slice with 1 item",
			value: []int{1},
			valid: true,
		},
		{
			name:  "slice with 3 items",
			value: []int{1, 2, 3},
			valid: true,
		},
		{
			name:  "slice with 4 items",
			value: []int{1, 2, 3, 4},
			valid: false,
		},
		{
			name:  "empty slice",
			value: []int{},
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
	validator := ExactLen[string](3)

	tests := []struct {
		name  string
		value []string
		valid bool
	}{
		{
			name:  "slice with exact length",
			value: []string{"a", "b", "c"},
			valid: true,
		},
		{
			name:  "slice with fewer items",
			value: []string{"a", "b"},
			valid: false,
		},
		{
			name:  "slice with more items",
			value: []string{"a", "b", "c", "d"},
			valid: false,
		},
		{
			name:  "empty slice",
			value: []string{},
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
	validator := InRange[int](2, 4)

	tests := []struct {
		name  string
		value []int
		valid bool
	}{
		{
			name:  "slice with 2 items (min)",
			value: []int{1, 2},
			valid: true,
		},
		{
			name:  "slice with 3 items",
			value: []int{1, 2, 3},
			valid: true,
		},
		{
			name:  "slice with 4 items (max)",
			value: []int{1, 2, 3, 4},
			valid: true,
		},
		{
			name:  "slice with 1 item (below range)",
			value: []int{1},
			valid: false,
		},
		{
			name:  "slice with 5 items (above range)",
			value: []int{1, 2, 3, 4, 5},
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

func Test_Each_String(t *testing.T) {
	minLenValidator := validator.GenericValidator[string](vstring.MinLen(2))
	validator := Each[string](minLenValidator)

	tests := []struct {
		name  string
		value []string
		valid bool
	}{
		{
			name:  "all items meet min length",
			value: []string{"ab", "cde", "fg"},
			valid: true,
		},
		{
			name:  "one item too short",
			value: []string{"ab", "a", "cde"},
			valid: false,
		},
		{
			name:  "all items too short",
			value: []string{"a", "b"},
			valid: false,
		},
		{
			name:  "empty slice",
			value: []string{},
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

func Test_Each_Number(t *testing.T) {
	isPositiveValidator := validator.GenericValidator[int](vnumber.IsPositive[int]())
	validator := Each[int](isPositiveValidator)

	tests := []struct {
		name  string
		value []int
		valid bool
	}{
		{
			name:  "all items positive",
			value: []int{1, 2, 3},
			valid: true,
		},
		{
			name:  "one item not positive",
			value: []int{1, 0, 3},
			valid: false,
		},
		{
			name:  "negative item",
			value: []int{1, -5},
			valid: false,
		},
		{
			name:  "empty slice",
			value: []int{},
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

func Test_Contains(t *testing.T) {
	validator := Contains[string]("target")

	tests := []struct {
		name  string
		value []string
		valid bool
	}{
		{
			name:  "slice contains item",
			value: []string{"a", "target", "c"},
			valid: true,
		},
		{
			name:  "slice does not contain item",
			value: []string{"a", "b", "c"},
			valid: false,
		},
		{
			name:  "empty slice",
			value: []string{},
			valid: false,
		},
		{
			name:  "single item match",
			value: []string{"target"},
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

func Test_NotContains(t *testing.T) {
	validator := NotContains[int](99)

	tests := []struct {
		name  string
		value []int
		valid bool
	}{
		{
			name:  "slice does not contain forbidden item",
			value: []int{1, 2, 3},
			valid: true,
		},
		{
			name:  "slice contains forbidden item",
			value: []int{1, 99, 3},
			valid: false,
		},
		{
			name:  "empty slice",
			value: []int{},
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

func Test_NoDuplicates(t *testing.T) {
	validator := NoDuplicates[string]()

	tests := []struct {
		name  string
		value []string
		valid bool
	}{
		{
			name:  "slice with no duplicates",
			value: []string{"a", "b", "c"},
			valid: true,
		},
		{
			name:  "slice with duplicates",
			value: []string{"a", "b", "a"},
			valid: false,
		},
		{
			name:  "slice with consecutive duplicates",
			value: []string{"a", "a", "b"},
			valid: false,
		},
		{
			name:  "empty slice",
			value: []string{},
			valid: true,
		},
		{
			name:  "single item",
			value: []string{"a"},
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

func Test_ComplexSlice_Validation(t *testing.T) {
	// Complex validation: slice must have 1-4 items, all strings at least 2 chars, no duplicates
	minLenValidator := validator.GenericValidator[string](vstring.MinLen(2))
	validator := Slice("words",
		InRange[string](1, 4),
		Each[string](minLenValidator),
		NoDuplicates[string](),
	)

	tests := []struct {
		name  string
		value []string
		valid bool
	}{
		{
			name:  "valid slice",
			value: []string{"ab", "cd", "ef"},
			valid: true,
		},
		{
			name:  "empty slice",
			value: []string{},
			valid: false,
		},
		{
			name:  "item too short",
			value: []string{"ab", "c", "de"},
			valid: false,
		},
		{
			name:  "duplicate items",
			value: []string{"ab", "cd", "ab"},
			valid: false,
		},
		{
			name:  "too many items",
			value: []string{"ab", "cd", "ef", "gh", "ij"},
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
