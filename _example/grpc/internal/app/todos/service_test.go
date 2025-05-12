package todos

import (
	"context"
	"github.com/0B1t322/zero-validation/grpc-example/pkg/api/todos"
	"github.com/0B1t322/zero-validation/validate"
	"testing"
)

func TestSearchTodosValidator(t *testing.T) {
	s := searchTodosValidator{}

	req := &todos.SearchTodosRequest{
		Filter: &todos.SearchTodosRequest_Filter{
			By: &todos.SearchTodosRequest_Filter_Name{},
		},
		Limit: 100,
	}

	rules := s.Rules()

	ctx := context.Background()

	err := validate.Struct(ctx, req, rules...)
	t.Log(err)
}
