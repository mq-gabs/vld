package vnumber

import (
	"strings"
	"testing"
)

func Test_Number_Int_GroupedValidator(t *testing.T) {
	validator := Number("age",
		IsPositive[int](),
		MaxValue(150),
	)

	tests := []struct {
		name  string
		value int
		valid bool
	}{
		{
			name:  "valid age",
			value: 25,
			valid: true,
		},
		{
			name:  "minimum valid (1)",
			value: 1,
			valid: true,
		},
		{
			name:  "maximum valid",
			value: 150,
			valid: true,
		},
		{
			name:  "negative value",
			value: -5,
			valid: false,
		},
		{
			name:  "above maximum",
			value: 200,
			valid: false,
		},
		{
			name:  "zero with IsPositive",
			value: 0,
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

func Test_MinValue_Int(t *testing.T) {
	validator := MinValue(10)

	tests := []struct {
		name  string
		value int
		valid bool
	}{
		{"above minimum", 15, true},
		{"equal to minimum", 10, true},
		{"below minimum", 5, false},
		{"negative below minimum", -10, false},
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

func Test_MaxValue_Int(t *testing.T) {
	validator := MaxValue(100)

	tests := []struct {
		name  string
		value int
		valid bool
	}{
		{"below maximum", 50, true},
		{"equal to maximum", 100, true},
		{"above maximum", 150, false},
		{"way above maximum", 1000, false},
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

func Test_InRange_Int(t *testing.T) {
	validator := InRange(10, 100)

	tests := []struct {
		name  string
		value int
		valid bool
	}{
		{"below range", 5, false},
		{"at minimum", 10, true},
		{"inside range", 50, true},
		{"at maximum", 100, true},
		{"above range", 150, false},
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

func Test_IsPositive_Int(t *testing.T) {
	validator := IsPositive[int]()

	tests := []struct {
		name  string
		value int
		valid bool
	}{
		{"positive", 1, true},
		{"large positive", 1000, true},
		{"zero", 0, false},
		{"negative", -1, false},
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

func Test_IsPositive_Float(t *testing.T) {
	validator := IsPositive[float64]()

	tests := []struct {
		name  string
		value float64
		valid bool
	}{
		{"positive float", 0.1, true},
		{"large positive", 1000.5, true},
		{"zero", 0.0, false},
		{"negative", -1.5, false},
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

func Test_IsNegative_Int(t *testing.T) {
	validator := IsNegative[int]()

	tests := []struct {
		name  string
		value int
		valid bool
	}{
		{"negative", -1, true},
		{"large negative", -1000, true},
		{"zero", 0, false},
		{"positive", 1, false},
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

func Test_IsNegative_Float(t *testing.T) {
	validator := IsNegative[float64]()

	tests := []struct {
		name  string
		value float64
		valid bool
	}{
		{"negative float", -0.5, true},
		{"large negative", -100.5, true},
		{"zero", 0.0, false},
		{"positive", 1.5, false},
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

func Test_IsZero_Int(t *testing.T) {
	validator := IsZero[int]()

	tests := []struct {
		name  string
		value int
		valid bool
	}{
		{"zero", 0, true},
		{"positive", 1, false},
		{"negative", -1, false},
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

func Test_IsZero_Float(t *testing.T) {
	validator := IsZero[float64]()

	tests := []struct {
		name  string
		value float64
		valid bool
	}{
		{"zero", 0.0, true},
		{"small positive", 0.1, false},
		{"small negative", -0.1, false},
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

func Test_IsNonZero_Int(t *testing.T) {
	validator := IsNonZero[int]()

	tests := []struct {
		name  string
		value int
		valid bool
	}{
		{"positive", 1, true},
		{"negative", -1, true},
		{"zero", 0, false},
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

func Test_IsEven_Int(t *testing.T) {
	validator := IsEven[int]()

	tests := []struct {
		name  string
		value int
		valid bool
	}{
		{"even positive", 2, true},
		{"even negative", -4, true},
		{"zero", 0, true},
		{"odd positive", 3, false},
		{"odd negative", -5, false},
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

func Test_IsEven_Int64(t *testing.T) {
	validator := IsEven[int64]()

	tests := []struct {
		name  string
		value int64
		valid bool
	}{
		{"large even", 1000, true},
		{"large odd", 1001, false},
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

func Test_IsOdd_Int(t *testing.T) {
	validator := IsOdd[int]()

	tests := []struct {
		name  string
		value int
		valid bool
	}{
		{"odd positive", 3, true},
		{"odd negative", -5, true},
		{"even positive", 2, false},
		{"even negative", -4, false},
		{"zero", 0, false},
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

func Test_IsOdd_Int32(t *testing.T) {
	validator := IsOdd[int32]()

	tests := []struct {
		name  string
		value int32
		valid bool
	}{
		{"odd", 7, true},
		{"even", 8, false},
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

func Test_OneOf_Int(t *testing.T) {
	validator := NumberOneOf[int](10, 20, 30)

	tests := []struct {
		name  string
		value int
		valid bool
	}{
		{"allowed value 10", 10, true},
		{"allowed value 20", 20, true},
		{"allowed value 30", 30, true},
		{"not allowed value 15", 15, false},
		{"not allowed value 40", 40, false},
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

func Test_OneOf_Float(t *testing.T) {
	validator := NumberOneOf(1.5, 2.5, 3.5)

	tests := []struct {
		name  string
		value float64
		valid bool
	}{
		{"allowed value 1.5", 1.5, true},
		{"allowed value 2.5", 2.5, true},
		{"allowed value 3.5", 3.5, true},
		{"not allowed value 1.0", 1.0, false},
		{"not allowed value 4.0", 4.0, false},
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

func Test_MultipleValidatorsJoined(t *testing.T) {
	// Test that multiple errors are properly joined
	validator := Number("value",
		MinValue(10),
		MaxValue(20),
		IsEven[int](),
	)

	err := validator(5)
	if err == nil {
		t.Fatalf("expected error for value 5")
	}

	// Check that the error contains the validation name
	if !strings.Contains(err.Error(), "value") {
		t.Fatalf("expected error to contain field name, got: %v", err)
	}
}
