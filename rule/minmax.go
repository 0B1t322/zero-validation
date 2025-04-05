package rule

import (
	"github.com/0B1t322/zero-validation/errors"
	"github.com/0B1t322/zero-validation/internal/optional"
	validatecontext "github.com/0B1t322/zero-validation/validate/context"
)

var (
	// ErrValueShouldBeGTEMin ...
	ErrValueShouldBeGTEMin = errors.DefaultErrorObjectFactory().NewErrorObject("validation_value_gte_min")
	// ErrValueShouldBeLTEMax ...
	ErrValueShouldBeLTEMax = errors.DefaultErrorObjectFactory().NewErrorObject("validation_value_lte_max")
	// ErrValueBetweenRequired ...
	ErrValueBetweenRequired = errors.DefaultErrorObjectFactory().NewErrorObject("validation_value_between_required")
)

type minMaxOp uint8

const (
	gte minMaxOp = iota
	lte
	btw
)

type MinMaxRule[T Number] struct {
	op  minMaxOp
	min optional.Optional[T]
	max optional.Optional[T]
	err error
}

func (r *MinMaxRule[T]) Validate(_ validatecontext.Context, value T) error {
	return r.compare(value)
}

func (r *MinMaxRule[T]) compare(value T) error {
	switch r.op {
	case lte:
		maxValue := r.max.MustValue()
		if value > maxValue {
			return r.err
		}
	case gte:
		minValue := r.min.MustValue()
		if value < minValue {
			return r.err
		}
	case btw:
		minValue := r.min.MustValue()
		maxValue := r.max.MustValue()
		if !(value >= minValue && value <= maxValue) {
			return r.err
		}
	}
	return nil
}

type minMaxParams[T Number] struct {
	Min T
	Max T
}

func Min[T Number](min T) MinMaxRule[T] {
	return MinMaxRule[T]{
		op:  gte,
		min: optional.New(min),
		err: ErrValueShouldBeGTEMin.SetParams(
			minMaxParams[T]{
				Min: min,
			},
		),
	}
}

func Max[T Number](max T) MinMaxRule[T] {
	return MinMaxRule[T]{
		op:  lte,
		max: optional.New(max),
		err: ErrValueShouldBeLTEMax.SetParams(
			minMaxParams[T]{
				Max: max,
			},
		),
	}
}

func Between[T Number](min, max T) *MinMaxRule[T] {
	return &MinMaxRule[T]{
		op:  btw,
		min: optional.New(min),
		max: optional.New(max),
		err: ErrValueBetweenRequired.
			SetParams(
				minMaxParams[T]{
					Max: max,
					Min: min,
				},
			),
	}
}
