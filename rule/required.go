package rule

import (
	errors "github.com/0B1t322/zero-validaton/errors"
	validatecontext "github.com/0B1t322/zero-validaton/validate/context"
)

var (
	ErrRequired = errors.DefaultErrorObjectFactory().NewErrorObject("required")
)

// Required check that value is not default
func Required[T comparable]() *RequiredRule[T] {
	return &RequiredRule[T]{
		err: ErrRequired,
	}
}

type RequiredRule[T comparable] struct {
	err error
}

func (r *RequiredRule[T]) Validate(_ validatecontext.Context, value T) error {
	var def T
	if def == value {
		return r.err
	}

	return nil
}

func (r *RequiredRule[T]) SetError(err error) *RequiredRule[T] {
	r.err = err
	return r
}

type RequiredPtrRule[T comparable, P *T] struct {
	err error
}

func RequiredPtr[T comparable, P *T]() *RequiredPtrRule[T, P] {
	return &RequiredPtrRule[T, P]{
		err: ErrRequired,
	}
}

func (r *RequiredPtrRule[T, P]) Validate(_ validatecontext.Context, value P) error {
	if value == nil {
		return r.err
	}

	var def T
	if *value == def {
		return r.err
	}

	return nil
}

func (r *RequiredPtrRule[T, P]) SetError(err error) *RequiredPtrRule[T, P] {
	r.err = err
	return r
}

type RequiredSliceRule[T any, P []T] struct {
	err error
}

// RequiredSlice ...
func RequiredSlice[T any, P []T]() *RequiredSliceRule[T, P] {
	return &RequiredSliceRule[T, P]{
		err: ErrRequired,
	}
}

func (r *RequiredSliceRule[T, P]) SetError(err error) *RequiredSliceRule[T, P] {
	r.err = err
	return r
}

func (r *RequiredSliceRule[T, P]) Validate(_ validatecontext.Context, value P) error {
	if len(value) == 0 {
		return r.err
	}

	return nil
}

type NotNilRule[T any, P *T] struct {
	err error
}

func NotNil[T any, P *T]() *NotNilRule[T, P] {
	return &NotNilRule[T, P]{
		err: ErrRequired,
	}
}

func (r *NotNilRule[T, P]) Validate(_ validatecontext.Context, value P) error {
	if value == nil {
		return r.err
	}

	return nil
}
