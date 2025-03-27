package fieldname

import "context"

type ByKeyGetter interface {
	TryGetAdditionalName(key string) (string, bool)
	GetAdditionalName(key string) string
	Name() string
}

type Getter interface {
	GetFieldName(fieldNameByKeyGetter ByKeyGetter) string
}

type GetterStrategy struct {
	fieldNameKeys []Key
}

func NewGetterStrategy(fieldNameKeys ...Key) Getter {
	return GetterStrategy{fieldNameKeys: fieldNameKeys}
}

func (f GetterStrategy) GetFieldName(fieldNameByKeyGetter ByKeyGetter) string {
	for _, key := range f.fieldNameKeys {
		fieldName, isOk := fieldNameByKeyGetter.TryGetAdditionalName(string(key))
		if isOk {
			return fieldName
		}
	}

	return fieldNameByKeyGetter.Name()
}

// Key ...
type Key string

const (
	// Proto ...
	Proto = Key("proto")
	// JSON ...
	JSON = Key("json")
)

func (f Key) GetFieldName(fieldNameByKeyGetter ByKeyGetter) string {
	return fieldNameByKeyGetter.GetAdditionalName(string(f))
}

type fieldNameKey struct{}

func GetterToContext(ctx context.Context, fieldName ByKeyGetter) context.Context {
	return context.WithValue(ctx, fieldNameKey{}, fieldName)
}

func GetterFromContext(ctx context.Context) (Getter, bool) {
	fieldName, ok := ctx.Value(fieldNameKey{}).(Getter)
	return fieldName, ok
}
