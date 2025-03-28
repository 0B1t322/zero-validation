package rule

import (
	"github.com/0B1t322/zero-validaton/errors"
	validatecontext "github.com/0B1t322/zero-validaton/validate/context"
)

var ErrValueNotIn = errors.DefaultErrorObjectFactory().NewErrorObject("validation_in_invalid")
var ErrValueIn = errors.DefaultErrorObjectFactory().NewErrorObject("validation_not_in_invalid")

type inComparableParams[T comparable] struct {
	In []T
}

type inOp uint8

const (
	inOpIn inOp = iota
	inOpNot
)

type InComparableRule[T comparable] struct {
	In  []T
	err error

	op inOp
}

func In[T comparable](values ...T) *InComparableRule[T] {
	return &InComparableRule[T]{
		In:  values,
		err: ErrValueNotIn.SetParams(inComparableParams[T]{In: values}),
		op:  inOpIn,
	}
}

func NotIn[T comparable](values ...T) *InComparableRule[T] {
	return &InComparableRule[T]{
		In:  values,
		err: ErrValueIn.SetParams(inComparableParams[T]{In: values}),
		op:  inOpNot,
	}
}

func (r *InComparableRule[T]) Validate(_ validatecontext.Context, v T) error {
	return r.check(v)
}

func (r *InComparableRule[T]) check(v T) error {
	switch r.op {
	case inOpIn:
		for _, in := range r.In {
			if v == in {
				return nil
			}
		}
		return r.err
	case inOpNot:
		for _, in := range r.In {
			if v == in {
				return r.err
			}
			return nil
		}
	}

	return nil
}
