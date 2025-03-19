package field_type

import (
	"testing"

	"github.com/0B1t322/zero-validaton/internal/pointer"
	"github.com/stretchr/testify/assert"
)

func TestIsOrdered(t *testing.T) {
	t.Parallel()
	t.Skip()

	type testCase struct {
		fieldTyper       FieldTyper
		targets          []any
		want             bool
		targetsAssertion func(t *testing.T, targets []any)
	}

	_ = []testCase{
		{
			fieldTyper: PtrField(SliceField(BasicInt)),
			targets:    []any{&Ptr{}, &Slice{}, pointer.Of(Basic(""))},
			want:       true,
			targetsAssertion: func(t *testing.T, targets []any) {
				ptr := targets[0].(*Ptr)
				slice := targets[1].(*Slice)
				basic := targets[2].(*Basic)

				assert.NotNil(t, ptr.Field)
				assert.NotNil(t, slice.Field)
				assert.Equal(t, *basic, BasicInt)
			},
		},
		{
			fieldTyper: PtrField(SliceField(BasicInt)),
			targets:    []any{&Ptr{}, pointer.Of(Basic(""))},
			want:       false,
			targetsAssertion: func(t *testing.T, targets []any) {
				ptr := targets[0].(*Ptr)
				basic := targets[1].(*Basic)

				assert.NotNil(t, ptr.Field)
				assert.NotEqual(t, *basic, BasicInt)
			},
		},
		{
			fieldTyper: PtrField(SliceField(BasicInt)),
			targets:    []any{&Slice{}, pointer.Of(Basic(""))},
			want:       false,
			targetsAssertion: func(t *testing.T, targets []any) {
				slice := targets[0].(*Slice)
				basic := targets[1].(*Basic)

				assert.Nil(t, slice.Field)
				assert.NotEqual(t, *basic, BasicInt)
			},
		},
		{
			fieldTyper: PtrField(GenericField(CustomFiled("Some", ""), BasicInt)),
			targets:    []any{&Generic{}, &Custom{}, pointer.Of(Basic(""))},
			want:       false,
			targetsAssertion: func(t *testing.T, targets []any) {
				generic := targets[0].(*Generic)
				custom := targets[1].(*Custom)
				basic := targets[2].(*Basic)

				assert.Nil(t, generic.GenericType)
				assert.Nil(t, generic.ParameterType)
				assert.Empty(t, custom)
				assert.NotEqual(t, *basic, BasicInt)
			},
		},
	}

	//for _, tt := range tts {
	//	//tt := tt
	//	t.Run("case", func(t *testing.T) {
	//		t.Parallel()
	//
	//		//got := IsOrdered(tt.fieldTyper, tt.targets...)
	//		//assert.Equal(t, tt.want, got)
	//		//tt.targetsAssertion(t, tt.targets)
	//	})
	//}
}
