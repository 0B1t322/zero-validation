package errors_old

type DirectErrors struct {
	errors Errors
}

func (d *DirectErrors) Get() Errors {
	return d.errors
}

func (d *DirectErrors) Set(key string, value error) {
	d.errors[key] = value
}

func (d *DirectErrors) Join(with Errors) {
	d.errors.Join(with)
}

func NewDirectErrors(startLen int) ErrorsBuilder {
	return &DirectErrors{
		errors: make(Errors, startLen),
	}
}
