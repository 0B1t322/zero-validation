package errors_old

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFieldErrorToErrors(t *testing.T) {
	t.Parallel()

	err1 := fmt.Errorf("err1")
	err2 := fmt.Errorf("err2")

	type testCase struct {
		fs   []*FieldError
		want Errors
	}

	testCases := []testCase{
		{
			fs: []*FieldError{
				{
					Field: NewStringField("a"),
					err:   err1,
				},
				{
					Field: NewStringField("b"),
					err:   err1,
				},
				{
					Field: NewStringField("c"),
					err:   err1,
				},
			},
			want: map[string]error{
				"a": err1,
				"b": err1,
				"c": err1,
			},
		},
		{
			fs: []*FieldError{
				{
					Field: NewStringField("a"),
					err:   err1,
				},
				{
					Field: NewStringField("b"),
					err:   err1,
				},
				{
					Field: NewStringField("c"),
					err:   err1,
				},
				{
					Field: NewStringField("c"),
					err:   err2,
				},
			},
			want: map[string]error{
				"a": err1,
				"b": err1,
				"c": ErrorSlice{err1, err2},
			},
		},
		{
			fs: []*FieldError{
				{
					Field: NewStringField("a"),
					err:   err1,
				},
				{
					Field: NewStringField("b"),
					err: append(
						NewFieldErrors(),
						NewFieldError(
							NewStringField("a"),
							err1,
						),
						NewFieldError(
							NewStringField("b"),
							err2,
						),
					),
				},
			},
			want: map[string]error{
				"a": err1,
				"b": Errors{
					"a": err1,
					"b": err2,
				},
			},
		},
		{
			fs: []*FieldError{
				{
					Field: NewStringField("a"),
					err:   err1,
				},
				{
					Field: NewStringField("b"),
					err: append(
						NewFieldErrors(),
						NewFieldError(
							NewStringField("a"),
							err1,
						),
						NewFieldError(
							NewStringField("b"),
							err2,
						),
					),
				},
				{
					Field: NewStringField("b"),
					err: append(
						NewFieldErrors(),
						NewFieldError(
							NewStringField("c"),
							err1,
						),
						NewFieldError(
							NewStringField("d"),
							err2,
						),
					),
				},
			},
			want: map[string]error{
				"a": err1,
				"b": Errors{
					"a": err1,
					"b": err2,
					"c": err1,
					"d": err2,
				},
			},
		},
	}

	for _, tt := range testCases {
		tt := tt
		t.Run(
			"case", func(t *testing.T) {
				t.Parallel()

				got := FieldErrorToErrors(tt.fs)
				assert.Equal(t, tt.want, got)
			},
		)
	}
}
