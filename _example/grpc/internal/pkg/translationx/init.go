package translationx

import (
	"github.com/0B1t322/zero-validation/translation"
	"github.com/0B1t322/zero-validation/translation/locales/ru"
)

func init() {
	translation.GlobalRegistry().RegisterLocale(ru.Locale())
	translation.GlobalRegistry().SetDefaultLocale("ru")
}
