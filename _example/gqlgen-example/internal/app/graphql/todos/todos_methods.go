package todos

import (
	"context"
	model1 "github.com/0B1t322/zero-validation/gqlgen-example/internal/model"
	"github.com/0B1t322/zero-validation/gqlgen-example/internal/pkg/graphql/api/model"
)

type Todos struct {
}

func (t *Todos) GetTodos(ctx context.Context, input model.GetTodosInput) ([]*model.Todo, error) {
	panic("implement me")
}

func (t *Todos) CreateTodo(ctx context.Context, input model.CreateTodoInput) ([]*model.Todo, error) {
	panic("implement me")
}

func (t *Todos) SuggestTodos(ctx context.Context, input model1.SuggestTodosInput) ([]*model.Todo, error) {
	panic("implement me")
}
