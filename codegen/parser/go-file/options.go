package parser

import "github.com/0B1t322/zero-validaton/codegen/matcher"

type Option func(*options)

type options struct {
	structMatcherBuilder matcher.StructMatcherBuilder
}

func (o *options) buildStructMatcher() matcher.StructMatcher {
	if o.structMatcherBuilder == nil {
		return matcher.NewAlwaysTrueStructMatcher()
	}

	return o.structMatcherBuilder.Build()
}

func newOptions(opts ...Option) *options {
	p := &options{}

	for _, o := range opts {
		o(p)
	}

	return p
}

// WithStructMatcherBuilder ...
func WithStructMatcherBuilder(builder matcher.StructMatcherBuilder) Option {
	return func(o *options) {
		o.structMatcherBuilder = builder
	}
}
