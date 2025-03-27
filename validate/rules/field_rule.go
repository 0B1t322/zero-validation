package rules

import (
	"github.com/0B1t322/zero-validaton/errors"
	validatecontext "github.com/0B1t322/zero-validaton/validate/context"
)

type FieldRule[T any] interface {
	Validate(ctx validatecontext.Context, obj T) *errors.FieldError
}

type Rule[T any] interface {
	Validate(ctx validatecontext.Context, obj T) error
}
