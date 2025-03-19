package parser

import "github.com/0B1t322/zero-validaton/codegen/matcher"

type Option func(o *options)

type options struct {
	excludes []string
}

func WithExcludes(excludes []string) Option {
	return func(o *options) {
		o.excludes = append(o.excludes, excludes...)
	}
}

func newOptions(opts ...Option) *options {
	o := &options{}

	for _, opt := range opts {
		opt(o)
	}

	return o
}

func (o *options) structMatcher() matcher.StructMatcher {
	b := matcher.NewBuilder()

	b = b.AddRegexpExcludes(o.excludes...)

	return b.Build()
}
