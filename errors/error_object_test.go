package errors

import (
	"testing"
	"text/template"
)

func BenchmarkErrorObject_Error(b *testing.B) {
	e := NewErrorObject("code", "msg")
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = e.Error()
		}
	})
}

func BenchmarkErrorObject_SetErrorTemplate(b *testing.B) {
	e := NewErrorObject("code", "msg")
	newTmpl := template.Must(template.New("").Parse("new"))
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			newE := e.SetErrorTemplate(newTmpl)
			_ = newE
		}
	})
}
