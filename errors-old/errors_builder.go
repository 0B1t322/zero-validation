package errors_old

type ErrorsBuilder interface {
	Set(key string, value error)
	Join(with Errors)
	Get() Errors
}
