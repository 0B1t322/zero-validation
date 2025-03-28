package en

import "text/template"

func Locale() locale {
	return locale{}
}

type locale struct{}

func (l locale) GetCodesTemplate() map[string]*template.Template {
	return map[string]*template.Template{
		"required":                          template.Must(template.New("required").Parse("field is required")),
		"validation_in_invalid":             template.Must(template.New("validation_in_invalid").Parse("must be in {{.In}}")),
		"validation_value_gte_min":          template.Must(template.New("validation_value_gte_min").Parse("value must be greater or equal then {{.Min}}")),
		"validation_value_lte_max":          template.Must(template.New("validation_value_lte_max").Parse("value must be less or equal then {{.Max}}")),
		"validation_value_between_required": template.Must(template.New("validation_value_between_required").Parse("value bust be between {{.Min}} and {{.Max}}")),
		"validation_not_in_invalid":         template.Must(template.New("validation_not_in_invalid").Parse("must not be in {{.In}}")),
	}
}

func (l locale) GetName() string {
	return "en"
}
