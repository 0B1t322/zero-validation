package errors_old

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrorObject_Error(t *testing.T) {
	t.Parallel()

	type testCase struct {
		errorObject ErrorObject
		wantErr     string
	}

	testCases := []testCase{
		{
			errorObject: NewError(
				"error_1",
				"msg with params {{.Param}}",
			).SetParams(&struct{ Param int }{Param: 12}),
			wantErr: "msg with params 12",
		},
	}

	for i, tt := range testCases {
		tt := tt
		t.Run(
			fmt.Sprintf("%d", i), func(t *testing.T) {
				t.Parallel()
				assert.Equal(t, tt.wantErr, tt.errorObject.Error())
			},
		)
	}
}
