package translationx

import (
	"github.com/0B1t322/zero-validation/translation"
	"github.com/0B1t322/zero-validation/translation/locales/override"
	"github.com/0B1t322/zero-validation/translation/locales/ru"
)

func init() {
	ruLocale := ru.Locale(
		override.NewTemplateByCodeFromText("custom-code", "текст ошибки"),
	)
	translation.GlobalRegistry().RegisterLocale(ruLocale)
	translation.GlobalRegistry().SetDefaultLocale("ru")
}
