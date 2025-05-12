package grpcmw

import (
	"context"
	fieldname "github.com/0B1t322/zero-validation/field/name"
	"github.com/0B1t322/zero-validation/translation"
	validatecontext "github.com/0B1t322/zero-validation/validate/context"
	"google.golang.org/grpc"
)

func WithValidateContext() grpc.UnaryServerInterceptor {
	vCtx := validatecontext.New(
		translation.GlobalRegistry(),
		"ru",
		validatecontext.WithFieldNameGetter(fieldname.NewGetterStrategy(fieldname.Proto, fieldname.JSON)),
	)
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		ctx = validatecontext.ToContext(ctx, vCtx)

		return handler(ctx, req)
	}
}
