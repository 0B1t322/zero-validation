package errors

import (
	"slices"
	"strings"
	"sync"
)

var (
	fieldErrorPool = &sync.Pool{
		New: func() any {
			return &FieldError{}
		},
	}

	fieldErrorsPool = &sync.Pool{
		New: func() any {
			return make(FieldErrors, 0)
		},
	}
)

type FieldError struct {
	Field Field
	err   error
}

func (f *FieldError) Error() string {
	return f.err.Error()
}

type FieldErrors []*FieldError

func (fes FieldErrors) Error() string {
	if len(fes) == 0 {
		return ""
	}

	b := strings.Builder{}

	for i, fe := range fes {
		_, _ = b.WriteString(fe.Error())

		if i != len(fes)-1 {
			_, _ = b.WriteString("\n")
		}
	}

	return b.String()
}

func NewFieldError(field Field, err error) *FieldError {
	fieldErr := fieldErrorPool.Get().(*FieldError)
	fieldErr.Field = field
	fieldErr.err = err
	return fieldErr
}

type newFieldErrorsOpt struct {
	startCap int
}

type NewFieldErrorsOpt func(opt *newFieldErrorsOpt)

func WithStartCap(startCap int) NewFieldErrorsOpt {
	return func(opt *newFieldErrorsOpt) {
		opt.startCap = startCap
	}
}

func NewFieldErrors(opts ...NewFieldErrorsOpt) FieldErrors {
	fes := fieldErrorsPool.Get().(FieldErrors)

	o := &newFieldErrorsOpt{}
	for _, opt := range opts {
		opt(o)
	}

	if cap(fes) < o.startCap {
		fes = slices.Grow(fes, o.startCap)
	}

	return fes
}

func ReleaseFieldErrors(fe FieldErrors) {
	fe = fe[:0]
	fieldErrorsPool.Put(fe)
}

func ReleaseFieldError(fe *FieldError) {
	fe.Field = nil
	fe.err = nil
	fieldErrorPool.Put(fe)
}

func FieldErrorToErrors(fs FieldErrors) Errors {
	errs := make(Errors, len(fs))
	for _, fe := range fs {
		field := fe.Field.GetName()

		err := fe.err
		if fes, ok := err.(FieldErrors); ok {
			err = FieldErrorToErrors(fes)
		}

		if findedErr, find := errs[field]; find {
			errs[field] = makeNewError(findedErr, err)
			ReleaseFieldError(fe)
			continue
		}

		errs[field] = err
		ReleaseFieldError(fe)
	}

	ReleaseFieldErrors(fs)
	return errs
}

func makeNewError(err1 error, err2 error) error {
	switch err := err1.(type) {
	case Errors:
		return tryJoinErrors(err, err2)
	case ErrorSlice:
		return err.AppendError(err2)
	}

	switch err := err2.(type) {
	case Errors:
		return tryJoinErrors(err, err1)
	case ErrorSlice:
		s := make(ErrorSlice, 0, len(err)+1)
		s = append(s, err1)
		s = append(s, err...)
		return s
	}

	return ErrorSlice{err1, err2}
}

func tryJoinErrors(errs Errors, err error) Errors {
	if err, ok := err.(Errors); ok {
		return errs.Join(err)
	}
	return errs
}
