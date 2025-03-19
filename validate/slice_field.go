package validate

import (
	"fmt"
	"github.com/0B1t322/zero-validaton/errors"
	"github.com/0B1t322/zero-validaton/field"
)

type sliceFieldRule[T any, V any] struct {
	fielder        *fielder[T, []V]
	valueExtractor field.StructField[T, []V]
	rules          []Rule[V]
}

func (s *sliceFieldRule[T, V]) GetFieldName() string {
	return s.valueExtractor.Name()
}

func (s *sliceFieldRule[T, V]) Validate(ctx Context, obj T) *errors.FieldError {
	slice := s.valueExtractor.ExtractValue(obj)

	var errs errors.FieldErrors
	for _, rule := range s.rules {
		for i, value := range slice {
			var err error
			err = rule.Validate(value)
			if err != nil {
				if errs == nil {
					errs = errors.NewFieldErrors()
				}

				err = TranslateError(ctx, err)
				errs = append(
					errs,
					errors.NewFieldError(
						errors.NewStringField(fmt.Sprint(i)),
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
		s.fielder,
		errors.FieldErrorToErrors(errs),
	)
}

func SliceField[T any, V any](
	field field.StructField[T, []V],
	rules ...Rule[V],
) FieldRule[T] {
	return &sliceFieldRule[T, V]{
		valueExtractor: field,
		fielder: &fielder[T, []V]{
			valueExtractor: field,
		},
		rules: rules,
	}
}
