package tasks

import (
	"context"
	"github.com/0B1t322/zero-validation/gqlgen-example/internal/pkg/graphql/api/model"
)

type Tasks struct {
}

func (t *Tasks) CreateTasks(ctx context.Context, input model.CreateTasksInput) ([]*model.Task, error) {
	panic("implement me")
}
