package errorspresenter

import (
	stgerror "errors"
	"github.com/0B1t322/zero-validation/errors"
	"strings"
)

// PresentErrorAsOnePathError present error as path of error in fields and error.
func PresentErrorAsOnePathError(errs errors.Errors) error {
	sb := strings.Builder{}
	presetErrorAsOnePathError(errs, &sb)

	return stgerror.New(sb.String())
}

func presetErrorAsOnePathError(err error, sb *strings.Builder) {
	switch e := err.(type) {
	case errors.Errors:
		for vfield, verr := range e {
			if isSliceIndex(vfield) {
				sb.WriteString("[")
				sb.WriteString(vfield)
				sb.WriteString("]")
			} else {
				sb.WriteString(vfield)
			}

			nextErrType := getErrorType(verr)
			if !nextErrType.isLast() && nextErrType != errorTypeErrorSlice {
				sb.WriteString(".")
			}

			presetErrorAsOnePathError(verr, sb)
			return
		}
	case errors.ErrorSlice:
		if len(e) == 0 {
			return
		}

		presetErrorAsOnePathError(e[0], sb)
	default:
		sb.WriteString(": ")
		sb.WriteString(e.Error())
	}
}

func getErrorType(err error) errorType {
	switch err.(type) {
	case errors.Errors:
		return errorTypeErrors
	case errors.ErrorSlice:
		return errorTypeErrorSlice
	default:
		return errorTypeError
	}
}

type errorType int

func (e errorType) isLast() bool {
	return e == errorTypeError
}

const (
	errorTypeErrors = iota + 1
	errorTypeErrorSlice
	errorTypeError
)
