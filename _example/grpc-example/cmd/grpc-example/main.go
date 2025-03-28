package main

import (
	"context"
	"encoding/json"
	"github.com/0B1t322/zero-validation/grpc-example/pkg/api/todos"
	_ "github.com/0B1t322/zero-validation/grpc-example/pkg/translationx"
	fieldname "github.com/0B1t322/zero-validaton/field/name"
	"github.com/0B1t322/zero-validaton/rule"
	"github.com/0B1t322/zero-validaton/translation"
	"github.com/0B1t322/zero-validaton/validate"
	validatecontext "github.com/0B1t322/zero-validaton/validate/context"
	"github.com/0B1t322/zero-validaton/validate/validators"
	"os"
)

type createSomeRequestValidator struct{}

func (createSomeRequestValidator) Name() string {
	return "createSomeRequestValidator"
}

func (createSomeRequestValidator) Rules() []validate.FieldRule[*todos.CreateSomeRequest] {
	return []validate.FieldRule[*todos.CreateSomeRequest]{
		validate.Field(
			todos.ValidateCreateSomeRequest.BaseType,
			rule.In[uint64](3, 4),
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
	}
}

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
				InnerMessage: &todos.CreateSomeRequest_InnerMessage{
					Some: "some",
				},
			},
		},
		validators.GetOrInitValidatorRules[createSomeRequestValidator]()...,
	)
	json.NewEncoder(os.Stdout).Encode(err)
}
