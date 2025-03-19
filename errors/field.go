package errors

type Field interface {
	GetName() string
	SetNameKey(key string)
}

type stringField struct {
	field string
}

func (s *stringField) GetName() string {
	return s.field
}

func (s *stringField) SetNameKey(_ string) {}

func NewStringField(field string) Field {
	return &stringField{
		field: field,
	}
}
