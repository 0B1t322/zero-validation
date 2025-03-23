package todos

import (
	"context"
	"github.com/0B1t322/zero-validation/gqlgen-example/internal/pkg/graphql/api/model"
)

type TodoResolver struct {
}

func (t *TodoResolver) Tasks(ctx context.Context, obj *model.Todo) ([]*model.Task, error) {
	panic("implement me")
}
