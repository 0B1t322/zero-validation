package rule

import (
	"github.com/0B1t322/zero-validation/errors"
	validatecontext "github.com/0B1t322/zero-validation/validate/context"
)

var (
	ErrMaxSliceLenInvalid = errors.DefaultErrorObjectFactory().NewErrorObject("max_slice_len_invalid")
)

type lenParams struct {
	Len int
}

type MaxSliceLenRule[T any, P []T] struct {
	maxLen int
	err    error
}

func MaxSliceLen[T any, P []T](maxLen int) *MaxSliceLenRule[T, P] {
	return &MaxSliceLenRule[T, P]{
		maxLen: maxLen,
		err:    ErrMaxSliceLenInvalid.SetParams(lenParams{Len: maxLen}),
	}
}

func (m *MaxSliceLenRule[T, P]) Validate(_ validatecontext.Context, v P) error {
	if len(v) > m.maxLen {
		return m.err
	}

	return nil
}
