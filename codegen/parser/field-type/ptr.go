package field_type

type Ptr struct {
	Field FieldTyper
}

func (p Ptr) Kind() Kind {
	return p.Field.Kind()
}

func (p Ptr) Unwrap() FieldTyper {
	return p.Field
}

func (p Ptr) Unwraps() []FieldTyper {
	return []FieldTyper{p.Unwrap()}
}

func (p Ptr) GoTypeString() string {
	return "*" + p.Field.GoTypeString()
}

func (p Ptr) GoTypeStringWithAlias(alias string) string {
	return "*" + p.Field.GoTypeStringWithAlias(alias)
}

func (p Ptr) String() string {
	return p.GoTypeString()
}

func (p Ptr) Accept(v Visitor) {
	v.VisitPtr(p)
	Visit(p.Field, v)
}

func PtrField(f FieldTyper) FieldTyper {
	return Ptr{Field: f}
}
