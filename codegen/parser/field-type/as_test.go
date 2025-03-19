package field_type

import (
	"fmt"
	"testing"

	"github.com/0B1t322/zero-validaton/internal/pointer"
	"github.com/stretchr/testify/assert"
)

func TestAs(t *testing.T) {
	t.Parallel()

	type testCase struct {
		fieldTyper      FieldTyper
		target          any
		want            bool
		targetAssertion func(t *testing.T, target any)
	}

	testCases := []testCase{
		{
			fieldTyper: BasicInt,
			target:     pointer.Of[Basic](""),
			want:       true,
			targetAssertion: func(t *testing.T, target any) {
				basic, ok := target.(*Basic)
				assert.True(t, ok)
				assert.True(t, *basic == BasicInt)
			},
		},
		{
			fieldTyper: PtrField(BasicInt),
			target:     &Slice{},
			want:       false,
			targetAssertion: func(t *testing.T, target any) {
				s, ok := target.(*Slice)
				assert.True(t, ok)
				assert.Nil(t, s.Field)
			},
		},
	}

	for idx, tt := range testCases {
		tt := tt
		t.Run(fmt.Sprintf("case_%v", idx), func(t *testing.T) {
			t.Parallel()

			got := As(tt.fieldTyper, tt.target)
			assert.Equal(t, tt.want, got)
			tt.targetAssertion(t, tt.target)
		})
	}
}
