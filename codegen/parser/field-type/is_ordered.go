package field_type

//func IsOrdered(fieldTyper FieldTyper, targets ...any) bool {
//	if fieldTyper == nil {
//		return false
//	}
//
//	if len(targets) == 0 {
//		return false
//	}
//
//	for idx, target := range targets {
//		val := reflect.ValueOf(target)
//		typ := val.Type()
//		if typ.Kind() != reflect.Ptr || val.IsNil() {
//			panic("FieldTyperIsOrdered: target must be a non-nil pointer")
//		}
//
//		targetType := typ.Elem()
//		if targetType.Kind() != reflect.Interface && !targetType.Implements(fieldTyperType) {
//			panic("FieldTyperIsOrdered: *target must be interface or implement error")
//		}
//
//		if reflect.TypeOf(fieldTyper).AssignableTo(targetType) {
//			val.Elem().Set(reflect.ValueOf(fieldTyper))
//		} else {
//			return false
//		}
//
//		fieldTyper = fieldTyper.Unwrap()
//		if fieldTyper == nil && idx != len(targets)-1 {
//			return false
//		}
//	}
//
//	return true
//}
