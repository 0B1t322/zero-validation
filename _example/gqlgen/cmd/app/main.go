package main

import (
	todos2 "github.com/0B1t322/zero-validation/gqlgen-example/internal/app/todos"
	"github.com/0B1t322/zero-validation/gqlgen-example/internal/pkg/graphql/api"
	"github.com/0B1t322/zero-validation/gqlgen-example/internal/pkg/graphql/api/resolver"
	_ "github.com/0B1t322/zero-validation/gqlgen-example/internal/pkg/translationx"
	"github.com/0B1t322/zero-validation/validate/validators"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

func main() {
	store := validators.NewDefaultMapStore()
	todos := todos2.NewImplementation(store)

	executableSchema := api.NewExecutableSchema(api.Config{
		Resolvers: resolver.New(todos),
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
