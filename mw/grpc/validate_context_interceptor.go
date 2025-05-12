package zerovalidationgrpcmw

import (
	"context"
	fieldname "github.com/0B1t322/zero-validation/field/name"
	"github.com/0B1t322/zero-validation/translation"
	validatecontext "github.com/0B1t322/zero-validation/validate/context"
	"google.golang.org/grpc"
)

// OnceValidateContextFactory ...
type OnceValidateContextFactory func() validatecontext.Context

// DefaultOnceValidateContextFactory ...
func DefaultOnceValidateContextFactory() OnceValidateContextFactory {
	return func() validatecontext.Context {
		vCtx := validatecontext.New(
			translation.GlobalRegistry(),
			"ru",
			validatecontext.WithFieldNameGetter(fieldname.NewGetterStrategy(fieldname.Proto, fieldname.JSON)),
		)

		return vCtx
	}
}

// WithOnceValidateContext ...
func WithOnceValidateContext(factory OnceValidateContextFactory) grpc.UnaryServerInterceptor {
	vCtx := factory()
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		ctx = validatecontext.ToContext(ctx, vCtx)

		return handler(ctx, req)
	}
}
