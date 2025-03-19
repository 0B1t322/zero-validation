package ru

import (
	"github.com/0B1t322/zero-validaton/translation/locales/override"
	"text/template"
)

func Locale(templateByCodeOverrides ...override.TemplateByCode) locale {
	return locale{
		templateByCodeOverrides: templateByCodeOverrides,
	}
}

type locale struct {
	templateByCodeOverrides []override.TemplateByCode
}

func (l locale) GetCodesTemplate() map[string]*template.Template {
	templateByCode := map[string]*template.Template{
		"required": template.Must(template.New("required").Parse("поле обязательно для заполнения")),
	}

	override.Locale(templateByCode, l.templateByCodeOverrides...)

	return templateByCode
}

func (l locale) GetName() string {
	return "ru"
}
