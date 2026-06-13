package slicev

import (
	"testing"

	"github.com/mq-gabs/vld/config/str"
	"github.com/mq-gabs/vld/config/utils"
)

/* =========================
   MIN / MAX LENGTH
========================= */

func TestSliceMinMaxLen(t *testing.T) {
	tests := []struct {
		name    string
		cfg     ConfigSlice[string]
		value   []string
		wantErr bool
	}{
		{
			name:    "min len pass",
			cfg:     ConfigSlice[string]{MinLen: utils.W(2)},
			value:   []string{"a", "b"},
			wantErr: false,
		},
		{
			name:    "min len fail",
			cfg:     ConfigSlice[string]{MinLen: utils.W(3)},
			value:   []string{"a"},
			wantErr: true,
		},
		{
			name:    "max len pass",
			cfg:     ConfigSlice[string]{MaxLen: utils.W(2)},
			value:   []string{"a"},
			wantErr: false,
		},
		{
			name:    "max len fail",
			cfg:     ConfigSlice[string]{MaxLen: utils.W(1)},
			value:   []string{"a", "b"},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.cfg.Build().Validate(tt.value)

			if tt.wantErr && err == nil {
				t.Fatalf("expected error but got nil")
			}
			if !tt.wantErr && err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
		})
	}
}

/* =========================
   NON EMPTY
========================= */

func TestSliceNonEmpty(t *testing.T) {
	tests := []struct {
		name    string
		cfg     ConfigSlice[string]
		value   []string
		wantErr bool
	}{
		{
			name:    "non empty pass",
			cfg:     ConfigSlice[string]{NonEmpty: true},
			value:   []string{"a"},
			wantErr: false,
		},
		{
			name:    "non empty fail",
			cfg:     ConfigSlice[string]{NonEmpty: true},
			value:   []string{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.cfg.Build().Validate(tt.value)

			if tt.wantErr && err == nil {
				t.Fatalf("expected error but got nil")
			}
			if !tt.wantErr && err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
		})
	}
}

/* =========================
   VALUE VALIDATOR (REAL STRING VALIDATOR)
========================= */

func TestSliceValueValidator(t *testing.T) {
	strCfg := str.ConfigString{
		Alpha: true,
	}
	cfg := ConfigSlice[string]{
		ValueValidator: strCfg.Build(),
	}

	tests := []struct {
		name    string
		value   []string
		wantErr bool
	}{
		{
			name:    "all valid",
			value:   []string{"abc", "xyz"},
			wantErr: false,
		},
		{
			name:    "one invalid",
			value:   []string{"abc", "x1y"},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := cfg.Build().Validate(tt.value)

			if tt.wantErr && err == nil {
				t.Fatalf("expected error but got nil")
			}
			if !tt.wantErr && err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
		})
	}
}

/* =========================
   COMBINED RULES
========================= */

func TestSliceCombinedRules(t *testing.T) {
	strCfg := str.ConfigString{
		Alpha:  true,
		MinLen: utils.W(1),
		MaxLen: utils.W(3),
	}
	cfg := ConfigSlice[string]{
		MinLen:   utils.W(1),
		MaxLen:   utils.W(3),
		NonEmpty: true,

		ValueValidator: strCfg.Build(),
	}

	tests := []struct {
		name    string
		value   []string
		wantErr bool
	}{
		{
			name:    "valid slice",
			value:   []string{"ok"},
			wantErr: false,
		},
		{
			name:    "invalid value",
			value:   []string{"1k"},
			wantErr: true,
		},
		{
			name:    "too large slice",
			value:   []string{"a", "b", "c", "d"},
			wantErr: true,
		},
		{
			name:    "empty slice fails non-empty",
			value:   []string{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := cfg.Build().Validate(tt.value)

			if tt.wantErr && err == nil {
				t.Fatalf("expected error but got nil")
			}
			if !tt.wantErr && err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
		})
	}
}

/* =========================
   EMPTY CONFIG
========================= */

func TestEmptySliceConfig(t *testing.T) {
	cfg := ConfigSlice[string]{}

	err := cfg.Build().Validate([]string{"a", "b", "c"})

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}
