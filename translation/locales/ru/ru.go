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
		"required":                          template.Must(template.New("required").Parse("поле обязательно для заполнения")),
		"validation_in_invalid":             template.Must(template.New("validation_in_invalid").Parse("значение должно быть одним из {{.In}}")),
		"validation_value_gte_min":          template.Must(template.New("validation_value_gte_min").Parse("значение должно быть больше или равно {{.Min}}")),
		"validation_value_lte_max":          template.Must(template.New("validation_value_lte_max").Parse("значение должно быть меньше или равно {{.Max}}")),
		"validation_value_between_required": template.Must(template.New("validation_value_between_required").Parse("значение должно быть между {{.Min}} и {{.Max}}")),
		"validation_not_in_invalid":         template.Must(template.New("validation_not_in_invalid").Parse("значение не должно быть одним из {{.In}}")),
	}

	override.Locale(templateByCode, l.templateByCodeOverrides...)

	return templateByCode
}

func (l locale) GetName() string {
	return "ru"
}
