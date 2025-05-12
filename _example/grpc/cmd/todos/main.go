package main

import (
	"fmt"
	errorspresenter "github.com/0B1t322/zero-validation/errors/presenter"
	fieldname "github.com/0B1t322/zero-validation/field/name"
	"github.com/0B1t322/zero-validation/grpc-example/internal/app/todos"
	_ "github.com/0B1t322/zero-validation/grpc-example/internal/pkg/translationx"
	todos2 "github.com/0B1t322/zero-validation/grpc-example/pkg/api/todos"
	zerovalidationgrpcmw "github.com/0B1t322/zero-validation/mw/grpc"
	"github.com/0B1t322/zero-validation/translation"
	validatecontext "github.com/0B1t322/zero-validation/validate/context"
	"github.com/0B1t322/zero-validation/validate/validators"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func main() {
	store := validators.NewDefaultMapStore()
	impl := todos.New(store)

	grpcSrv := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			zerovalidationgrpcmw.WithOnceValidateContext(func() validatecontext.Context {
				vCtx := validatecontext.New(
					translation.GlobalRegistry(),
					"ru",
					validatecontext.WithFieldNameGetter(fieldname.NewGetterStrategy("ru", fieldname.Proto, fieldname.JSON)),
				)

				return vCtx
			}),
			zerovalidationgrpcmw.WithErrorPresenter(
				errorspresenter.PresenterFunc(errorspresenter.PresentErrorsAsSimpleOneError),
			),
		),
	)

	todos2.RegisterTodoServiceServer(grpcSrv, impl)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 8080))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	reflection.Register(grpcSrv)

	if err := grpcSrv.Serve(lis); err != nil {
		panic(err)
	}
}
