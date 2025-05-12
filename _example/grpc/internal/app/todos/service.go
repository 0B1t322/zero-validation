package todos

import (
	"context"
	"github.com/0B1t322/zero-validation/grpc-example/pkg/api/todos"
	"github.com/0B1t322/zero-validation/rule"
	"github.com/0B1t322/zero-validation/validate"
	"github.com/0B1t322/zero-validation/validate/validators"
)

type Implementation struct {
	todos.UnimplementedTodoServiceServer

	store validators.ValidatorStore
}

func New(store validators.ValidatorStore) *Implementation {
	return &Implementation{
		store: store,
	}
}

func (i *Implementation) SearchTodos(ctx context.Context, req *todos.SearchTodosRequest) (*todos.SearchTodosResponse, error) {
	err := validate.Struct(ctx, req, validators.GetOrInitValidatorRulesFromStore[searchTodosValidator](i.store)...)
	if err != nil {
		return nil, err
	}

	return &todos.SearchTodosResponse{}, nil
}

type searchTodosValidator struct{}

func (s searchTodosValidator) Name() string {
	return "searchTodosValidator"
}
func (s searchTodosValidator) Rules() []validate.FieldRule[*todos.SearchTodosRequest] {
	searchByIDsRules := []validate.FieldRule[*todos.SearchTodosRequest_Filter_ByIDs]{
		validate.Field(todos.ValidateSearchTodosRequest_Filter_ByIDs.Ids,
			rule.RequiredSlice[uint64](), rule.MaxSliceLen[uint64](2)),
	}

	searchByNameRules := []validate.FieldRule[*todos.SearchTodosRequest_Filter_ByName]{
		validate.Field(todos.ValidateSearchTodosRequest_Filter_ByName.Name,
			rule.MinStringRunesCount(3)),
	}

	filterRules := []validate.FieldRule[*todos.SearchTodosRequest_Filter]{
		validate.IfFieldTypeOf[*todos.SearchTodosRequest_Filter_ByName_](todos.ValidateSearchTodosRequest_Filter.By,
			validate.Field(todos.ValidateSearchTodosRequest_Filter_ByName_.ByName, rule.NotNil[todos.SearchTodosRequest_Filter_ByName]()),
			validate.ObjectField(todos.ValidateSearchTodosRequest_Filter_ByName_.ByName, searchByNameRules...)),
		validate.IfFieldTypeOf[*todos.SearchTodosRequest_Filter_Ids](todos.ValidateSearchTodosRequest_Filter.By,
			validate.Field(todos.ValidateSearchTodosRequest_Filter_Ids.Ids, rule.NotNil[todos.SearchTodosRequest_Filter_ByIDs]()),
			validate.ObjectField(todos.ValidateSearchTodosRequest_Filter_Ids.Ids, searchByIDsRules...)),
		validate.Field(todos.ValidateSearchTodosRequest_Filter.By,
			rule.Required[todos.IsSearchTodosRequest_Filter_By]()),
	}

	return []validate.FieldRule[*todos.SearchTodosRequest]{
		validate.Field(todos.ValidateSearchTodosRequest.Limit, rule.Between[uint64](1, 100)),
		validate.Field(todos.ValidateSearchTodosRequest.Filter, rule.NotNil[todos.SearchTodosRequest_Filter]()),
		validate.If(s.isSearchTodosRequestFilterNotNil,
			validate.ObjectField(todos.ValidateSearchTodosRequest.Filter, filterRules...)),
	}
}
func (s searchTodosValidator) isSearchTodosRequestFilterNotNil(req *todos.SearchTodosRequest) bool {
	return req.GetFilter() != nil
}
