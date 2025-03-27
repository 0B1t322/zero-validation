package main

import (
	tasks2 "github.com/0B1t322/zero-validation/gqlgen-example/internal/app/graphql/tasks"
	todos2 "github.com/0B1t322/zero-validation/gqlgen-example/internal/app/graphql/todos"
	"github.com/0B1t322/zero-validation/gqlgen-example/internal/pkg/graphql/api"
	"github.com/0B1t322/zero-validation/gqlgen-example/internal/pkg/graphql/api/resolver"
	_ "github.com/0B1t322/zero-validation/gqlgen-example/internal/pkg/translationx"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

func main() {
	todos := &todos2.Todos{}
	tasks := &tasks2.Tasks{}

	executableSchema := api.NewExecutableSchema(api.Config{
		Resolvers:  resolver.NewResolver(todos, tasks),
		Directives: api.DirectiveRoot{},
		Complexity: api.ComplexityRoot{},
	})

	srv := handler.New(executableSchema)

	srv.AddTransport(transport.SSE{})
	srv.AddTransport(transport.POST{})

	srv.Use(extension.Introspection{})

	r := chi.NewRouter()

	r.Handle("/", playground.Handler("GraphQL playground", "/query"))
	r.Handle("/query", srv)

	log.Fatal(http.ListenAndServe(":8080", r))
}
