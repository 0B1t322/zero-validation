package field_type

type Generic struct {
	GenericType   FieldTyper
	ParameterType FieldTyper
}

func (g Generic) Kind() Kind {
	return KindCustom
}

func (g Generic) Unwrap() FieldTyper {
	return g.ParameterType
}

func (g Generic) Unwraps() []FieldTyper {
	return []FieldTyper{
		g.GenericType,
		g.ParameterType,
	}
}

func (g Generic) GoTypeString() string {
	genericTypeName := g.GenericType.GoTypeString()

	return genericTypeName + "[" + g.ParameterType.GoTypeString() + "]"
}

func (g Generic) GoTypeStringWithAlias(alias string) string {
	genericTypeName := g.GenericType.GoTypeStringWithAlias(alias)

	return genericTypeName + "[" + g.ParameterType.GoTypeStringWithAlias(alias) + "]"
}

func (g Generic) String() string {
	return g.GoTypeString()
}

func (g Generic) Accept(visitor Visitor) {
	visitor.VisitGeneric(g)
	Visit(g.GenericType, visitor)
	Visit(g.ParameterType, visitor)
}

func GenericField(genericType FieldTyper, parameterType FieldTyper) FieldTyper {
	return Generic{
		GenericType:   genericType,
		ParameterType: parameterType,
	}
}
