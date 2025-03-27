LOCAL_BIN:=$(CURDIR)/bin


GQLGEN_VERSON:=v0.17.68
GQLGEN_BIN:=$(LOCAL_BIN)/gqlgen
.PHONY: gqlgen-bin
gqlgen-bin: export GOBIN=$(LOCAL_BIN)
gqlgen-bin:
	go install github.com/99designs/gqlgen@$(GQLGEN_VERSON)

.PHONY: generate-graphql
generate-graphql: gqlgen-bin
	go generate ./internal/pkg/graphql/...