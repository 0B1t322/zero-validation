package validatecontext

import (
	"context"
	fieldname "github.com/0B1t322/zero-validation/field/name"
	"github.com/0B1t322/zero-validation/translation"
)

type Context interface {
	GetRegistry() translation.Registry
	GetPreferredLocale() string
	FieldNameGetter() fieldname.Getter
	IsStopAfterFirstError() bool
}

var defaultFieldNameKey = fieldname.Key("defaultFieldNameKey")

type validateContext struct {
	registry translation.Registry

	preferredLocale string

	fieldNameGetter fieldname.Getter

	stopAfterFirstError bool
}

func (v validateContext) GetRegistry() translation.Registry {
	return v.registry
}

func (v validateContext) GetPreferredLocale() string {
	return v.preferredLocale
}

func (v validateContext) FieldNameGetter() fieldname.Getter {
	return v.fieldNameGetter
}

func (v validateContext) IsStopAfterFirstError() bool {
	return v.stopAfterFirstError
}

func newValidateContextFromContext(ctx context.Context) Context {
	vCtx, ok := FromContext(ctx)
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

	fieldName, isFind := fieldname.GetterFromContext(ctx)
	if !isFind {
		fieldName = defaultFieldNameKey
	}

	return validateContext{
		registry:            registry,
		preferredLocale:     locale,
		fieldNameGetter:     fieldName,
		stopAfterFirstError: true,
	}
}

// NewFromContext ...
func NewFromContext(ctx context.Context) Context {
	return newValidateContextFromContext(ctx)
}

// New ...
func New(
	registry translation.Registry,
	preferredLocale string,
	opts ...ContextOption,
) Context {
	o := newValidateContextOptions(opts...)
	return validateContext{
		registry:            registry,
		preferredLocale:     preferredLocale,
		fieldNameGetter:     o.fieldNameGetter,
		stopAfterFirstError: o.stopAfterFirstError,
	}
}

type ContextOption func(o *validateContextOptions)

type validateContextOptions struct {
	fieldNameGetter     fieldname.Getter
	stopAfterFirstError bool
}

func newValidateContextOptions(options ...ContextOption) *validateContextOptions {
	o := &validateContextOptions{
		fieldNameGetter:     defaultFieldNameKey,
		stopAfterFirstError: true,
	}

	for _, option := range options {
		option(o)
	}

	return o
}

func WithFieldNameGetter(f fieldname.Getter) ContextOption {
	return func(o *validateContextOptions) {
		o.fieldNameGetter = f
	}
}

func WithStopAfterFirstError() ContextOption {
	return func(o *validateContextOptions) {
		o.stopAfterFirstError = true
	}
}

func WithNotStopAfterFirstError() ContextOption {
	return func(o *validateContextOptions) {
		o.stopAfterFirstError = false
	}
}

type validateContextKey struct{}

// ToContext ...
func ToContext(ctx context.Context, vCtx Context) context.Context {
	return context.WithValue(ctx, validateContextKey{}, vCtx)
}

// FromContext ...
func FromContext(ctx context.Context) (Context, bool) {
	vCtx, ok := ctx.Value(validateContextKey{}).(Context)
	if !ok {
		return nil, false
	}

	return vCtx, true
}
