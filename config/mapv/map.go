package mapv

import (
	"errors"

	"github.com/mq-gabs/vld/config/utils"
	"github.com/mq-gabs/vld/config/validate"
)

type ConfigMap[K comparable, V any] struct {
	MinLen utils.V[int]
	MaxLen utils.V[int]

	RequiredKeys utils.V[[]K]
	AllowedKeys  utils.V[[]K]

	NonEmpty bool

	// NEW: external value validator (composed from string/number/etc)
	ValueValidator validate.Validator[V]
}

type mapValidator[K comparable, V any] struct {
	valueValidator validate.Validator[V]
	validations    []validate.Validate[map[K]V]
}

func (mv *mapValidator[K, V]) append(v validate.Validate[map[K]V]) {
	if v == nil {
		return
	}
	mv.validations = append(mv.validations, v)
}

func (mv *mapValidator[K, V]) Validate(m map[K]V) error {
	var err error

	// 1. structure validations
	for _, validation := range mv.validations {
		err = errors.Join(err, validation(m))
	}

	// 2. value validations (delegated)
	if mv.valueValidator != nil {
		for _, v := range m {
			if e := mv.valueValidator.Validate(v); e != nil {
				err = errors.Join(err, e)
			}
		}
	}

	return err
}

func (c *ConfigMap[K, V]) Build() validate.Validator[map[K]V] {
	mv := &mapValidator[K, V]{
		valueValidator: c.ValueValidator,
	}

	if c.MinLen.IsSet() {
		mv.append(buildMinLenMap[K, V](c.MinLen.Get()))
	}

	if c.MaxLen.IsSet() {
		mv.append(buildMaxLenMap[K, V](c.MaxLen.Get()))
	}

	if c.RequiredKeys.IsSet() {
		mv.append(buildRequiredKeys[K, V](c.RequiredKeys.Get()))
	}

	if c.AllowedKeys.IsSet() {
		mv.append(buildAllowedKeys[K, V](c.AllowedKeys.Get()))
	}

	if c.NonEmpty {
		mv.append(buildNonEmptyMap[K, V]())
	}

	return mv
}

/* =========================
   SIZE RULES
========================= */

func buildMinLenMap[K comparable, V any](min int) validate.Validate[map[K]V] {
	return func(m map[K]V) error {
		if len(m) < min {
			return errors.New("map has fewer items than min length")
		}
		return nil
	}
}

func buildMaxLenMap[K comparable, V any](max int) validate.Validate[map[K]V] {
	return func(m map[K]V) error {
		if len(m) > max {
			return errors.New("map exceeds max length")
		}
		return nil
	}
}

/* =========================
   KEY RULES
========================= */

func buildRequiredKeys[K comparable, V any](keys []K) validate.Validate[map[K]V] {
	return func(m map[K]V) error {
		for _, k := range keys {
			if _, ok := m[k]; !ok {
				return errors.New("missing required key")
			}
		}
		return nil
	}
}

func buildAllowedKeys[K comparable, V any](keys []K) validate.Validate[map[K]V] {
	allowed := make(map[K]struct{}, len(keys))
	for _, k := range keys {
		allowed[k] = struct{}{}
	}

	return func(m map[K]V) error {
		for k := range m {
			if _, ok := allowed[k]; !ok {
				return errors.New("map contains disallowed key")
			}
		}
		return nil
	}
}

/* =========================
   STRUCTURE RULES
========================= */

func buildNonEmptyMap[K comparable, V any]() validate.Validate[map[K]V] {
	return func(m map[K]V) error {
		if len(m) == 0 {
			return errors.New("map must not be empty")
		}
		return nil
	}
}
