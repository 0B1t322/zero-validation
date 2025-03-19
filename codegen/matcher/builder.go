package matcher

import "strings"

type StructMatcherBuilder interface {
	AddRegexpMatches(matches ...string) StructMatcherBuilder
	AddRegexpExcludes(excludes ...string) StructMatcherBuilder
	//AddFullMatches(matches ...string) StructMatcherBuilder
	Build() StructMatcher
}

type structMatcherBuilder struct {
	matchStrBuilder    *strings.Builder
	notMatchStrBuilder *strings.Builder
}

func NewBuilder() StructMatcherBuilder {
	return &structMatcherBuilder{
		matchStrBuilder:    &strings.Builder{},
		notMatchStrBuilder: &strings.Builder{},
	}
}

func (s *structMatcherBuilder) AddRegexpMatches(matches ...string) StructMatcherBuilder {
	if len(matches) == 0 {
		return s
	}

	if s.matchStrBuilder.Len() > 0 {
		s.matchStrBuilder.WriteString("|")
	}

	for i, match := range matches {
		s.matchStrBuilder.WriteString(match)
		if i < len(matches)-1 {
			s.matchStrBuilder.WriteString("|")
		}
	}

	return s
}

func (s *structMatcherBuilder) AddRegexpExcludes(excludes ...string) StructMatcherBuilder {
	if len(excludes) == 0 {
		return s
	}

	if s.notMatchStrBuilder.Len() > 0 {
		s.notMatchStrBuilder.WriteString("|")
	}

	for i, match := range excludes {
		s.notMatchStrBuilder.WriteString(match)
		if i < len(excludes)-1 {
			s.notMatchStrBuilder.WriteString("|")
		}
	}

	return s
}

func (s *structMatcherBuilder) AddFullMatches(matches ...string) StructMatcherBuilder {
	if len(matches) == 0 {
		return s
	}

	if s.matchStrBuilder.Len() > 0 {
		s.matchStrBuilder.WriteString("|")
	}

	for i, match := range matches {
		s.matchStrBuilder.WriteString("^")
		s.matchStrBuilder.WriteString(match)
		s.matchStrBuilder.WriteString("$")

		if i < len(matches)-1 {
			s.matchStrBuilder.WriteString("|")
		}
	}

	return s
}

func (s *structMatcherBuilder) Build() StructMatcher {
	matcher := newStructMatcher(s.matchStrBuilder.String())
	excluder := newStructExcluder(s.notMatchStrBuilder.String())

	return NewStructMatcherDecorator(matcher, excluder)
}
