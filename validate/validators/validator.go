package validators

import (
	"github.com/0B1t322/zero-validaton/validate"
	"sync"
)

type Validator[T any] interface {
	Name() string
	Rules() []validate.FieldRule[T]
}

type ValidatorWrapper[T any] struct {
	Validator[T]
	rules []validate.FieldRule[T]
	once  sync.Once
}

func Wrap[V Validator[T], T any](validator V) Validator[T] {
	wrapper := &ValidatorWrapper[T]{
		Validator: validator,
		once:      sync.Once{},
	}

	return wrapper
}

func (v *ValidatorWrapper[T]) Rules() []validate.FieldRule[T] {
	v.once.Do(
		func() {
			v.rules = v.Validator.Rules()
		},
	)

	return v.rules
}
