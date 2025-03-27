package matcher

import "strings"

type StructMatcherBuilder interface {
	AddRegexpMatches(matches ...string) StructMatcherBuilder
	AddRegexpExcludes(excludes ...string) StructMatcherBuilder
	Build() StructMatcher
}

type structMatcherBuilder struct {
	matchStrBuilder    *strings.Builder
	notMatchStrBuilder *strings.Builder
}

func NewBuilder() StructMatcherBuilder {
	return &structMatcherBuilder{matchStrBuilder: &strings.Builder{}, notMatchStrBuilder: &strings.Builder{}}
}

func (s *structMatcherBuilder) AddRegexpMatches(matches ...string) StructMatcherBuilder {
	matchesToStringBuilder(s.matchStrBuilder, matches...)
	return s
}

func (s *structMatcherBuilder) AddRegexpExcludes(excludes ...string) StructMatcherBuilder {
	matchesToStringBuilder(s.notMatchStrBuilder, excludes...)
	return s
}

func (s *structMatcherBuilder) Build() StructMatcher {
	matcher := newStructMatcher(s.matchStrBuilder.String())
	excluder := newStructExcluder(s.notMatchStrBuilder.String())
	return NewStructMatcherDecorator(matcher, excluder)
}

func matchesToStringBuilder(stringBuilder *strings.Builder, matches ...string) {
	if len(matches) == 0 {
		return
	}

	if stringBuilder.Len() > 0 {
		stringBuilder.WriteString("|")
	}

	for i, match := range matches {
		stringBuilder.WriteString(match)
		if i < len(matches)-1 {
			stringBuilder.WriteString("|")
		}
	}
}
