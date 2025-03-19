package translation

import "text/template"

// Registry ...
type Registry interface {
	// MustGetErrorTemplate return error template by code. Panic if not found template by code or locale
	MustGetErrorTemplate(code string, locale string) *template.Template
	// GetErrorTemplate return error template by code
	GetErrorTemplate(code string, locale string) (*template.Template, error)
	// DefaultLocale return default locale
	DefaultLocale() string
	// SetDefaultLocale set default locale
	SetDefaultLocale(locale string)
	RegisterTemplate(code string, locale string, template *template.Template)
	RegisterLocale(locale locale)
}

type locale interface {
	GetCodesTemplate() map[string]*template.Template
	GetName() string
}
