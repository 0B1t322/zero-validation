package todos

import (
	"context"
	"encoding/json"
	"github.com/0B1t322/zero-validation/ogen/pkg/api/todos"
	"github.com/0B1t322/zero-validation/validate"
	"github.com/0B1t322/zero-validation/validate/validators"
	"net/http"
)

type Implementation struct {
	store validators.ValidatorStore
}

func (i Implementation) TodosV1SearchGet(ctx context.Context, params api.TodosV1SearchGetParams) (api.TodosV1SearchGetRes, error) {
	err := validate.Struct(ctx, params)
	if err != nil {
		validationErrorsBytes, _ := json.Marshal(err)
		raw := api.TodosV1SearchGetBadRequestApplicationJSON(validationErrorsBytes)
		return &raw, nil
	}

	return &api.TodosV1SearchGetOK{}, nil
}

func (i Implementation) NewError(ctx context.Context, err error) *api.ErrRespStatusCode {
	return &api.ErrRespStatusCode{
		StatusCode: http.StatusInternalServerError,
		Response:   api.ErrResp{},
	}
}

func NewImplementation() *Implementation {
	return &Implementation{}
}
