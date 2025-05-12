package errorspresenter

import (
	"fmt"
	"github.com/0B1t322/zero-validation/errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPresentErrorAsOnePathError(t *testing.T) {
	t.Parallel()

	type args struct {
		errs errors.Errors
	}

	type testCase struct {
		name string
		args args
		want assert.ErrorAssertionFunc
	}

	tests := []testCase{
		{
			name: "simple one error",
			args: args{
				errs: errors.Errors{
					"id": fmt.Errorf("some error"),
				},
			},
			want: func(t assert.TestingT, err error, _ ...interface{}) bool {
				return assert.Equal(t, "id: some error", err.Error())
			},
		},
		{
			name: "simple 2 depth error",
			args: args{
				errs: errors.Errors{
					"entity": errors.Errors{
						"id": fmt.Errorf("some error"),
					},
				},
			},
			want: func(t assert.TestingT, err error, _ ...interface{}) bool {
				return assert.Equal(t, "entity.id: some error", err.Error())
			},
		},
		{
			name: "simple 3 depth error",
			args: args{
				errs: errors.Errors{
					"entity": errors.Errors{
						"body": errors.Errors{
							"id": fmt.Errorf("some error"),
						},
					},
				},
			},
			want: func(t assert.TestingT, err error, _ ...interface{}) bool {
				return assert.Equal(t, "entity.body.id: some error", err.Error())
			},
		},
		{
			name: "simple 2 depth error with slice",
			args: args{
				errs: errors.Errors{
					"entity": errors.ErrorSlice{
						errors.Errors{
							"0": errors.Errors{
								"id": fmt.Errorf("some error"),
							},
						},
					},
				},
			},
			want: func(t assert.TestingT, err error, _ ...interface{}) bool {
				return assert.Equal(t, "entity[0].id: some error", err.Error())
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			err := PresentErrorAsOnePathError(tt.args.errs)

			tt.want(t, err)
		})
	}
}
