package rule

import (
	"github.com/0B1t322/zero-validaton/errors"
	validatecontext "github.com/0B1t322/zero-validaton/validate/context"
)

type fieldRule[T any] interface {
	Validate(ctx validatecontext.Context, obj T) *errors.FieldError
}

type IfTypeOfRule[T any, V any] struct {
	ifTypeOfRules []fieldRule[V]
}

func IfTypeOf[T any, V any](ifTypeOfRules ...fieldRule[V]) *IfTypeOfRule[T, V] {
	return &IfTypeOfRule[T, V]{
		ifTypeOfRules: ifTypeOfRules,
	}
}

func (r *IfTypeOfRule[T, V]) Validate(ctx validatecontext.Context, value T) error {
	casted, isImplement := tryCast[T, V](value)
	if !isImplement {
		return nil
	}

	var errs errors.FieldErrors
	for _, rule := range r.ifTypeOfRules {
		if err := rule.Validate(ctx, casted); err != nil {
			if errs == nil {
				errs = errors.NewFieldErrors(errors.WithStartCap(len(r.ifTypeOfRules)))
			}
			errs = append(errs, err)
		}
	}

	if len(errs) == 0 {
		return nil
	}

	return errors.FieldErrorToErrors(errs)
}

func tryCast[T any, V any](t T) (V, bool) {
	casted, ok := any(t).(V)
	return casted, ok
}
