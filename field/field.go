package field

type (
	ValueExtractor[T any, V any] func(from T) V

	field[T any, V any] struct {
		name                 string
		additionalNameGetter additionalNameGetter
		extractor            ValueExtractor[T, V]
	}
)

func (f *field[T, V]) Name() string {
	return f.name
}

func (f *field[T, V]) GetAdditionalName(key string) string {
	return f.additionalNameGetter.GetAdditionalName(key)
}

func (f *field[T, V]) ExtractValue(from T) V {
	return f.extractor(from)
}

func FromPtr[T any, V any](s StructField[T, V]) StructField[*T, V] {
	return &field[*T, V]{
		name:                 s.Name(),
		additionalNameGetter: s,
		extractor: func(from *T) V {
			if from == nil {
				var def V
				return def
			}

			return s.ExtractValue(*from)
		},
	}
}

type StructField[T any, V any] interface {
	Name() string
	GetAdditionalName(key string) string
	ExtractValue(from T) V
}

type additionalNameGetter interface {
	GetAdditionalName(key string) string
}

type mapAdditionFieldGetter struct {
	defaultName     string
	additionalNames map[string]string
}

func (m mapAdditionFieldGetter) GetAdditionalName(key string) string {
	if len(m.additionalNames) == 0 {
		return m.defaultName
	}

	additionalName, find := m.additionalNames[key]
	if !find {
		return m.defaultName
	}

	return additionalName
}

func newMapAdditionFieldGetter(defaultName string, additionalNames map[string]string) mapAdditionFieldGetter {
	return mapAdditionFieldGetter{
		defaultName:     defaultName,
		additionalNames: additionalNames,
	}
}

func NewField[T any, V any](
	name string,
	additionalNames map[string]string,
	extractor ValueExtractor[T, V],
) *field[T, V] {
	return &field[T, V]{
		name:                 name,
		additionalNameGetter: newMapAdditionFieldGetter(name, additionalNames),
		extractor:            extractor,
	}
}
