package number

import (
	"errors"

	"github.com/mq-gabs/vld/config/utils"
	"github.com/mq-gabs/vld/config/validate"
)

type ConfigNumber[T typeNumber] struct {
	Max utils.V[T]
	Min utils.V[T]
	Gt  utils.V[T]
	Gte utils.V[T]
	Lt  utils.V[T]
	Lte utils.V[T]
	Eq  utils.V[T]
	Ne  utils.V[T]

	Pos  bool
	NPos bool
	Neg  bool
	NNeg bool

	Even bool
	Odd  bool

	Custom validate.Validate[T]
}

type numberValidator[T typeNumber] struct {
	validations []validate.Validate[T]
}

func (nv *numberValidator[T]) append(v validate.Validate[T]) {
	if v == nil {
		return
	}

	nv.validations = append(nv.validations, v)
}

func (nv *numberValidator[T]) Validate(v T) error {
	var err error
	for _, validate := range nv.validations {
		err = errors.Join(err, validate(v))
	}

	return err
}

func (cn ConfigNumber[T]) Build() validate.Validator[T] {
	nv := &numberValidator[T]{}

	if cn.Max.IsSet() {
		nv.append(buildMax(cn.Max.Get()))
	}

	if cn.Min.IsSet() {
		nv.append(buildMin(cn.Min.Get()))
	}

	if cn.Gt.IsSet() {
		nv.append(buildGt(cn.Gt.Get()))
	}

	if cn.Gte.IsSet() {
		nv.append(buildGte(cn.Gte.Get()))
	}

	if cn.Lt.IsSet() {
		nv.append(buildLt(cn.Lt.Get()))
	}

	if cn.Lte.IsSet() {
		nv.append(buildLte(cn.Lte.Get()))
	}

	if cn.Eq.IsSet() {
		nv.append(buildEq(cn.Eq.Get()))
	}

	if cn.Ne.IsSet() {
		nv.append(buildNe(cn.Ne.Get()))
	}

	if cn.Pos {
		nv.append(buildPos[T]())
	}

	if cn.NPos {
		nv.append(buildNPos[T]())
	}

	if cn.Neg {
		nv.append(buildNeg[T]())
	}

	if cn.NNeg {
		nv.append(buildNNeg[T]())
	}

	if cn.Even {
		nv.append(buildEven[T]())
	}

	if cn.Odd {
		nv.append(buildOdd[T]())
	}

	if cn.Custom != nil {
		nv.append(cn.Custom)
	}

	return nv
}

func buildMax[T typeNumber](v T) validate.Validate[T] {
	return func(t T) error {
		if t > v {
			return errors.New("greater than max value")
		}

		return nil
	}
}

func buildMin[T typeNumber](v T) validate.Validate[T] {
	return func(t T) error {
		if t < v {
			return errors.New("less than min value")
		}

		return nil
	}
}

func buildGt[T typeNumber](v T) validate.Validate[T] {
	return func(t T) error {
		if t <= v {
			return errors.New("must be greater than value")
		}

		return nil
	}
}

func buildGte[T typeNumber](v T) validate.Validate[T] {
	return func(t T) error {
		if t < v {
			return errors.New("must be greater than or equal to value")
		}

		return nil
	}
}

func buildLt[T typeNumber](v T) validate.Validate[T] {
	return func(t T) error {
		if t >= v {
			return errors.New("must be less than value")
		}

		return nil
	}
}

func buildLte[T typeNumber](v T) validate.Validate[T] {
	return func(t T) error {
		if t > v {
			return errors.New("must be less than or equal to value")
		}

		return nil
	}
}

func buildEq[T typeNumber](v T) validate.Validate[T] {
	return func(t T) error {
		if t != v {
			return errors.New("value does not match")
		}

		return nil
	}
}

func buildNe[T typeNumber](v T) validate.Validate[T] {
	return func(t T) error {
		if t == v {
			return errors.New("value must not match")
		}

		return nil
	}
}

func buildPos[T typeNumber]() validate.Validate[T] {
	return func(t T) error {
		if t <= 0 {
			return errors.New("must be positive")
		}

		return nil
	}
}

func buildNPos[T typeNumber]() validate.Validate[T] {
	return func(t T) error {
		if t > 0 {
			return errors.New("must be non-positive")
		}

		return nil
	}
}

func buildNeg[T typeNumber]() validate.Validate[T] {
	return func(t T) error {
		if t >= 0 {
			return errors.New("must be negative")
		}

		return nil
	}
}

func buildNNeg[T typeNumber]() validate.Validate[T] {
	return func(t T) error {
		if t < 0 {
			return errors.New("must be non-negative")
		}

		return nil
	}
}

func buildEven[T typeNumber]() validate.Validate[T] {
	return func(t T) error {
		if int64(t)%2 != 0 {
			return errors.New("must be even")
		}

		return nil
	}
}

func buildOdd[T typeNumber]() validate.Validate[T] {
	return func(t T) error {
		if int64(t)%2 == 0 {
			return errors.New("must be odd")
		}

		return nil
	}
}
