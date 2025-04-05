package rule

import (
	"github.com/0B1t322/zero-validation/translation"
	"github.com/0B1t322/zero-validation/translation/locales/ru"
	"testing"
)

func TestRequired(t *testing.T) {
	t.Parallel()

	type testCase struct {
		name string
	}

	testCases := []testCase{}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

		})
	}
}

func TestSome(t *testing.T) {
	t.Parallel()

	err := Required[int]().Validate(0)
	t.Log(err)
}

func BenchmarkName(b *testing.B) {
	translation.GlobalRegistry().RegisterLocale(ru.Locale())
	req := Required[int]()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			err := req.Validate(0)
			//err = err.(interface {
			//	Translate(registry translation.Registry, locale string) errors-old.ErrorObject
			//}).Translate(translation.GlobalRegistry(), "ru")
			_ = err.Error()
		}
	})
}
