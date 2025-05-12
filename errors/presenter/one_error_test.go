package errorspresenter

import (
	"fmt"
	"github.com/0B1t322/zero-validation/errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPresentErrorsAsSimpleOneError(t *testing.T) {
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
			name: "empty errors",
			args: args{},
			want: assert.NoError,
		},
		{
			name: "one depth",
			args: args{
				errs: errors.Errors{
					"some": fmt.Errorf("some error"),
				},
			},
			want: func(t assert.TestingT, err error, _ ...interface{}) bool {
				return assert.Equal(t, err.Error(), "some: some error")
			},
		},
		{
			name: "one depth slice",
			args: args{
				errs: errors.Errors{
					"some": errors.ErrorSlice{fmt.Errorf("some error")},
				},
			},
			want: func(t assert.TestingT, err error, _ ...interface{}) bool {
				return assert.Equal(t, err.Error(), "some: some error")
			},
		},
		{
			name: "one depth slice with idx",
			args: args{
				errs: errors.Errors{
					"some": errors.ErrorSlice{
						errors.Errors{
							"0": fmt.Errorf("some error"),
						},
					},
				},
			},
			want: func(t assert.TestingT, err error, _ ...interface{}) bool {
				return assert.Equal(t, "some: some error", err.Error())
			},
		},
		{
			name: "two depth",
			args: args{
				errs: errors.Errors{
					"some": errors.Errors{
						"id": fmt.Errorf("some error"),
					},
				},
			},
			want: func(t assert.TestingT, err error, _ ...interface{}) bool {
				return assert.Equal(t, err.Error(), "id: some error")
			},
		},
		{
			name: "two depth",
			args: args{
				errs: errors.Errors{
					"some": errors.Errors{
						"id": fmt.Errorf("some error"),
					},
				},
			},
			want: func(t assert.TestingT, err error, _ ...interface{}) bool {
				return assert.Equal(t, err.Error(), "id: some error")
			},
		},
		{
			name: "3 depth",
			args: args{
				errs: errors.Errors{
					"some": errors.Errors{
						"entity": errors.Errors{
							"id": fmt.Errorf("some error"),
						},
					},
				},
			},
			want: func(t assert.TestingT, err error, _ ...interface{}) bool {
				return assert.Equal(t, err.Error(), "id: some error")
			},
		},
		{
			name: "3 depth with error slice in end",
			args: args{
				errs: errors.Errors{
					"some": errors.Errors{
						"entity": errors.Errors{
							"id": errors.ErrorSlice{
								fmt.Errorf("some error"),
							},
						},
					},
				},
			},
			want: func(t assert.TestingT, err error, _ ...interface{}) bool {
				return assert.Equal(t, err.Error(), "id: some error")
			},
		},
		{
			name: "3 depth with error slice in end",
			args: args{
				errs: errors.Errors{
					"some": errors.Errors{
						"entity": errors.Errors{
							"id": errors.ErrorSlice{
								errors.Errors{
									"0": fmt.Errorf("some error"),
								},
							},
						},
					},
				},
			},
			want: func(t assert.TestingT, err error, _ ...interface{}) bool {
				return assert.Equal(t, err.Error(), "id: some error")
			},
		},
		{
			name: "3 depth with error slice in entity",
			args: args{
				errs: errors.Errors{
					"some": errors.Errors{
						"entity": errors.ErrorSlice{
							errors.Errors{
								"0": errors.Errors{
									"id": fmt.Errorf("some error"),
								},
							},
						},
					},
				},
			},
			want: func(t assert.TestingT, err error, _ ...interface{}) bool {
				return assert.Equal(t, err.Error(), "id: some error")
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := PresentErrorsAsSimpleOneError(tt.args.errs)
			tt.want(t, got)
		})
	}
}
