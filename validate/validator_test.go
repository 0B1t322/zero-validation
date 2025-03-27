package validate

import (
	"context"
	"github.com/0B1t322/zero-validaton/field"
	fieldname "github.com/0B1t322/zero-validaton/field/name"
	"github.com/0B1t322/zero-validaton/rule"
	"github.com/0B1t322/zero-validaton/translation"
	validatecontext "github.com/0B1t322/zero-validaton/validate/context"
	"testing"
)

type Object struct {
	ID   uint64
	Name string
	Some someInterface
}

type someInterface interface {
	isSomeInterface()
}

type someImpl1 struct {
	Value uint64
}

type someImpl1Extractor struct {
	Value field.StructField[someImpl1, uint64]
}

var SomeImpl1 = someImpl1Extractor{
	Value: field.NewField("Value", nil, func(from someImpl1) uint64 {
		return from.Value
	}),
}

type someImpl2 struct {
	AnotherValue uint64
}

func (someImpl1) isSomeInterface() {}

func (someImpl2) isSomeInterface() {}

type objectExtractor struct {
	ID   field.StructField[Object, uint64]
	Name field.StructField[Object, string]
	Some field.StructField[Object, someInterface]
}

var ExtractObject = objectExtractor{
	ID: field.NewField("ID", map[string]string{
		"ru":    "ID Объекта",
		"proto": "id",
	}, func(from Object) uint64 {
		return from.ID
	}),
	Name: field.NewField("Name", map[string]string{
		"proto": "name",
	}, func(from Object) string {
		return from.Name
	}),
	Some: field.NewField("Some", nil, func(from Object) someInterface {
		return from.Some
	}),
}

func TestSome(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	obj := Object{
		Name: "special",
	}
	vCtx := validatecontext.New(
		translation.GlobalRegistry(),
		"ru",
		validatecontext.WithFieldNameGetter(fieldname.NewGetterStrategy(
			"ru",
			fieldname.Proto,
		)),
		validatecontext.WithStopAfterFirstError(),
	)

	ctx = validatecontext.ToContext(ctx, vCtx)

	err := Struct(
		ctx,
		obj,
		If(Object.IsNameNotSpecial,
			Field(
				ExtractObject.ID,
				rule.Required[uint64](),
			),
		),
		IfFieldTypeOf[someImpl1](
			ExtractObject.Some,
			Field(
				SomeImpl1.Value,
				rule.Required[uint64](),
			),
		),
		Field(
			ExtractObject.Name,
			rule.Required[string](),
		),
	)
	t.Log(err)
}

func (o Object) IsIdEval() bool {
	return o.ID%2 == 0
}

func (o Object) IsNameNotSpecial() bool {
	return o.Name != "special"
}

func BenchmarkTranslate(b *testing.B) {
	ctx := context.Background()
	ctx = validatecontext.ToContext(ctx, validatecontext.New(
		translation.GlobalRegistry(),
		"en",
		validatecontext.WithFieldNameGetter(fieldname.NewGetterStrategy(
			"ru",
			fieldname.Proto,
		)),
	))

	obj := Object{
		//ID:   1,
		//Name: "Some",
	}

	fieldRules := []FieldRule[Object]{
		Field(
			ExtractObject.ID,
			rule.Required[uint64](),
		),
		Field(
			ExtractObject.Name,
			rule.Required[string](),
		),
	}

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = Struct(ctx, obj, fieldRules...)
		}
	})
}
