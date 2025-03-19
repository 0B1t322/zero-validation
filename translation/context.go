package translation

import "context"

type registryCtxKey struct{}

// RegistryToContext ...
func RegistryToContext(ctx context.Context, registry Registry) context.Context {
	return context.WithValue(ctx, registryCtxKey{}, registry)
}

// RegistryFromContext ...
func RegistryFromContext(ctx context.Context) (Registry, bool) {
	registry, ok := ctx.Value(registryCtxKey{}).(Registry)
	return registry, ok
}

type localeCtxKey struct{}

func LocaleToContext(ctx context.Context, locale string) context.Context {
	return context.WithValue(ctx, localeCtxKey{}, locale)
}

func LocaleFromContext(ctx context.Context) (string, bool) {
	locale, ok := ctx.Value(localeCtxKey{}).(string)
	return locale, ok
}
