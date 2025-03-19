package field_type

type Basic string

const (
	BasicString  = Basic("string")
	BasicInt     = Basic("int")
	BasicInt32   = Basic("int32")
	BasicInt64   = Basic("int64")
	BasicUint    = Basic("uint")
	BasicUint32  = Basic("uint32")
	BasicUint64  = Basic("uint64")
	BasicFloat32 = Basic("float32")
	BasicFloat64 = Basic("float64")
	BasicBool    = Basic("bool")
	BasicByte    = Basic("byte")
)

func (b Basic) GoName() string {
	return string(b)
}

func (Basic) Kind() Kind {
	return KindBasic
}

//func (Basic) Unwrap() FieldTyper {
//	return nil
//}

func (Basic) Unwraps() []FieldTyper {
	return nil
}

func (b Basic) GoTypeString() string {
	return string(b)
}

func (b Basic) GoTypeStringWithAlias(_ string) string {
	return b.GoTypeString()
}

func (b Basic) String() string {
	return b.GoTypeString()
}

func (b Basic) Accept(v Visitor) {
	v.VisitBasic(b)
}

var basicTypes = map[Basic]struct{}{
	BasicString:  {},
	BasicInt:     {},
	BasicInt32:   {},
	BasicInt64:   {},
	BasicUint:    {},
	BasicUint32:  {},
	BasicUint64:  {},
	BasicFloat32: {},
	BasicFloat64: {},
	BasicBool:    {},
	BasicByte:    {},
}

func ParseBasic(t string) (Basic, bool) {
	b := Basic(t)
	_, ok := basicTypes[b]
	return b, ok
}
