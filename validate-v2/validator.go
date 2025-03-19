package validate

import (
	"context"
	"github.com/0B1t322/zero-validaton/errors-old"
	"github.com/0B1t322/zero-validaton/translation"
)

type Validator[T any] struct {
	registry translation.Registry
}

func (v Validator[T]) Struct(ctx context.Context, obj T, fieldRules ...FieldRule[T]) error {
	var errs errors_old.FieldErrors
	for _, fieldRule := range fieldRules {
		if err := fieldRule.Validate(ctx, obj); err != nil {
			if errs == nil {
				errs = errors_old.NewFieldErrors(errors_old.WithStartCap(len(fieldRules)))
			}
			errs = append(errs, err)
		}
	}

	if len(errs) == 0 {
		return nil
	}

	return errors_old.FieldErrorToErrors(errs)
}
