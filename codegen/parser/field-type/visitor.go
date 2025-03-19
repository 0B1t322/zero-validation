package field_type

// Visitor ...
type Visitor interface {
	VisitBasic(basic Basic)
	VisitCustom(custom Custom)
	VisitGeneric(generic Generic)
	VisitPtr(ptr Ptr)
	VisitSlice(slice Slice)
}

// VisitAll ...
func VisitAll(fieldTypes []FieldTyper, visitor Visitor) {
	for _, fieldType := range fieldTypes {
		Visit(fieldType, visitor)
	}
}

// Visit ...
func Visit(fieldType FieldTyper, visitor Visitor) {
	fieldType.Accept(visitor)
}
