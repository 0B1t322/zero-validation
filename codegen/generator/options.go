package generator

import (
	"github.com/0B1t322/zero-validaton/codegen/parser"
)

type Option func(o *options)

type options struct {
	tagsAdder tagsAdder
}

func newOptions(opts ...Option) *options {
	o := &options{}

	for _, opt := range opts {
		opt(o)
	}

	return o
}

type tagsAdder interface {
	AddTags(structs []parser.Struct) []parser.Struct
}

// WithTagsAdder ...
func WithTagsAdder(tagsAdder tagsAdder) Option {
	return func(o *options) {
		o.tagsAdder = tagsAdder
	}
}
