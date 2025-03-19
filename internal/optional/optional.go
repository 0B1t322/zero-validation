package optional

type Optional[T any] interface {
	Value() (T, bool)
	MustValue() T
	Set(value T)
	IsPresent() bool
	ValueOrDefault(def T) T
}

type optional[T any] struct {
	value     T
	isPresent bool
}

func (o *optional[T]) Value() (T, bool) {
	return o.value, o.isPresent
}

func (o optional[T]) ValueOrDefault(def T) T {
	if o.isPresent {
		return o.value
	}

	return def
}

func (o *optional[T]) MustValue() T {
	return o.value
}

func (o *optional[T]) Set(value T) {
	o.value = value
	o.isPresent = true
}

func (o *optional[T]) IsPresent() bool {
	return o.isPresent
}

func Default[T any]() Optional[T] {
	return &optional[T]{isPresent: false}
}

func New[T any](v T) Optional[T] {
	return &optional[T]{
		value:     v,
		isPresent: true,
	}
}
