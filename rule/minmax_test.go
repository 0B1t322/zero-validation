package rule

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMin(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		minRule := Min(10)
		assert.NoError(t, minRule.Validate(nil, 10))
		assert.NoError(t, minRule.Validate(nil, 11))
	})

	t.Run("failed", func(t *testing.T) {
		t.Parallel()

		minRule := Min(10)
		assert.Error(t, minRule.Validate(nil, 9))
		assert.Error(t, minRule.Validate(nil, 8))
	})
}

func TestMax(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		maxRule := Max(10)
		assert.NoError(t, maxRule.Validate(nil, 9))
		assert.NoError(t, maxRule.Validate(nil, 8))
	})

	t.Run("failed", func(t *testing.T) {
		t.Parallel()

		maxRule := Max(10)
		assert.Error(t, maxRule.Validate(nil, 11))
		assert.Error(t, maxRule.Validate(nil, 12))
	})
}

func TestBetween(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		betweenRule := Between(7, 10)
		assert.NoError(t, betweenRule.Validate(nil, 7))
		assert.NoError(t, betweenRule.Validate(nil, 8))
		assert.NoError(t, betweenRule.Validate(nil, 9))
		assert.NoError(t, betweenRule.Validate(nil, 10))
	})

	t.Run("failed", func(t *testing.T) {
		t.Parallel()

		betweenRule := Between(7, 10)
		assert.Error(t, betweenRule.Validate(nil, 5))
		assert.Error(t, betweenRule.Validate(nil, 6))
		assert.Error(t, betweenRule.Validate(nil, 11))
		assert.Error(t, betweenRule.Validate(nil, 12))
	})
}
