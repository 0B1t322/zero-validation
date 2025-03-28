package validators

import (
	"github.com/0B1t322/zero-validaton/field"
	"github.com/0B1t322/zero-validaton/rule"
	"github.com/0B1t322/zero-validaton/validate"
	"testing"
)

type validator struct{}

func (v *validator) Name() string {
	return "validator"
}

func (v *validator) Rules() []validate.FieldRule[string] {
	f := field.NewField[string, string](
		"some field", nil, func(from string) string {
			return from
		},
	)
	return []validate.FieldRule[string]{
		validate.Field[string, string](
			f,
			rule.Required[string](),
		),
	}
}

type validatorTwo struct{}

func (v *validatorTwo) Name() string {
	return "validatorTwo"
}

func (v *validatorTwo) Rules() []validate.FieldRule[string] {
	f := field.NewField[string, string](
		"some field", nil, func(from string) string {
			return from
		},
	)

	return []validate.FieldRule[string]{
		validate.Field[string, string](
			f,
			rule.Required[string](),
		),
	}
}

func TestName(t *testing.T) {
	t.Parallel()
	//mapStore := NewDefaultMapStore()

	//rules := GetOrInitValidatorRulesFromStore[*validator](mapStore)
	//t.Log(rules)
	//m := sync.Map{}

}

func BenchmarkDefaultMapStore_Get(b *testing.B) {
	mapStore := NewDefaultMapStore()
	InitValidatorInStore(mapStore, &validator{})

	b.RunParallel(
		func(pb *testing.PB) {
			for pb.Next() {
				_ = GetValidatorRulesFromStore[*validator](mapStore)
			}
		},
	)
}
func BenchmarkName(b *testing.B) {
	m := NewConcurrentMapStore()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = m.Get("")
		}
	})
}

func BenchmarkName2(b *testing.B) {
	m := NewDefaultMapStore()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = m.Get("")
		}
	})
}
