package main

import (
	"context"
	"encoding/json"
	"github.com/0B1t322/zero-validation/grpc-example/pkg/api/todos"
	fieldname "github.com/0B1t322/zero-validaton/field/name"
	"github.com/0B1t322/zero-validaton/rule"
	"github.com/0B1t322/zero-validaton/translation"
	"github.com/0B1t322/zero-validaton/validate"
	validatecontext "github.com/0B1t322/zero-validaton/validate/context"
	"os"
)

func main() {
	ctx := context.Background()
	ctx = validatecontext.ToContext(
		ctx,
		validatecontext.New(
			translation.GlobalRegistry(),
			translation.GlobalRegistry().DefaultLocale(),
			validatecontext.WithFieldNameGetter(fieldname.Proto),
		),
	)
	err := validate.Struct(
		ctx,
		&todos.CreateSomeRequest{
			BaseType: 1,
			OneofExample: &todos.CreateSomeRequest_InnerMessage_{
				InnerMessage: &todos.CreateSomeRequest_InnerMessage{},
			},
		},
		validate.Field(
			todos.ValidateCreateSomeRequest.BaseType,
			rule.Required[uint64](),
		),
		validate.IfFieldTypeOf[*todos.CreateSomeRequest_InnerMessage_](
			todos.ValidateCreateSomeRequest.OneofExample,
			validate.ObjectField(
				todos.ValidateCreateSomeRequest_InnerMessage_.InnerMessage,
				validate.Field(
					todos.ValidateCreateSomeRequest_InnerMessage.Some,
					rule.Required[string](),
				),
			),
		),
	)
	json.NewEncoder(os.Stdout).Encode(err)
}

func isOneOfExampleUint(req *todos.CreateSomeRequest) bool {
	_, ok := req.GetOneofExample().(*todos.CreateSomeRequest_Uint)
	return ok
}
