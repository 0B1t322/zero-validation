package translation

import (
	"fmt"
	"github.com/0B1t322/zero-validaton/translation/locales/en"
	"sync/atomic"
	"text/template"
)

func GlobalRegistry() Registry {
	return globalRegistry
}

func SetGlobalRegistry(registry Registry) {
	globalRegistry = registry
}

var globalRegistry = NewDefaultRegistry()

func NewDefaultRegistry() Registry {
	return newDefaultRegistry(en.Locale())
}

func newDefaultRegistry(defaultLocale locale) *defaultRegistry {
	r := &defaultRegistry{
		registryMap: make(map[string]map[string]*template.Template),
	}

	r.defaultLocale.Store(defaultLocale.GetName())
	r.RegisterLocale(defaultLocale)

	return r
}

type defaultRegistry struct {
	// code -> locale -> template
	registryMap map[string]map[string]*template.Template

	defaultLocale atomic.Value
}

var singletonDefaultRegistry Registry = NewDefaultRegistry()

func (d *defaultRegistry) MustGetErrorTemplate(code string, locale string) *template.Template {
	localedTemplate, err := d.GetErrorTemplate(code, locale)
	if err != nil {
		panic(err)
	}

	return localedTemplate
}

func (d *defaultRegistry) GetErrorTemplate(code string, locale string) (*template.Template, error) {
	templateLocales, isCodeRegistered := d.registryMap[code]
	if !isCodeRegistered {
		return nil, fmt.Errorf("code %s is not registered", code)
	}

	localedTemplate, isTemplateLocaleRegistered := templateLocales[locale]
	if !isTemplateLocaleRegistered {
		return nil, fmt.Errorf("locale %s is not registered for code %s", locale, code)
	}

	return localedTemplate, nil
}

func (d *defaultRegistry) RegisterTemplate(code string, locale string, tmpl *template.Template) {
	templateLocales, isCodeRegisterd := d.registryMap[code]
	if !isCodeRegisterd {
		templateLocales = make(map[string]*template.Template)
		d.registryMap[code] = templateLocales
	}

	templateLocales[locale] = tmpl
}

func (d *defaultRegistry) DefaultLocale() string {
	return d.defaultLocale.Load().(string)
}

func (d *defaultRegistry) SetDefaultLocale(locale string) {
	d.defaultLocale.Swap(locale)
}

func (d *defaultRegistry) RegisterLocale(locale locale) {
	codesTemplate := locale.GetCodesTemplate()
	localeName := locale.GetName()
	for code, tmpl := range codesTemplate {
		templateLocales, isCodeRegistered := d.registryMap[code]
		if !isCodeRegistered {
			templateLocales = make(map[string]*template.Template)
			d.registryMap[code] = templateLocales
		}

		templateLocales[localeName] = tmpl
	}

}
