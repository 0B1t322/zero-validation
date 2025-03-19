package field_type

type Struct struct{}

func (s Struct) Kind() Kind {
	return KindStruct
}

func (s Struct) Unwrap() FieldTyper {
	return nil
}

func (s Struct) Unwraps() []FieldTyper {
	return nil
}

func StructField() Struct {
	return Struct{}
}
