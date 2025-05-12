package errorspresenter

import (
	"fmt"
	"github.com/0B1t322/zero-validation/errors"
	"strconv"
)

// PresentErrorsAsSimpleOneError find error in [errors.Errors], where error is
// not a [errors.Errors], [errors.ErrorSlice].
func PresentErrorsAsSimpleOneError(errs errors.Errors) error {
	if len(errs) == 0 {
		return nil
	}

	return presentErrorAsSimpleOneError("", errs)
}

func presentErrorAsSimpleOneError(field string, err error) error {
	switch e := err.(type) {
	case errors.Errors:
		for vfield, verr := range e {
			if isSliceIndex(vfield) {
				vfield = field
			}
			if mappedErr := presentErrorAsSimpleOneError(vfield, verr); mappedErr != nil {
				return mappedErr
			}
		}
		return nil
	case errors.ErrorSlice:
		if len(e) == 0 {
			return nil
		}

		return presentErrorAsSimpleOneError(field, e[0])
	default:
		return fmt.Errorf("%s: %w", field, e)
	}
}

func isSliceIndex(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}
