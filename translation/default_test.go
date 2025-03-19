package translation

import (
	errors_v2 "github.com/0B1t322/zero-validaton/errors"
	"github.com/stretchr/testify/assert"
	"testing"
	"text/template"
)

type locale1 struct {
}

func (l locale1) GetCodesTemplate() map[string]*template.Template {
	return map[string]*template.Template{
		"code": template.Must(template.New("code").Parse("locale1")),
	}
}

func (l locale1) GetName() string {
	return "local1"
}

type locale2 struct {
}

func (l locale2) GetCodesTemplate() map[string]*template.Template {
	return map[string]*template.Template{
		"code": template.Must(template.New("code").Parse("locale2")),
	}
}

func (l locale2) GetName() string {
	return "local2"
}

func TestDefaultRegistry_SetDefaultLocale(t *testing.T) {
	t.Parallel()

	r := newDefaultRegistry(locale1{})

	factory := errors_v2.NewErrorObjectFactory(r)

	err := factory.NewErrorObject("code")

	assert.Equal(t, "locale1", err.Error())

	r.SetDefaultLocale("locale2")

	assert.Equal(t, "locale2", err.Error())
}
