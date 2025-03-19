package field_type

import "reflect"

var fieldTyperType = reflect.TypeOf((*FieldTyper)(nil)).Elem()

func As(fieldTyper FieldTyper, target any) bool {
	if fieldTyper == nil {
		return false
	}

	if target == nil {
		panic("FieldTyperAs: target cannot be nil")
	}

	val := reflect.ValueOf(target)
	typ := val.Type()
	if typ.Kind() != reflect.Ptr || val.IsNil() {
		panic("FieldTyperAs: target must be a non-nil pointer")
	}

	targetType := typ.Elem()
	if targetType.Kind() != reflect.Interface && !targetType.Implements(fieldTyperType) {
		panic("FieldTyperAs: *target must be interface or implement error")
	}

	stack := []FieldTyper{fieldTyper}
	for {
		nextFieldTyper := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if reflect.TypeOf(nextFieldTyper).AssignableTo(targetType) {
			val.Elem().Set(reflect.ValueOf(nextFieldTyper))
			return true
		}

		nextFieldTypers := nextFieldTyper.Unwraps()
		if len(nextFieldTypers) == 0 {
			return false
		}

		stack = append(stack, nextFieldTypers...)
	}
}
