package translationx

import (
	"github.com/0B1t322/zero-validaton/translation"
	"github.com/0B1t322/zero-validaton/translation/locales/ru"
)

func init() {
	translation.GlobalRegistry().RegisterLocale(ru.Locale())
	translation.GlobalRegistry().SetDefaultLocale("ru")
}
