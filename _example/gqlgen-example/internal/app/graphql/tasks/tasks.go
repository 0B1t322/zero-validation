package tasks

import (
	"context"
	model2 "github.com/0B1t322/zero-validation/gqlgen-example/internal/model"
	"github.com/0B1t322/zero-validation/gqlgen-example/internal/pkg/graphql/api/model"
	"github.com/0B1t322/zero-validation/gqlgen-example/internal/pkg/graphql/validation_errors"
	"github.com/0B1t322/zero-validaton/rule"
	"github.com/0B1t322/zero-validaton/validate"
)

type Tasks struct {
}

func (t *Tasks) CreateTasks(ctx context.Context, input model.CreateTasksInput) ([]*model.Task, error) {
	err := validate.Struct(
		ctx,
		&input,
		validate.Field(
			model.ValidateCreateTasksInput.TodoID,
			rule.Required[string](),
		),
		validate.Field(
			model.ValidateCreateTasksInput.Tasks,
			rule.RequiredSlice[*model2.CreateTasksInputItem](),
		),
		validate.ObjectSliceField(
			model.ValidateCreateTasksInput.Tasks,
			validate.Field(model2.ValidateCreateTasksInputItem.Name, rule.Required[string]()),
		),
	)
	if err != nil {
		gqlError, _ := validation_errors.MapError(ctx, input, err)
		return nil, gqlError
	}

	return nil, nil
}
