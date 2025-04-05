package rule

import (
	"github.com/0B1t322/zero-validaton/errors"
	validatecontext "github.com/0B1t322/zero-validaton/validate/context"
)

var (
	ErrMinStringLenInvalid = errors.DefaultErrorObjectFactory().NewErrorObject("min_string_len_invalid")
)

type MinStringRunesCountRule struct {
	minLength int
	err       error
}

func MinStringRunesCount(count int) *MinStringRunesCountRule {
	return &MinStringRunesCountRule{
		minLength: count,
		err: ErrMinStringLenInvalid.SetParams(lenParams{
			Len: count,
		}),
	}
}

func (r *MinStringRunesCountRule) Validate(_ validatecontext.Context, v string) error {
	if len([]rune(v)) < r.minLength {
		return r.err
	}

	return nil
}
