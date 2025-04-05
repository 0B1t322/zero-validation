package errors

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func BenchmarkFieldErrorToErrors(b *testing.B) {
	err1 := fmt.Errorf("error1")
	err2 := fmt.Errorf("error2")
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			fs := NewFieldErrors(WithStartCap(3))

			fs = append(fs, NewFieldError("a", err1))
			fs = append(fs, NewFieldError("b", err2))
			fs = append(fs, NewFieldError("c", err1))

			_ = FieldErrorToErrors(fs)
		}
	})
}

func TestFieldErrorToErrors_SimpleCase(t *testing.T) {
	t.Parallel()
	err1 := fmt.Errorf("error1")
	err2 := fmt.Errorf("error2")

	fs := NewFieldErrors(WithStartCap(5))

	fs = append(fs, NewFieldError("a", err1))
	fs = append(fs, NewFieldError("b", err2))
	fs = append(fs, NewFieldError("a", err2))

	errs := FieldErrorToErrors(fs)
	assert.Equal(t, errs["a"], ErrorSlice{err1, err2})
	assert.Equal(t, errs["b"], ErrorSlice{err2})
	t.Log(errs)
}
