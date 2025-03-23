package validate

import (
	"context"
	"github.com/0B1t322/zero-validaton/translation"
)

type Context interface {
	GetRegistry() translation.Registry
	GetPreferredLocale() string
	FieldNameGetter() FieldNameGetter
}

var defaultFieldNameKey = FieldNameKey("defaultFieldNameKey")

type validateContext struct {
	registry translation.Registry

	preferredLocale string

	fieldNameGetter FieldNameGetter
}

func (v validateContext) GetRegistry() translation.Registry {
	return v.registry
}

func (v validateContext) GetPreferredLocale() string {
	return v.preferredLocale
}

func (v validateContext) FieldNameGetter() FieldNameGetter {
	return v.fieldNameGetter
}

func newValidateContextFromContext(ctx context.Context) Context {
	vCtx, ok := ValidateContextFromContext(ctx)
	if ok {
		return vCtx
	}

	registry, isFind := translation.RegistryFromContext(ctx)
	if !isFind {
		registry = translation.GlobalRegistry()
	}

	locale, isFind := translation.LocaleFromContext(ctx)
	if !isFind {
		locale = registry.DefaultLocale()
	}

	fieldName, isFind := FieldNameGetterFromContext(ctx)
	if !isFind {
		fieldName = defaultFieldNameKey
	}

	return validateContext{
		registry:        registry,
		preferredLocale: locale,
		fieldNameGetter: fieldName,
	}
}

// NewValidateContext ...
func NewValidateContext(
	registry translation.Registry,
	preferredLocale string,
	opts ...ContextOption,
) Context {
	o := newValidateContextOptions(opts...)
	return validateContext{
		registry:        registry,
		preferredLocale: preferredLocale,
		fieldNameGetter: o.fieldNameGetter,
	}
}

type ContextOption func(o *validateContextOptions)

type validateContextOptions struct {
	fieldNameGetter FieldNameGetter
}

func newValidateContextOptions(options ...ContextOption) *validateContextOptions {
	o := &validateContextOptions{
		fieldNameGetter: defaultFieldNameKey,
	}

	for _, option := range options {
		option(o)
	}

	return o
}

func WithFieldNameGetter(f FieldNameGetter) ContextOption {
	return func(o *validateContextOptions) {
		o.fieldNameGetter = f
	}
}

type validateContextKey struct{}

// ValidateContextToContext ...
func ValidateContextToContext(ctx context.Context, vCtx Context) context.Context {
	return context.WithValue(ctx, validateContextKey{}, vCtx)
}

// ValidateContextFromContext ...
func ValidateContextFromContext(ctx context.Context) (Context, bool) {
	vCtx, ok := ctx.Value(validateContextKey{}).(Context)
	if !ok {
		return nil, false
	}

	return vCtx, true
}
