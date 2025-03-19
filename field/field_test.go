package field

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestField_ExtractValue(t *testing.T) {
	type object struct {
		Some string
	}

	f := NewField(
		"Some", nil, func(from *object) string {
			return from.Some
		},
	)

	v := f.ExtractValue(&object{Some: "s"})
	assert.Equal(t, v, "s")
}

func BenchmarkFiled_ExtractValue_ValueTypePointer(b *testing.B) {
	type object struct {
		Some string
	}

	f := NewField(
		"Some", nil, func(from *object) string {
			return from.Some
		},
	)

	obj := &object{Some: "s"}

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = f.ExtractValue(obj)
		}
	})
}

func BenchmarkFiled_ExtractValue_ValueTypeNotPointer(b *testing.B) {
	type object struct {
		Some string
	}

	f := NewField(
		"Some", nil, func(from object) string {
			return from.Some
		},
	)

	obj := object{Some: "s"}

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = f.ExtractValue(obj)
		}
	})
}
