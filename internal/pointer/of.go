package pointer

func Of[T any](v T) *T {
	return &v
}

func GetValueOrDefault[T any](v *T) T {
	if v == nil {
		var def T
		return def
	}

	return *v
}
