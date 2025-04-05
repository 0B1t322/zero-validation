package errors

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestErrors_Is(t *testing.T) {
	t.Parallel()

	type args struct {
		err error
	}

	type testCase struct {
		name   string
		err    error
		wantOk assert.BoolAssertionFunc
	}

	testCases := []testCase{
		{
			name:   "direct true",
			err:    Errors{},
			wantOk: assert.True,
		},

		{
			name:   "wrapped true",
			err:    fmt.Errorf("wrapper: %w", Errors{}),
			wantOk: assert.True,
		},

		{
			name:   "wrapped two times true",
			err:    fmt.Errorf("failed: %w", fmt.Errorf("wrapper: %w", Errors{})),
			wantOk: assert.True,
		},
		{
			name:   "another error",
			err:    fmt.Errorf("another errror"),
			wantOk: assert.False,
		},
		{
			name:   "nil",
			err:    nil,
			wantOk: assert.False,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			tt.wantOk(t, errors.Is(tt.err, Errors{}))
		})
	}
}
