package validation_errors

import (
	"context"
	"errors"
	verrors "github.com/0B1t322/zero-validaton/errors"
	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/ast"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"reflect"
)

func MapError(ctx context.Context, arg any, err error) (gqlerror.List, bool) {
	if !graphql.HasOperationContext(ctx) {
		return nil, false
	}

	validationErrors := verrors.Errors{}
	if !errors.As(err, &validationErrors) {
		return nil, false
	}

	fieldCtx := graphql.GetRootFieldContext(ctx)
	gqlError := translateValidationErrorsToGqlError(ctx, fieldCtx, arg, validationErrors)
	if gqlError == nil {
		return nil, false
	}
	return gqlError, true
}

func translateValidationErrorsToGqlError(ctx context.Context, rootFieldContext *graphql.RootFieldContext, arg any, validationErrors verrors.Errors) gqlerror.List {
	// надо в лист добавлять все ошибки с разным path
	typeName := reflect.TypeOf(arg).Name()

	argument, isFind := findArgumentOfType(rootFieldContext.Field.Arguments, typeName)
	if !isFind {
		return nil
	}

	patchCtx := graphql.NewPathWithField(argument.Name)
	patchCtx.Parent = graphql.GetPathContext(ctx)

	return validationErrorsToGqlListErrorsByChildrens(patchCtx, argument.Value.Children, validationErrors)
}

func validationErrorsToGqlErrors(pathCtx *graphql.PathContext, argument *ast.Argument, validationErrors verrors.Errors) gqlerror.List {
	list := make(gqlerror.List, 0, len(validationErrors))
	for _, childValue := range argument.Value.Children {
		fieldError, isFind := validationErrors[childValue.Name]
		if !isFind {
			continue
		}

		nextPathCtx := graphql.NewPathWithField(childValue.Name)
		nextPathCtx.Parent = pathCtx

		list = append(list, &gqlerror.Error{
			Path:    nextPathCtx.Path(),
			Err:     fieldError,
			Message: fieldError.Error(),
		})
	}

	return list
}

func validationErrorsToGqlListErrorsByChildrens(pathCtx *graphql.PathContext, childValues ast.ChildValueList, validationErrors verrors.Errors) gqlerror.List {
	list := make(gqlerror.List, 0, len(validationErrors))
	for _, childValue := range childValues {
		fieldError, isFind := validationErrors[childValue.Name]
		if !isFind {
			continue
		}

		nextPathCtx := graphql.NewPathWithField(childValue.Name)
		nextPathCtx.Parent = pathCtx

		if innerFieldErrors, ok := fieldError.(verrors.Errors); ok {
			list = append(list,
				validationErrorsToGqlListErrorsByChildrens(pathCtx, childValue.Value.Children, innerFieldErrors)...)
			continue
		}

		list = append(list, &gqlerror.Error{
			Path:    nextPathCtx.Path(),
			Err:     fieldError,
			Message: fieldError.Error(),
		})
	}

	return list
}

func findArgumentOfType(arguments ast.ArgumentList, typeName string) (*ast.Argument, bool) {
	var findedArgument *ast.Argument
	for _, argument := range arguments {
		if argument.Value.Definition.Name == typeName {
			findedArgument = argument
			break
		}
	}

	if findedArgument == nil {
		return nil, false
	}

	return findedArgument, true
}
