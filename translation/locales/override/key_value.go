package override

import "text/template"

type templateByCode struct {
	code string
	tmpl *template.Template
}

// NewTemplateByCode ...
func NewTemplateByCode(code string, tmpl *template.Template) TemplateByCode {
	return templateByCode{
		code: code,
		tmpl: tmpl,
	}
}

func (t templateByCode) GetCode() string {
	return t.code
}

func (t templateByCode) GetTemplate() *template.Template {
	return t.tmpl
}

// TemplateByCode ...
type TemplateByCode interface {
	GetCode() string
	GetTemplate() *template.Template
}

// Locale ...
func Locale(locale map[string]*template.Template, overrides ...TemplateByCode) {
	for _, override := range overrides {
		locale[override.GetCode()] = override.GetTemplate()
	}
}
