package translationx

import (
	"github.com/0B1t322/zero-validaton/translation"
	"github.com/0B1t322/zero-validaton/translation/locales/override"
	"github.com/0B1t322/zero-validaton/translation/locales/ru"
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
