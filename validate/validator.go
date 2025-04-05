package validate

import (
	"context"
	validatecontext "github.com/0B1t322/zero-validation/validate/context"

	errors "github.com/0B1t322/zero-validation/errors"
)

func Struct[T any](ctx context.Context, obj T, fieldRules ...FieldRule[T]) error {
	vCtx := validatecontext.NewFromContext(ctx)

	var errs errors.FieldErrors
	for _, fieldRule := range fieldRules {
		if err := fieldRule.Validate(vCtx, obj); err != nil {
			if errs == nil {
				errs = errors.NewFieldErrors(
					errors.WithStartCap(fieldErrorsStartCap(vCtx.IsStopAfterFirstError(), len(fieldRules))),
				)
			}
			errs = append(errs, err)

			if vCtx.IsStopAfterFirstError() {
				break
			}
		}
	}

	if len(errs) == 0 {
		return nil
	}

	return errors.FieldErrorToErrors(errs)
}

func fieldErrorsStartCap(isStopAfterFirstErrors bool, fieldErrorsCount int) int {
	if isStopAfterFirstErrors {
		return 1
	}

	return fieldErrorsCount
}
