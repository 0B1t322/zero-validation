package translationx

import (
	"github.com/0B1t322/zero-validation/translation"
	"github.com/0B1t322/zero-validation/translation/locales/override"
	"github.com/0B1t322/zero-validation/translation/locales/ru"
)

var registry = translation.GlobalRegistry()

func init() {
	overrides := []override.TemplateByCode{
		override.NewTemplateByCodeFromText("required", "обязательное поле"),
	}

	registry.RegisterLocale(
		ru.Locale(overrides...),
	)

	registry.SetDefaultLocale("ru")
}
