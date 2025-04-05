package validate

import validatecontext "github.com/0B1t322/zero-validation/validate/context"

type Rule[T any] interface {
	Validate(ctx validatecontext.Context, v T) error
}
