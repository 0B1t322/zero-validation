package validate

import (
	"context"
	"github.com/0B1t322/zero-validaton/translation"
)

type Context interface {
	context.Context
	GetRegistry() translation.Registry
	GetPreferredLocale() string
}

type validateContext struct {
	context.Context
	registry translation.Registry

	preferredLocale string
}

func (v *validateContext) GetRegistry() translation.Registry {
	return v.registry
}

func (v *validateContext) GetPreferredLocale() string {
	return v.preferredLocale
}

func newValidateContext(ctx context.Context) Context {
	registry, isFind := translation.RegistryFromContext(ctx)
	if !isFind {
		registry = translation.GlobalRegistry()
	}

	locale, isFind := translation.LocaleFromContext(ctx)
	if !isFind {
		locale = registry.DefaultLocale()
	}

	return &validateContext{
		Context:         ctx,
		registry:        registry,
		preferredLocale: locale,
	}
}
