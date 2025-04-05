package validate

import (
	errors "github.com/0B1t322/zero-validation/errors"
	"github.com/0B1t322/zero-validation/field"
	validatecontext "github.com/0B1t322/zero-validation/validate/context"
)

type fielder[T any, V any] struct {
	valueExtractor field.StructField[T, V]
	key            string
}

func (f *fielder[T, V]) GetName() string {
	if f.key != "" {
		return f.valueExtractor.GetAdditionalName(f.key)
	}

	return f.valueExtractor.Name()
}

func (f *fielder[T, V]) SetNameKey(key string) {
	f.key = key
}

type fieldRule[T any, V any] struct {
	filder         *fielder[T, V]
	valueExtractor field.StructField[T, V]
	rules          []Rule[V]
}

type FieldRule[T any] interface {
	Validate(ctx validatecontext.Context, obj T) *errors.FieldError
}

func (f *fieldRule[T, V]) GetFieldName() string {
	return f.valueExtractor.Name()
}

func (f *fieldRule[T, V]) Validate(ctx validatecontext.Context, obj T) *errors.FieldError {
	v := f.valueExtractor.ExtractValue(obj)
	for _, rule := range f.rules {
		if err := rule.Validate(ctx, v); err != nil {
			err = TranslateError(ctx, err)
			return errors.NewFieldError(
				ctx.FieldNameGetter().GetFieldName(f.filder.valueExtractor),
				err,
			)
		}
	}
	return nil
}

func Field[T any, V any](structField field.StructField[T, V], rules ...Rule[V]) FieldRule[T] {
	return &fieldRule[T, V]{
		valueExtractor: structField,
		rules:          rules,
		filder: &fielder[T, V]{
			valueExtractor: structField,
		},
	}
}
