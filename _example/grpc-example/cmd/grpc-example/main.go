package main

import (
	"context"
	"fmt"
	"github.com/0B1t322/zero-validation/grpc-example/pkg/api/todos"
	"github.com/0B1t322/zero-validaton/rule"
	"github.com/0B1t322/zero-validaton/validate"
)

func main() {
	ctx := context.Background()
	err := validate.Struct(
		ctx,
		&todos.CreateSomeRequest{
			BaseType: 1,
		},
		validate.Field(
			todos.ValidateCreateSomeRequest.BaseType,
			rule.Required[uint64](),
		),
	)
	fmt.Println(err)
}
