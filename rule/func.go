package rule

import validatecontext "github.com/0B1t322/zero-validaton/validate/context"

type FuncRule[T any] func(ctx validatecontext.Context, v T) error

func (f FuncRule[T]) Validate(ctx validatecontext.Context, v T) error {
	return f(ctx, v)
}
