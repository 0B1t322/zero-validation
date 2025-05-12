package main

import (
	"fmt"
	"github.com/0B1t322/zero-validation/grpc-example/internal/app/todos"
	grpcmw "github.com/0B1t322/zero-validation/grpc-example/internal/pkg/mw/grpc"
	_ "github.com/0B1t322/zero-validation/grpc-example/internal/pkg/translationx"
	todos2 "github.com/0B1t322/zero-validation/grpc-example/pkg/api/todos"
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
			grpcmw.WithValidateContext(),
			grpcmw.ZeroValidationErrorHandler(
				grpcmw.NewZeroValidationErrorHandlerWithOneError(),
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
