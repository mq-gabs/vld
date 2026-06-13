package number

import (
	"testing"

	"github.com/mq-gabs/vld/config/utils"
)

func TestNumberValidations(t *testing.T) {
	tests := []struct {
		name    string
		cfg     ConfigNumber[int]
		value   int
		wantErr bool
	}{
		// Max
		{
			name:    "max pass",
			cfg:     ConfigNumber[int]{Max: utils.W(10)},
			value:   10,
			wantErr: false,
		},
		{
			name:    "max fail",
			cfg:     ConfigNumber[int]{Max: utils.W(10)},
			value:   11,
			wantErr: true,
		},

		// Min
		{
			name:    "min pass",
			cfg:     ConfigNumber[int]{Min: utils.W(10)},
			value:   10,
			wantErr: false,
		},
		{
			name:    "min fail",
			cfg:     ConfigNumber[int]{Min: utils.W(10)},
			value:   9,
			wantErr: true,
		},

		// Gt
		{
			name:    "gt pass",
			cfg:     ConfigNumber[int]{Gt: utils.W(10)},
			value:   11,
			wantErr: false,
		},
		{
			name:    "gt fail equal",
			cfg:     ConfigNumber[int]{Gt: utils.W(10)},
			value:   10,
			wantErr: true,
		},
		{
			name:    "gt fail less",
			cfg:     ConfigNumber[int]{Gt: utils.W(10)},
			value:   9,
			wantErr: true,
		},

		// Gte
		{
			name:    "gte pass equal",
			cfg:     ConfigNumber[int]{Gte: utils.W(10)},
			value:   10,
			wantErr: false,
		},
		{
			name:    "gte pass greater",
			cfg:     ConfigNumber[int]{Gte: utils.W(10)},
			value:   11,
			wantErr: false,
		},
		{
			name:    "gte fail",
			cfg:     ConfigNumber[int]{Gte: utils.W(10)},
			value:   9,
			wantErr: true,
		},

		// Lt
		{
			name:    "lt pass",
			cfg:     ConfigNumber[int]{Lt: utils.W(10)},
			value:   9,
			wantErr: false,
		},
		{
			name:    "lt fail equal",
			cfg:     ConfigNumber[int]{Lt: utils.W(10)},
			value:   10,
			wantErr: true,
		},
		{
			name:    "lt fail greater",
			cfg:     ConfigNumber[int]{Lt: utils.W(10)},
			value:   11,
			wantErr: true,
		},

		// Lte
		{
			name:    "lte pass equal",
			cfg:     ConfigNumber[int]{Lte: utils.W(10)},
			value:   10,
			wantErr: false,
		},
		{
			name:    "lte pass less",
			cfg:     ConfigNumber[int]{Lte: utils.W(10)},
			value:   9,
			wantErr: false,
		},
		{
			name:    "lte fail",
			cfg:     ConfigNumber[int]{Lte: utils.W(10)},
			value:   11,
			wantErr: true,
		},

		// Eq
		{
			name:    "eq pass",
			cfg:     ConfigNumber[int]{Eq: utils.W(10)},
			value:   10,
			wantErr: false,
		},
		{
			name:    "eq fail",
			cfg:     ConfigNumber[int]{Eq: utils.W(10)},
			value:   11,
			wantErr: true,
		},

		// Ne
		{
			name:    "ne pass",
			cfg:     ConfigNumber[int]{Ne: utils.W(10)},
			value:   11,
			wantErr: false,
		},
		{
			name:    "ne fail",
			cfg:     ConfigNumber[int]{Ne: utils.W(10)},
			value:   10,
			wantErr: true,
		},

		// Positive
		{
			name:    "positive pass",
			cfg:     ConfigNumber[int]{Pos: true},
			value:   1,
			wantErr: false,
		},
		{
			name:    "positive fail zero",
			cfg:     ConfigNumber[int]{Pos: true},
			value:   0,
			wantErr: true,
		},
		{
			name:    "positive fail negative",
			cfg:     ConfigNumber[int]{Pos: true},
			value:   -1,
			wantErr: true,
		},

		// Non Positive
		{
			name:    "non positive pass zero",
			cfg:     ConfigNumber[int]{NPos: true},
			value:   0,
			wantErr: false,
		},
		{
			name:    "non positive pass negative",
			cfg:     ConfigNumber[int]{NPos: true},
			value:   -1,
			wantErr: false,
		},
		{
			name:    "non positive fail",
			cfg:     ConfigNumber[int]{NPos: true},
			value:   1,
			wantErr: true,
		},

		// Negative
		{
			name:    "negative pass",
			cfg:     ConfigNumber[int]{Neg: true},
			value:   -1,
			wantErr: false,
		},
		{
			name:    "negative fail zero",
			cfg:     ConfigNumber[int]{Neg: true},
			value:   0,
			wantErr: true,
		},
		{
			name:    "negative fail positive",
			cfg:     ConfigNumber[int]{Neg: true},
			value:   1,
			wantErr: true,
		},

		// Non Negative
		{
			name:    "non negative pass positive",
			cfg:     ConfigNumber[int]{NNeg: true},
			value:   1,
			wantErr: false,
		},
		{
			name:    "non negative pass zero",
			cfg:     ConfigNumber[int]{NNeg: true},
			value:   0,
			wantErr: false,
		},
		{
			name:    "non negative fail",
			cfg:     ConfigNumber[int]{NNeg: true},
			value:   -1,
			wantErr: true,
		},

		// Even
		{
			name:    "even pass",
			cfg:     ConfigNumber[int]{Even: true},
			value:   2,
			wantErr: false,
		},
		{
			name:    "even fail",
			cfg:     ConfigNumber[int]{Even: true},
			value:   3,
			wantErr: true,
		},

		// Odd
		{
			name:    "odd pass",
			cfg:     ConfigNumber[int]{Odd: true},
			value:   3,
			wantErr: false,
		},
		{
			name:    "odd fail",
			cfg:     ConfigNumber[int]{Odd: true},
			value:   2,
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

func TestMultipleValidations(t *testing.T) {
	cfg := ConfigNumber[int]{
		Min:  utils.W(1),
		Max:  utils.W(10),
		Even: true,
	}

	validator := cfg.Build()

	tests := []struct {
		name    string
		value   int
		wantErr bool
	}{
		{
			name:    "valid minimum",
			value:   2,
			wantErr: false,
		},
		{
			name:    "valid maximum",
			value:   10,
			wantErr: false,
		},
		{
			name:    "fails max",
			value:   12,
			wantErr: true,
		},
		{
			name:    "fails min",
			value:   0,
			wantErr: true,
		},
		{
			name:    "fails even",
			value:   3,
			wantErr: true,
		},
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

func TestEmptyConfig(t *testing.T) {
	cfg := ConfigNumber[int]{}

	err := cfg.Build().Validate(123)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}
