package validate

import (
	"context"
	"github.com/0B1t322/zero-validaton/errors-old"
	"github.com/0B1t322/zero-validaton/field"
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
	GetFieldName() string
	Validate(ctx context.Context, obj T) *errors_old.FieldError
}

func (f *fieldRule[T, V]) Validate(ctx context.Context, obj T) *errors_old.FieldError {
	v := f.valueExtractor.ExtractValue(obj)
	for _, rule := range f.rules {
		if err := rule.Validate(v); err != nil {
			return errors_old.NewFieldError(f.filder, err)
		}
	}

	return nil
}

func (f *fieldRule[T, V]) GetFieldName() string {
	return f.valueExtractor.Name()
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
