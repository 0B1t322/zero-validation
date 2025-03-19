package validators

import (
	"github.com/0B1t322/zero-validaton/validate"
)

type ValidatorStore interface {
	Get(key string) any
	Set(key string, v any)
}

func GetValidatorRules[V Validator[T], T any](store ValidatorStore) []validate.FieldRule[T] {
	name := getValidatorName[T, V]()
	wrap := store.Get(name).(Validator[T])
	return wrap.Rules()
}

// InitValidator init validator into store
func InitValidator[V Validator[T], T any](
	store ValidatorStore,
	validator V,
) {
	wrapped := Wrap(validator)

	store.Set(getValidatorName[T, V](), wrapped)
}

func getValidatorName[T any, V Validator[T]]() string {
	var def V
	return def.Name()
}
