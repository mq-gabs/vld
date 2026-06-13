package str

import (
	"errors"
	"strings"

	"github.com/mq-gabs/vld/config/utils"
	"github.com/mq-gabs/vld/config/validate"
)

type ConfigString struct {
	MinLen utils.V[int]
	MaxLen utils.V[int]

	Eq utils.V[string]
	Ne utils.V[string]

	OneOf utils.V[[]string]

	Trim     bool
	Lower    bool
	Upper    bool
	Alpha    bool
	Numeric  bool
	Alphanum bool

	Email bool
	URL   bool
	UUID  bool
}

type stringValidator struct {
	validations []validate.Validate[string]
}

func (sv *stringValidator) append(v validate.Validate[string]) {
	if v == nil {
		return
	}
	sv.validations = append(sv.validations, v)
}

func (sv *stringValidator) Validate(v string) error {
	var err error
	for _, validation := range sv.validations {
		err = errors.Join(err, validation(v))
	}
	return err
}

func (c *ConfigString) Build() validate.Validator[string] {
	sv := &stringValidator{}

	if c.MinLen.IsSet() {
		sv.append(buildMinLen(c.MinLen.Get()))
	}

	if c.MaxLen.IsSet() {
		sv.append(buildMaxLen(c.MaxLen.Get()))
	}

	if c.Eq.IsSet() {
		sv.append(buildEq(c.Eq.Get()))
	}

	if c.Ne.IsSet() {
		sv.append(buildNe(c.Ne.Get()))
	}

	if c.OneOf.IsSet() {
		sv.append(buildOneOf(c.OneOf.Get()))
	}

	if c.Trim {
		sv.append(buildTrim())
	}

	if c.Lower {
		sv.append(buildLower())
	}

	if c.Upper {
		sv.append(buildUpper())
	}

	if c.Alpha {
		sv.append(buildAlpha())
	}

	if c.Numeric {
		sv.append(buildNumeric())
	}

	if c.Alphanum {
		sv.append(buildAlphanum())
	}

	if c.Email {
		sv.append(buildEmail())
	}

	if c.URL {
		sv.append(buildURL())
	}

	if c.UUID {
		sv.append(buildUUID())
	}

	return sv
}

/* =========================
   LENGTH RULES
========================= */

func buildMinLen(min int) validate.Validate[string] {
	return func(s string) error {
		if len(s) < min {
			return errors.New("string shorter than min length")
		}
		return nil
	}
}

func buildMaxLen(max int) validate.Validate[string] {
	return func(s string) error {
		if len(s) > max {
			return errors.New("string exceeds max length")
		}
		return nil
	}
}

/* =========================
   EQUALITY RULES
========================= */

func buildEq(v string) validate.Validate[string] {
	return func(s string) error {
		if s != v {
			return errors.New("string does not match")
		}
		return nil
	}
}

func buildNe(v string) validate.Validate[string] {
	return func(s string) error {
		if s == v {
			return errors.New("string must not match")
		}
		return nil
	}
}

/* =========================
   ONE OF
========================= */

func buildOneOf(values []string) validate.Validate[string] {
	set := make(map[string]struct{}, len(values))
	for _, v := range values {
		set[v] = struct{}{}
	}

	return func(s string) error {
		if _, ok := set[s]; !ok {
			return errors.New("string not in allowed list")
		}
		return nil
	}
}

/* =========================
   TRANSFORMS / CHECKS
========================= */

func buildTrim() validate.Validate[string] {
	return func(s string) error {
		if strings.TrimSpace(s) != s {
			return errors.New("string must be trimmed")
		}
		return nil
	}
}

func buildLower() validate.Validate[string] {
	return func(s string) error {
		if strings.ToLower(s) != s {
			return errors.New("string must be lowercase")
		}
		return nil
	}
}

func buildUpper() validate.Validate[string] {
	return func(s string) error {
		if strings.ToUpper(s) != s {
			return errors.New("string must be uppercase")
		}
		return nil
	}
}

/* =========================
   CHARACTER RULES
========================= */

func buildAlpha() validate.Validate[string] {
	return func(s string) error {
		for _, r := range s {
			if !(r >= 'a' && r <= 'z' || r >= 'A' && r <= 'Z') {
				return errors.New("string must contain only letters")
			}
		}
		return nil
	}
}

func buildNumeric() validate.Validate[string] {
	return func(s string) error {
		for _, r := range s {
			if r < '0' || r > '9' {
				return errors.New("string must contain only numbers")
			}
		}
		return nil
	}
}

func buildAlphanum() validate.Validate[string] {
	return func(s string) error {
		for _, r := range s {
			if !((r >= 'a' && r <= 'z') ||
				(r >= 'A' && r <= 'Z') ||
				(r >= '0' && r <= '9')) {
				return errors.New("string must be alphanumeric")
			}
		}
		return nil
	}
}

/* =========================
   FORMAT RULES (simple stubs)
========================= */

func buildEmail() validate.Validate[string] {
	return func(s string) error {
		if !strings.Contains(s, "@") {
			return errors.New("invalid email")
		}
		return nil
	}
}

func buildURL() validate.Validate[string] {
	return func(s string) error {
		if !(strings.HasPrefix(s, "http://") || strings.HasPrefix(s, "https://")) {
			return errors.New("invalid URL")
		}
		return nil
	}
}

func buildUUID() validate.Validate[string] {
	return func(s string) error {
		if len(s) != 36 {
			return errors.New("invalid UUID")
		}
		return nil
	}
}
