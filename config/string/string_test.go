package str

import (
	"testing"

	"github.com/mq-gabs/vld/config/utils"
)

func TestStringValidations(t *testing.T) {
	tests := []struct {
		name    string
		cfg     ConfigString
		value   string
		wantErr bool
	}{
		// MinLen
		{
			name:    "min len pass",
			cfg:     ConfigString{MinLen: utils.W(3)},
			value:   "abc",
			wantErr: false,
		},
		{
			name:    "min len fail",
			cfg:     ConfigString{MinLen: utils.W(3)},
			value:   "ab",
			wantErr: true,
		},

		// MaxLen
		{
			name:    "max len pass",
			cfg:     ConfigString{MaxLen: utils.W(3)},
			value:   "abc",
			wantErr: false,
		},
		{
			name:    "max len fail",
			cfg:     ConfigString{MaxLen: utils.W(3)},
			value:   "abcd",
			wantErr: true,
		},

		// Eq
		{
			name:    "eq pass",
			cfg:     ConfigString{Eq: utils.W("test")},
			value:   "test",
			wantErr: false,
		},
		{
			name:    "eq fail",
			cfg:     ConfigString{Eq: utils.W("test")},
			value:   "wrong",
			wantErr: true,
		},

		// Ne
		{
			name:    "ne pass",
			cfg:     ConfigString{Ne: utils.W("test")},
			value:   "wrong",
			wantErr: false,
		},
		{
			name:    "ne fail",
			cfg:     ConfigString{Ne: utils.W("test")},
			value:   "test",
			wantErr: true,
		},

		// OneOf
		{
			name:    "oneof pass",
			cfg:     ConfigString{OneOf: utils.W([]string{"a", "b", "c"})},
			value:   "b",
			wantErr: false,
		},
		{
			name:    "oneof fail",
			cfg:     ConfigString{OneOf: utils.W([]string{"a", "b", "c"})},
			value:   "x",
			wantErr: true,
		},

		// Trim
		{
			name:    "trim pass",
			cfg:     ConfigString{Trim: true},
			value:   "abc",
			wantErr: false,
		},
		{
			name:    "trim fail",
			cfg:     ConfigString{Trim: true},
			value:   " abc ",
			wantErr: true,
		},

		// Lower
		{
			name:    "lower pass",
			cfg:     ConfigString{Lower: true},
			value:   "abc",
			wantErr: false,
		},
		{
			name:    "lower fail",
			cfg:     ConfigString{Lower: true},
			value:   "Abc",
			wantErr: true,
		},

		// Upper
		{
			name:    "upper pass",
			cfg:     ConfigString{Upper: true},
			value:   "ABC",
			wantErr: false,
		},
		{
			name:    "upper fail",
			cfg:     ConfigString{Upper: true},
			value:   "AbC",
			wantErr: true,
		},

		// Alpha
		{
			name:    "alpha pass",
			cfg:     ConfigString{Alpha: true},
			value:   "abcXYZ",
			wantErr: false,
		},
		{
			name:    "alpha fail",
			cfg:     ConfigString{Alpha: true},
			value:   "abc123",
			wantErr: true,
		},

		// Numeric
		{
			name:    "numeric pass",
			cfg:     ConfigString{Numeric: true},
			value:   "12345",
			wantErr: false,
		},
		{
			name:    "numeric fail",
			cfg:     ConfigString{Numeric: true},
			value:   "12a45",
			wantErr: true,
		},

		// Alphanum
		{
			name:    "alphanum pass",
			cfg:     ConfigString{Alphanum: true},
			value:   "abc123",
			wantErr: false,
		},
		{
			name:    "alphanum fail",
			cfg:     ConfigString{Alphanum: true},
			value:   "abc-123",
			wantErr: true,
		},

		// Email
		{
			name:    "email pass",
			cfg:     ConfigString{Email: true},
			value:   "test@mail.com",
			wantErr: false,
		},
		{
			name:    "email fail",
			cfg:     ConfigString{Email: true},
			value:   "testmail.com",
			wantErr: true,
		},

		// URL
		{
			name:    "url pass http",
			cfg:     ConfigString{URL: true},
			value:   "http://example.com",
			wantErr: false,
		},
		{
			name:    "url pass https",
			cfg:     ConfigString{URL: true},
			value:   "https://example.com",
			wantErr: false,
		},
		{
			name:    "url fail",
			cfg:     ConfigString{URL: true},
			value:   "ftp://example.com",
			wantErr: true,
		},

		// UUID
		{
			name:    "uuid pass",
			cfg:     ConfigString{UUID: true},
			value:   "123e4567-e89b-12d3-a456-426614174000",
			wantErr: false,
		},
		{
			name:    "uuid fail",
			cfg:     ConfigString{UUID: true},
			value:   "invalid-uuid",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			validator := tt.cfg.Build()

			err := validator.Validate(tt.value)

			if tt.wantErr && err == nil {
				t.Fatalf("expected error but got nil")
			}

			if !tt.wantErr && err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
		})
	}
}

func TestStringMultipleRules(t *testing.T) {
	cfg := ConfigString{
		MinLen: utils.W(3),
		MaxLen: utils.W(6),
		Alpha:  true,
	}

	validator := cfg.Build()

	tests := []struct {
		name    string
		value   string
		wantErr bool
	}{
		{"valid", "abc", false},
		{"too short", "ab", true},
		{"too long", "abcdefg", true},
		{"invalid chars", "ab1", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validator.Validate(tt.value)

			if tt.wantErr && err == nil {
				t.Fatalf("expected error but got nil")
			}

			if !tt.wantErr && err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
		})
	}
}

func TestEmptyStringConfig(t *testing.T) {
	cfg := ConfigString{}

	err := cfg.Build().Validate("anything")

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}
