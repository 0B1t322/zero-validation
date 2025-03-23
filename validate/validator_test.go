package validate

import (
	"context"
	"github.com/0B1t322/zero-validaton/field"
	"github.com/0B1t322/zero-validaton/rule"
	"github.com/0B1t322/zero-validaton/translation"
	"testing"
)

type Object struct {
	ID   uint64
	Name string
}

type objectExtractor struct {
	ID   field.StructField[Object, uint64]
	Name field.StructField[Object, string]
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
}

func TestSome(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	obj := Object{}
	ctx = FieldNameGetterToContext(ctx, NewFieldNameGetterStrategy(
		"ru",
		FieldNameProto,
	))

	err := Struct(
		ctx,
		obj,
		Field(
			ExtractObject.ID,
			rule.Required[uint64](),
		),
		Field(
			ExtractObject.Name,
			rule.Required[string](),
		),
	)
	t.Log(err)
}

func BenchmarkTranslate(b *testing.B) {
	ctx := context.Background()
	ctx = ValidateContextToContext(ctx, NewValidateContext(
		translation.GlobalRegistry(),
		"en",
		WithFieldNameGetter(NewFieldNameGetterStrategy(
			"ru",
			FieldNameProto,
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
