package field_type

type Interface struct{}

func (i Interface) Kind() Kind {
	return KindInterface
}

func (i Interface) Unwrap() FieldTyper {
	return nil
}

func InterfaceField() Interface {
	return Interface{}
}
