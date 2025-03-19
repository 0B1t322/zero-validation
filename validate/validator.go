package validate

import (
	"context"

	errors "github.com/0B1t322/zero-validaton/errors"
)

func Struct[T any](ctx context.Context, obj T, fieldRules ...FieldRule[T]) error {
	vCtx := newValidateContextFromContext(ctx)

	var errs errors.FieldErrors
	for _, fieldRule := range fieldRules {
		if err := fieldRule.Validate(vCtx, obj); err != nil {
			if errs == nil {
				errs = errors.NewFieldErrors(errors.WithStartCap(len(fieldRules)))
			}
			errs = append(errs, err)
		}
	}

	if len(errs) == 0 {
		return nil
	}

	return errors.FieldErrorToErrors(errs)
}
