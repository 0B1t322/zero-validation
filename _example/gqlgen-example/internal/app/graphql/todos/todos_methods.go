package todos

import (
	"context"
	model1 "github.com/0B1t322/zero-validation/gqlgen-example/internal/model"
	"github.com/0B1t322/zero-validation/gqlgen-example/internal/pkg/graphql/api/model"
	"github.com/0B1t322/zero-validation/gqlgen-example/internal/pkg/graphql/validation_errors"
	"github.com/0B1t322/zero-validaton/rule"
	"github.com/0B1t322/zero-validaton/validate"
)

type Todos struct {
}

func (t *Todos) GetTodos(ctx context.Context, input model.GetTodosInput) ([]*model.Todo, error) {
	err := validate.Struct(
		ctx,
		&input,
		validate.Field(
			model.ValidateGetTodosInput.IDs,
			rule.RequiredSlice[string](),
		),
		validate.Field(
			model.ValidateGetTodosInput.TasksIDs,
			rule.RequiredSlice[string](),
		),
	)
	if err != nil {
		gqlError, _ := validation_errors.MapError(ctx, input, err)
		return nil, gqlError
	}

	return nil, nil
}

func (t *Todos) CreateTodo(ctx context.Context, input model.CreateTodoInput) ([]*model.Todo, error) {
	panic("implement me")
}

func (t *Todos) SuggestTodos(ctx context.Context, input model1.SuggestTodosInput) ([]*model.Todo, error) {
	panic("implement me")
}
