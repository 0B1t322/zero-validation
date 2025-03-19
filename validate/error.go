package validate

import (
	errors_v2 "github.com/0B1t322/zero-validaton/errors"
)

func TranslateError(ctx Context, err error) error {
	if errObj, ok := err.(errors_v2.ErrorObject); ok {
		code := errObj.GetCode()
		registry := ctx.GetRegistry()

		tmpl, errGetTemplate := registry.GetErrorTemplate(code, ctx.GetPreferredLocale())
		if errGetTemplate != nil {
			return err
		}
		
		if tmpl == errObj.GetErrorTemplate() {
			return err
		}

		return errObj.SetErrorTemplate(tmpl)
	}

	return err
}
