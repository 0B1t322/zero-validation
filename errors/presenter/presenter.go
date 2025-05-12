package errorspresenter

import "github.com/0B1t322/zero-validation/errors"

// Presenter ...
type Presenter interface {
	Present(errors errors.Errors) error
}

// PresenterFunc ...
type PresenterFunc func(errors errors.Errors) error

// Present ...
func (f PresenterFunc) Present(errors errors.Errors) error {
	return f(errors)
}
