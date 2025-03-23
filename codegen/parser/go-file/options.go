package parser

import (
	"github.com/0B1t322/zero-validaton/codegen/matcher"
	"github.com/0B1t322/zero-validaton/codegen/parser/tags"
)

type Option func(*options)

type options struct {
	structMatcherBuilder matcher.StructMatcherBuilder
	tagsToParse          []string
}

func (o *options) buildStructMatcher() matcher.StructMatcher {
	if o.structMatcherBuilder == nil {
		return matcher.NewAlwaysTrueStructMatcher()
	}

	return o.structMatcherBuilder.Build()
}

func (o *options) buildTagsParser() tagParser {
	return tags.NewParser(o.tagsToParse...)
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

func WithTagsToParse(tagsToParse []string) Option {
	return func(o *options) {
		o.tagsToParse = tagsToParse
	}
}
