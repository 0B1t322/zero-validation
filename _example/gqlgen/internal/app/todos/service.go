package todos

import (
	"context"
	gqlgenerrors "github.com/0B1t322/zero-validation/errors/map/gqlgen"
	"github.com/0B1t322/zero-validation/gqlgen-example/internal/pkg/graphql/api/model"
	"github.com/0B1t322/zero-validation/rule"
	"github.com/0B1t322/zero-validation/validate"
	"github.com/0B1t322/zero-validation/validate/validators"
)

// Implementation ...
type Implementation struct {
	store validators.ValidatorStore
}

// NewImplementation ...
func NewImplementation(store validators.ValidatorStore) *Implementation {
	validators.InitValidatorInStore(store, searchTodosInputValidator{})
	return &Implementation{
		store: store,
	}
}

func (i *Implementation) SearchTodos(ctx context.Context, input model.SearchTodosInput) (*model.SearchTodosResponse, error) {
	if err := validate.Struct(ctx, &input, validators.GetValidatorRulesFromStore[searchTodosInputValidator](i.store)...); err != nil {
		mappedError, isMapped := gqlgenerrors.MapError(ctx, input, err)
		if isMapped {
			return nil, mappedError
		}
		return nil, err
	}

	return nil, nil
}

type searchTodosInputValidator struct{}

func (s searchTodosInputValidator) Name() string {
	return "searchTodosInputValidator"
}

func (s searchTodosInputValidator) Rules() []validate.FieldRule[*model.SearchTodosInput] {
	filterRules := []validate.FieldRule[*model.SearchTodosInputFilter]{
		validate.If(s.isByIDs,
			validate.ObjectField(model.ValidateSearchTodosInputFilter.ByIDs,
				validate.Field(model.ValidateSearchTodosInputFilterByIDs.IDs, rule.RequiredSlice[string](), rule.MaxSliceLen[string](2)),
			),
		),
		validate.If(s.isByName,
			validate.ObjectField(model.ValidateSearchTodosInputFilter.ByName,
				validate.Field(model.ValidateSearchTodosInputFilterByName.Name, rule.MinStringRunesCount(3)),
			)),
	}
	return []validate.FieldRule[*model.SearchTodosInput]{
		validate.Field(model.ValidateSearchTodosInput.Limit, rule.Between[int64](1, 100)),
		validate.If(s.isFilterNotNil,
			validate.ObjectField(model.ValidateSearchTodosInput.Filter,
				filterRules...,
			),
		),
	}
}
func (s searchTodosInputValidator) isFilterNotNil(inp *model.SearchTodosInput) bool {
	return inp.Filter != nil
}
func (s searchTodosInputValidator) isByIDs(filter *model.SearchTodosInputFilter) bool {
	return filter.ByIDs != nil
}
func (s searchTodosInputValidator) isByName(filter *model.SearchTodosInputFilter) bool {
	return filter.ByName != nil
}
