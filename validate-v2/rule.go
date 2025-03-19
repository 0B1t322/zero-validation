package validate

type Rule[T any] interface {
	Validate(v T) error
	SetErrorLocale(locale string)
}
