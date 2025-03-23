package validate

import (
	"fmt"
	"github.com/0B1t322/zero-validaton/errors"
	"github.com/0B1t322/zero-validaton/field"
)

type objectSliceFieldRule[T any, V any] struct {
	extractor  field.StructField[T, []V]
	fielder    *fielder[T, []V]
	fieldRules []FieldRule[V]
}

func (s *objectSliceFieldRule[T, V]) GetFieldName() string {
	return s.extractor.Name()
}

func (s *objectSliceFieldRule[T, V]) Validate(ctx Context, obj T) *errors.FieldError {
	slice := s.extractor.ExtractValue(obj)

	var errs errors.FieldErrors
	for _, fieldRule := range s.fieldRules {
		for i, value := range slice {
			var err *errors.FieldError
			err = fieldRule.Validate(ctx, value)

			if err != nil {
				if errs == nil {
					errs = errors.NewFieldErrors()
				}
				errs = append(
					errs,
					errors.NewFieldError(
						fmt.Sprint(i),
						err,
					),
				)
				break
			}
		}
	}

	if len(errs) == 0 {
		return nil
	}

	return errors.NewFieldError(
		ctx.FieldNameGetter().GetFieldName(s.fielder.valueExtractor),
		errors.FieldErrorToErrors(errs),
	)
}

func ObjectSliceField[T any, V any](
	field field.StructField[T, []V],
	fieldRules ...FieldRule[V],
) FieldRule[T] {
	return &objectSliceFieldRule[T, V]{
		extractor: field,
		fielder: &fielder[T, []V]{
			valueExtractor: field,
		},
		fieldRules: fieldRules,
	}
}
