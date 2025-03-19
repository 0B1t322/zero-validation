package validate

import (
	"context"
	"github.com/0B1t322/zero-validaton/translation"
)

type Context interface {
	GetRegistry() translation.Registry
	GetPreferredLocale() string
}

type validateContext struct {
	registry translation.Registry

	preferredLocale string
}

func (v validateContext) GetRegistry() translation.Registry {
	return v.registry
}

func (v validateContext) GetPreferredLocale() string {
	return v.preferredLocale
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

	return validateContext{
		registry:        registry,
		preferredLocale: locale,
	}
}

// NewValidateContext ...
func NewValidateContext(
	registry translation.Registry,
	preferredLocale string,
) Context {
	return validateContext{
		registry:        registry,
		preferredLocale: preferredLocale,
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
