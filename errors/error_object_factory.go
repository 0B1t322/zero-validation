package errors

import (
	"fmt"
	"github.com/0B1t322/zero-validaton/translation"
	"text/template"
)

type ErrorObjectFactory interface {
	NewErrorObject(code string) ErrorObject
}

type translationRegistry interface {
	DefaultLocale() string
	GetErrorTemplate(code string, locale string) (*template.Template, error)
}

type defaultErrorObjectFactory struct {
	registry translationRegistry
}

func (d defaultErrorObjectFactory) NewErrorObject(code string) ErrorObject {
	locale := d.registry.DefaultLocale()
	tmpl, err := d.registry.GetErrorTemplate(code, locale)
	if err != nil {
		panic(fmt.Errorf("failed registry GetErrorTemplate: %w", err))
	}

	errObj := ErrorObject{
		code: code,
		tmpl: tmpl,
	}

	return errObj
}

var singleDefaultErrorObjectFactory = defaultErrorObjectFactory{
	registry: translation.GlobalRegistry(),
}

// DefaultErrorObjectFactory ...
func DefaultErrorObjectFactory() ErrorObjectFactory {
	return singleDefaultErrorObjectFactory
}

func NewErrorObjectFactory(registry translationRegistry) ErrorObjectFactory {
	return defaultErrorObjectFactory{
		registry: registry,
	}
}
