package mapv

import (
	"errors"
	"testing"

	"github.com/mq-gabs/vld/config/str"
	"github.com/mq-gabs/vld/config/utils"
)

/* =========================
   MIN / MAX LENGTH
========================= */

func TestMapMinMaxLen(t *testing.T) {
	tests := []struct {
		name    string
		cfg     ConfigMap[string, string]
		value   map[string]string
		wantErr bool
	}{
		{
			name:    "min len pass",
			cfg:     ConfigMap[string, string]{MinLen: utils.W(2)},
			value:   map[string]string{"a": "x", "b": "y"},
			wantErr: false,
		},
		{
			name:    "min len fail",
			cfg:     ConfigMap[string, string]{MinLen: utils.W(3)},
			value:   map[string]string{"a": "x"},
			wantErr: true,
		},
		{
			name:    "max len pass",
			cfg:     ConfigMap[string, string]{MaxLen: utils.W(2)},
			value:   map[string]string{"a": "x"},
			wantErr: false,
		},
		{
			name:    "max len fail",
			cfg:     ConfigMap[string, string]{MaxLen: utils.W(1)},
			value:   map[string]string{"a": "x", "b": "y"},
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
   KEY RULES
========================= */

func TestMapKeyRules(t *testing.T) {
	tests := []struct {
		name    string
		cfg     ConfigMap[string, string]
		value   map[string]string
		wantErr bool
	}{
		{
			name: "required keys pass",
			cfg: ConfigMap[string, string]{
				RequiredKeys: utils.W([]string{"a", "b"}),
			},
			value:   map[string]string{"a": "x", "b": "y"},
			wantErr: false,
		},
		{
			name: "required keys fail",
			cfg: ConfigMap[string, string]{
				RequiredKeys: utils.W([]string{"a", "b"}),
			},
			value:   map[string]string{"a": "x"},
			wantErr: true,
		},
		{
			name: "allowed keys pass",
			cfg: ConfigMap[string, string]{
				AllowedKeys: utils.W([]string{"a", "b"}),
			},
			value:   map[string]string{"a": "x"},
			wantErr: false,
		},
		{
			name: "allowed keys fail",
			cfg: ConfigMap[string, string]{
				AllowedKeys: utils.W([]string{"a"}),
			},
			value:   map[string]string{"b": "x"},
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

func TestMapNonEmpty(t *testing.T) {
	tests := []struct {
		name    string
		cfg     ConfigMap[string, string]
		value   map[string]string
		wantErr bool
	}{
		{
			name:    "non empty pass",
			cfg:     ConfigMap[string, string]{NonEmpty: true},
			value:   map[string]string{"a": "x"},
			wantErr: false,
		},
		{
			name:    "non empty fail",
			cfg:     ConfigMap[string, string]{NonEmpty: true},
			value:   map[string]string{},
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

func TestMapValueValidator(t *testing.T) {
	strCfg := &str.ConfigString{
		Alpha: true,
	}
	cfg := ConfigMap[string, string]{
		ValueValidator: strCfg.Build(),
	}

	tests := []struct {
		name    string
		value   map[string]string
		wantErr bool
	}{
		{
			name:    "valid values",
			value:   map[string]string{"a": "abc", "b": "xyz"},
			wantErr: false,
		},
		{
			name:    "invalid value contains number",
			value:   map[string]string{"a": "abc", "b": "x1y"},
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

func TestMapCombinedRules(t *testing.T) {
	strCfg := str.ConfigString{
		Alpha:  true,
		MinLen: utils.W(1),
		MaxLen: utils.W(3),
	}
	cfg := ConfigMap[string, string]{
		MinLen:       utils.W(1),
		MaxLen:       utils.W(3),
		RequiredKeys: utils.W([]string{"a"}),
		AllowedKeys:  utils.W([]string{"a", "b"}),

		ValueValidator: strCfg.Build(),

		NonEmpty: true,
	}

	tests := []struct {
		name    string
		value   map[string]string
		wantErr bool
	}{
		{
			name:    "valid map",
			value:   map[string]string{"a": "ok"},
			wantErr: false,
		},
		{
			name:    "invalid value",
			value:   map[string]string{"a": "1k"},
			wantErr: true,
		},
		{
			name:    "missing required key",
			value:   map[string]string{"b": "ok"},
			wantErr: true,
		},
		{
			name:    "too many entries",
			value:   map[string]string{"a": "ok", "b": "ok", "c": "ok", "d": "ok"},
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

func TestEmptyConfigMap(t *testing.T) {
	cfg := ConfigMap[string, string]{}

	err := cfg.Build().Validate(map[string]string{"a": "x"})

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestCustomValidation(t *testing.T) {
	cfg := ConfigMap[string, int]{
		Custom: func(m map[string]int) error {
			if len(m) == 2 {
				return errors.New("length cannot be 2")
			}

			if _, ok := m["invalid"]; ok {
				return errors.New("invalid key cannot exists")
			}

			return nil
		},
	}.Build()

	err := cfg.Validate(map[string]int{
		"age": 18,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	err = cfg.Validate(map[string]int{
		"day":   18,
		"month": 1,
		"year":  2000,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	err = cfg.Validate(map[string]int{
		"day":   18,
		"month": 1,
	})
	if err == nil {
		t.Fatal("expected error")
	}

	err = cfg.Validate(map[string]int{
		"invalid": 18,
		"month":   1,
	})
	if err == nil {
		t.Fatal("expected error")
	}
}
