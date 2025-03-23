package validate

import "context"

type FieldNameGetterStrategy struct {
	fieldNameKeys []FieldNameKey
}

func NewFieldNameGetterStrategy(fieldNameKeys ...FieldNameKey) FieldNameGetter {
	return FieldNameGetterStrategy{fieldNameKeys: fieldNameKeys}
}

func (f FieldNameGetterStrategy) GetFieldName(fieldNameByKeyGetter fieldNameByKeyGetter) string {
	for _, key := range f.fieldNameKeys {
		fieldName, isOk := fieldNameByKeyGetter.TryGetAdditionalName(string(key))
		if isOk {
			return fieldName
		}
	}

	return fieldNameByKeyGetter.Name()
}

type fieldNameByKeyGetter interface {
	TryGetAdditionalName(key string) (string, bool)
	GetAdditionalName(key string) string
	Name() string
}

type FieldNameGetter interface {
	GetFieldName(fieldNameByKeyGetter fieldNameByKeyGetter) string
}

// FieldNameKey ...
type FieldNameKey string

const (
	// FieldNameProto ...
	FieldNameProto = FieldNameKey("proto")
	// FieldNameJSON ...
	FieldNameJSON = FieldNameKey("json")
)

func (f FieldNameKey) GetFieldName(fieldNameByKeyGetter fieldNameByKeyGetter) string {
	return fieldNameByKeyGetter.GetAdditionalName(string(f))
}

type fieldNameKey struct{}

func FieldNameGetterToContext(ctx context.Context, fieldName FieldNameGetter) context.Context {
	return context.WithValue(ctx, fieldNameKey{}, fieldName)
}

func FieldNameGetterFromContext(ctx context.Context) (FieldNameGetter, bool) {
	fieldName, ok := ctx.Value(fieldNameKey{}).(FieldNameGetter)
	return fieldName, ok
}
