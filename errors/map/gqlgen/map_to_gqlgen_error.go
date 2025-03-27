package gqlgen

import (
	"context"
	"errors"
	verrors "github.com/0B1t322/zero-validaton/errors"
	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func MapError(ctx context.Context, arg any, err error) (*gqlerror.Error, bool) {
	if !graphql.HasOperationContext(ctx) {
		return nil, false
	}

	validationErrors := new(verrors.Errors)
	if !errors.As(err, &validationErrors) {
		return nil, false
	}

	fieldCtx := graphql.GetRootFieldContext(ctx)
	// TODO: translate validation error to error with path
}

func buildPathContext()
