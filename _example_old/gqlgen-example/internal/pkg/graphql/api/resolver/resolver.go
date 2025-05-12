package resolver

// THIS CODE WILL BE UPDATED WITH SCHEMA CHANGES. PREVIOUS IMPLEMENTATION FOR SCHEMA CHANGES WILL BE KEPT IN THE COMMENT SECTION. IMPLEMENTATION FOR UNCHANGED SCHEMA WILL BE KEPT.

import (
	"github.com/0B1t322/zero-validation/gqlgen-example/internal/app/graphql/tasks"
	todos2 "github.com/0B1t322/zero-validation/gqlgen-example/internal/app/graphql/todos"
	"github.com/0B1t322/zero-validation/gqlgen-example/internal/pkg/graphql/api"
)

type Resolver struct {
	*todos2.Todos
	*tasks.Tasks
}

func NewResolver(todos *todos2.Todos, tasks *tasks.Tasks) *Resolver {
	return &Resolver{
		Todos: todos,
		Tasks: tasks,
	}
}

// Mutation returns api.MutationResolver implementation.
func (r *Resolver) Mutation() api.MutationResolver { return r }

// Query returns api.QueryResolver implementation.
func (r *Resolver) Query() api.QueryResolver { return r }

// Todo returns api.TodoResolver implementation.
func (r *Resolver) Todo() api.TodoResolver { return &todos2.TodoResolver{} }
