package errors_old

import (
	"encoding/json"
	"strings"
	"sync"
	"text/template"
)

var errStringBuilderPool = sync.Pool{
	New: func() any {
		return new(strings.Builder)
	},
}

type ErrorObject struct {
	code   string
	params any

	tmpl *template.Template
}

// NewErrorObject ...
func NewErrorObject(code string, message string) ErrorObject {
	e := ErrorObject{
		code: code,
		tmpl: template.Must(template.New("err").Parse(message)),
	}

	return e
}

func (e ErrorObject) GetCode() string {
	return e.code
}

func (e ErrorObject) SetErrorTemplate(newTemplate *template.Template) error {
	e.tmpl = newTemplate
	return e
}

func (e ErrorObject) GetErrorTemplate() *template.Template {
	return e.tmpl
}

func (e ErrorObject) Error() string {
	// TODO: Maybe pool here
	b := strings.Builder{}
	_ = e.tmpl.Execute(&b, e.params)

	text := b.String()

	b.Reset()

	errStringBuilderPool.Put(b)

	return text
}

func (e ErrorObject) SetParams(params any) ErrorObject {
	e.params = params
	return e
}

func (e ErrorObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(e.Error())
}
