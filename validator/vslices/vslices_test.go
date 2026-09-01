package vslices

import "testing"

func Test_Slice_MinLen_String(t *testing.T) {
	tests := []struct {
		name      string
		validator SliceValidator[string]
		value     []string
		valid     bool
	}{
		{
			name:      "min len string valid",
			validator: MinLen[string](4),
			value:     []string{"1", "2", "3", "4"},
			valid:     true,
		},
		{
			name:      "min len string invalid",
			validator: MinLen[string](8),
			value:     []string{"1", "2", "3", "4"},
			valid:     false,
		},
		{
			name:      "min len empty slice",
			validator: MinLen[string](0),
			value:     []string{},
			valid:     true,
		},
		{
			name:      "min len single item",
			validator: MinLen[string](1),
			value:     []string{"1", "2"},
			valid:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.validator(tt.value)

			if tt.valid && err != nil {
				t.Fatalf("expected valid but it's invalid: %v", err)
			}

			if !tt.valid && err == nil {
				t.Fatal("expected invalid but it's valid")
			}
		})
	}
}
