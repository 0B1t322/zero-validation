package validate

import (
	"github.com/0B1t322/zero-validation/errors"
	"github.com/0B1t322/zero-validation/field"
	validatecontext "github.com/0B1t322/zero-validation/validate/context"
)

type objectField[T any, V any] struct {
	extractor  field.StructField[T, V]
	fielder    *fielder[T, V]
	fieldRules []FieldRule[V]
}

func (c *objectField[T, V]) GetFieldName() string {
	return c.extractor.Name()
}

func (c *objectField[T, V]) Validate(ctx validatecontext.Context, obj T) *errors.FieldError {
	value := c.extractor.ExtractValue(obj)

	return c.validate(ctx, value)
}

func (c *objectField[T, V]) validateRules(ctx validatecontext.Context, obj V) errors.FieldErrors {
	var fieldErrors errors.FieldErrors
	for _, rule := range c.fieldRules {
		if err := rule.Validate(ctx, obj); err != nil {
			if fieldErrors == nil {
				fieldErrors = errors.NewFieldErrors()
			}
			fieldErrors = append(fieldErrors, err)
			if ctx.IsStopAfterFirstError() {
				break
			}
		}
	}

	return fieldErrors
}

func (c *objectField[T, V]) validate(ctx validatecontext.Context, obj V) *errors.FieldError {
	errs := c.validateRules(ctx, obj)

	if len(errs) == 0 {
		return nil
	}

	return errors.NewFieldError(
		ctx.FieldNameGetter().GetFieldName(c.fielder.valueExtractor),
		errs,
	)
}

func ObjectField[T any, V any](
	structField field.StructField[T, V],
	fieldRules ...FieldRule[V],
) FieldRule[T] {
	return &objectField[T, V]{
		extractor:  structField,
		fieldRules: fieldRules,
		fielder: &fielder[T, V]{
			valueExtractor: structField,
		},
	}
}
