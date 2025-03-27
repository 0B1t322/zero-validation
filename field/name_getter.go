package field

type fieldNameByKeyGetter interface {
	TryGetAdditionalName(key string) (string, bool)
	GetAdditionalName(key string) string
	Name() string
}

type FieldNameGetter interface {
	GetFieldName(fieldNameByKeyGetter fieldNameByKeyGetter) string
}
