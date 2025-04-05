package field_type

import (
	"github.com/0B1t322/zero-validation/internal/assertx"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFieldTyper_GoTypeString(t *testing.T) {
	t.Parallel()

	type testCase struct {
		name       string
		fieldTyper FieldTyper
		want       assert.ValueAssertionFunc
	}

	testCases := []testCase{
		{
			name:       "basic string",
			fieldTyper: BasicString,
			want:       assertx.Equal("string"),
		},
		{
			name:       "basic slice",
			fieldTyper: SliceField(BasicString),
			want:       assertx.Equal("[]string"),
		},
		{
			name:       "basic ptr",
			fieldTyper: PtrField(BasicString),
			want:       assertx.Equal("*string"),
		},
		{
			name:       "custom",
			fieldTyper: CustomField("Some", "", ""),
			want:       assertx.Equal("Some"),
		},
		{
			name:       "custom",
			fieldTyper: CustomField("Some", "some", "go/some"),
			want:       assertx.Equal("some.Some"),
		},
		{
			name:       "custom",
			fieldTyper: CustomField("Some", "", "go/some"),
			want:       assertx.Equal("some.Some"),
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			tt.want(t, tt.fieldTyper.GoTypeString())
		})
	}
}
