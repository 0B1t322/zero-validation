package validators

import (
	"github.com/0B1t322/zero-validaton/validate"
)

type ValidatorStore interface {
	Get(key string) any
	Set(key string, v any)
}

func GetValidatorRulesFromStore[V Validator[T], T any](store ValidatorStore) []validate.FieldRule[T] {
	name := getValidatorName[T, V]()
	wrap := store.Get(name).(Validator[T])
	return wrap.Rules()
}

// InitValidatorInStore init validator into store
func InitValidatorInStore[V Validator[T], T any](
	store ValidatorStore,
	validator V,
) {
	wrapped := Wrap(validator)

	store.Set(getValidatorName[T, V](), wrapped)
}

func GetOrInitValidatorRulesFromStore[V Validator[T], T any](store ValidatorStore) []validate.FieldRule[T] {
	name := getValidatorName[T, V]()
	wrap, isOk := store.Get(name).(Validator[T])
	if !isOk {
		var def V
		store.Set(name, def)
		wrap = def
	}
	return wrap.Rules()
}

func getValidatorName[T any, V Validator[T]]() string {
	var def V
	return def.Name()
}
