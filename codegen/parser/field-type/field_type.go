package field_type

type Kind uint

const (
	KindBasic Kind = iota
	KindStruct
	KindCustom
	KindInterface
)

type FieldTyper interface {
	Kind() Kind
	Unwraps() []FieldTyper
	GoTypeString() string
	GoTypeStringWithAlias(alias string) string
	Accept(v Visitor)
}
