package validate

import (
	"github.com/0B1t322/zero-validation/errors"
	"github.com/0B1t322/zero-validation/field"
	validatecontext "github.com/0B1t322/zero-validation/validate/context"
)

type fieldRuleIf[T any] struct {
	ifFunc func(T) bool

	fieldRules []FieldRule[T]
}

func (f *fieldRuleIf[T]) Validate(ctx validatecontext.Context, obj T) *errors.FieldError {
	if !f.ifFunc(obj) {
		return nil
	}

	for _, rule := range f.fieldRules {
		if err := rule.Validate(ctx, obj); err != nil {
			return err
		}
	}

	return nil
}

func If[T any](ifFunc func(v T) bool, fieldRules ...FieldRule[T]) FieldRule[T] {
	return &fieldRuleIf[T]{
		ifFunc:     ifFunc,
		fieldRules: fieldRules,
	}
}

type ifFieldTypeOfRule[R, T, V any] struct {
	structField field.StructField[T, V]
	fieldRules  []FieldRule[R]
}

func (r *ifFieldTypeOfRule[R, T, V]) Validate(ctx validatecontext.Context, obj T) *errors.FieldError {
	v := r.structField.ExtractValue(obj)
	casted, isImplement := tryCast[V, R](v)
	if !isImplement {
		return nil
	}

	for _, rule := range r.fieldRules {
		if err := rule.Validate(ctx, casted); err != nil {
			return err
		}
	}

	return nil
}

func IfFieldTypeOf[R any, T any, V any](structField field.StructField[T, V], fieldRules ...FieldRule[R]) FieldRule[T] {
	return &ifFieldTypeOfRule[R, T, V]{
		structField: structField,
		fieldRules:  fieldRules,
	}
}

func tryCast[T any, V any](t T) (V, bool) {
	casted, ok := any(t).(V)
	return casted, ok
}
