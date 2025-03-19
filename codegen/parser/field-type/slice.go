package field_type

type Slice struct {
	Field FieldTyper
}

func (s Slice) Kind() Kind {
	return s.Field.Kind()
}

func (s Slice) Unwrap() FieldTyper {
	return s.Field
}

func (s Slice) Unwraps() []FieldTyper {
	return []FieldTyper{s.Unwrap()}
}

func (s Slice) GoTypeString() string {
	return "[]" + s.Field.GoTypeString()
}

func (s Slice) GoTypeStringWithAlias(alias string) string {
	return "[]" + s.Field.GoTypeStringWithAlias(alias)
}

func (s Slice) String() string {
	return s.GoTypeString()
}

func (s Slice) Accept(v Visitor) {
	v.VisitSlice(s)
	Visit(s.Field, v)
}

func SliceField(f FieldTyper) FieldTyper {
	return Slice{Field: f}
}
