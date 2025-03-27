package assertx

import "github.com/stretchr/testify/assert"

func Equal[T any](a T) assert.ValueAssertionFunc {
	return func(t assert.TestingT, i interface{}, _ ...interface{}) bool {
		return assert.Equal(t, a, i)
	}
}
