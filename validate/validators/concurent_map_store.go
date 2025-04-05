package validators

import (
	"github.com/0B1t322/zero-validaton/validate"
	"sync"
)

type concurrentMapStore struct {
	m sync.Map
}

func (c *concurrentMapStore) Get(key string) any {
	v, isFind := c.m.Load(key)
	if !isFind {
		return nil
	}

	return v
}

func (c *concurrentMapStore) Set(key string, value any) {
	c.m.Store(key, value)
}

func NewConcurrentMapStore() *concurrentMapStore {
	return &concurrentMapStore{}
}

func GetValidatorRules[V Validator[T], T any]() []validate.FieldRule[T] {
	return GetOrInitValidatorRulesFromStore[V, T](globalMapStore)
}

func InitValidatorRules[V Validator[T], T any](validator V) {
	InitValidatorInStore(globalMapStore, validator)
}

func GetOrInitValidatorRules[V Validator[T], T any]() []validate.FieldRule[T] {
	return GetOrInitValidatorRulesFromStore[V, T](globalMapStore)
}
