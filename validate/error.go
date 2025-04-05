package validate

import (
	errors "github.com/0B1t322/zero-validaton/errors"
	validatecontext "github.com/0B1t322/zero-validaton/validate/context"
)

func TranslateError(ctx validatecontext.Context, err error) error {
	if errObj, ok := err.(errors.ErrorObject); ok {
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
