package validate

import (
	"fmt"
	errors_v2 "github.com/0B1t322/zero-validaton/errors"
	"github.com/0B1t322/zero-validaton/translation"
	"github.com/0B1t322/zero-validaton/translation/locales/ru"
	validatecontext "github.com/0B1t322/zero-validaton/validate/context"
	"testing"
)

func TestTranslateError(t *testing.T) {
	t.Parallel()

	registry := translation.GlobalRegistry()
	registry.RegisterLocale(ru.Locale())
	err := errors_v2.NewErrorObject("required", "some")
	//ctx := context.Background()

	vCtx := validatecontext.New(registry, registry.DefaultLocale())

	t.Log(err)

	newErr := TranslateError(vCtx, err)
	t.Log(newErr)
	t.Log(err)
}

func BenchmarkTranslateError(b *testing.B) {
	type benchCase struct {
		name string
		bc   func(b *testing.B)
	}

	registry := translation.GlobalRegistry()

	bc := []benchCase{
		{
			name: "error is not translatable",
			bc: func(b *testing.B) {
				err := fmt.Errorf("test error")
				//ctx := context.Background()

				vCtx := validatecontext.New(registry, registry.DefaultLocale())

				b.RunParallel(func(pb *testing.PB) {
					for pb.Next() {
						_ = TranslateError(vCtx, err)
					}
				})
			},
		},
		{
			name: "error is translatable",
			bc: func(b *testing.B) {
				err := errors_v2.NewErrorObject("required", "some")
				//ctx := context.Background()

				vCtx := validatecontext.New(registry, registry.DefaultLocale())

				b.RunParallel(func(pb *testing.PB) {
					for pb.Next() {
						_ = TranslateError(vCtx, err)
					}
				})
			},
		},
	}

	for _, br := range bc {
		b.Run(br.name, br.bc)
	}
}
