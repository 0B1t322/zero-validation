package en

import "text/template"

func Locale() locale {
	return locale{}
}

type locale struct{}

func (l locale) GetCodesTemplate() map[string]*template.Template {
	return map[string]*template.Template{
		"required": template.Must(template.New("required").Parse("field is required")),
	}
}

func (l locale) GetName() string {
	return "en"
}
