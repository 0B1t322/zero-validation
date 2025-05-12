package resolver

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

import (
	"github.com/0B1t322/zero-validation/gqlgen-example/internal/app/todos"
	"github.com/0B1t322/zero-validation/gqlgen-example/internal/pkg/graphql/api"
)

type Resolver struct {
	*todos.Implementation
}

func New(todosService *todos.Implementation) *Resolver {
	return &Resolver{
		Implementation: todosService,
	}
}

// Query returns api.QueryResolver implementation.
func (r *Resolver) Query() api.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
