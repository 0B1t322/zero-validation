package errors_old

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrors_Error(t *testing.T) {
	err := Errors{
		"id":   NewError("f", "required"),
		"name": NewError("f", "required"),
	}

	assert.Equal(t, "id: required; name: required.", err.Error())
}

func TestErrors_Join(t *testing.T) {
	err := Errors{
		"id":   NewError("f", "required"),
		"name": NewError("f", "required"),
	}

	err = err.Join(
		Errors{
			"id": NewError("s", "nil"),
		},
	)
	err = err.Join(
		Errors{
			"id": NewError("d", "aaa"),
		},
	)
	data, errMarshal := err.MarshalJSON()
	assert.NoError(t, errMarshal)
	assert.Equal(t, `{"id":["required","nil","aaa"],"name":"required"}`, string(data))
}

func TestErrors_MarshalJSON(t *testing.T) {
	err := Errors{
		"id": Errors{
			"f": NewError("f", "required"),
			"s": NewError("f", "required"),
		},
	}

	data, errMarshal := err.MarshalJSON()
	assert.NoError(t, errMarshal)
	assert.Equal(t, `{"id":{"f":"required","s":"required"}}`, string(data))
}

func BenchmarkJoin(b *testing.B) {
	b.ResetTimer()
	b.RunParallel(
		func(p *testing.PB) {
			for p.Next() {
				errs1 := Errors{
					"id": NewError("1", "2"),
				}
				errs2 := Errors{
					"name": NewError("3", "4"),
				}
				errs3 := Errors{
					"color": NewError("5", "6"),
				}
				errs1.Join(errs2)
				errs1.Join(errs3)
			}
		},
	)
}
