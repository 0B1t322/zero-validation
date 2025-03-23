package main

import (
	"context"
	"fmt"
	"github.com/0B1t322/zero-validation/grpc-example/pkg/api/todos"
	"github.com/0B1t322/zero-validaton/rule"
	"github.com/0B1t322/zero-validaton/translation"
	"github.com/0B1t322/zero-validaton/validate"
)

func main() {
	ctx := context.Background()
	ctx = validate.ValidateContextToContext(
		ctx,
		validate.NewValidateContext(
			translation.GlobalRegistry(),
			translation.GlobalRegistry().DefaultLocale(),
			validate.WithFieldNameGetter(validate.FieldNameKey("proto")),
		),
	)
	err := validate.Struct(
		ctx,
		&todos.CreateSomeRequest{
			//BaseType: 1,
		},
		validate.Field(
			todos.ValidateCreateSomeRequest.BaseType,
			rule.Required[uint64](),
		),
	)
	fmt.Println(err)
}
