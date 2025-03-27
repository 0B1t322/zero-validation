package matcher

import (
	"fmt"
	"regexp"
)

type StructMatcher interface {
	Match(structName string) bool
}

type structMatcher struct {
	regex *regexp.Regexp
}

func newStructMatcher(regexpString string) StructMatcher {
	if regexpString == "" {
		return &alwaysTrueStructMatcher{}
	}

	regex := regexp.MustCompile(regexpString)

	return &structMatcher{
		regex: regex,
	}
}

type structExcluder struct {
	excludeRegex *regexp.Regexp
}

func (s structExcluder) Match(structName string) bool {
	return !s.excludeRegex.MatchString(structName)
}

func (s structExcluder) String() string {
	return fmt.Sprintf("structExcluder: %v", s.excludeRegex.String())
}

func newStructExcluder(regexpString string) StructMatcher {
	if regexpString == "" {
		return &alwaysTrueStructMatcher{}
	}

	regex := regexp.MustCompile(regexpString)

	return &structExcluder{
		excludeRegex: regex,
	}
}

func (s *structMatcher) Match(structName string) bool {
	return s.regex.MatchString(structName)
}

type alwaysFalseStructMatcher struct{}

func (*alwaysFalseStructMatcher) Match(_ string) bool {
	return false
}

type alwaysTrueStructMatcher struct{}

func (*alwaysTrueStructMatcher) Match(_ string) bool { return true }

func NewAlwaysTrueStructMatcher() StructMatcher {
	return &alwaysTrueStructMatcher{}
}

type structMatcherDecorator struct {
	first  StructMatcher
	second StructMatcher
}

func (s *structMatcherDecorator) Match(structName string) bool {
	return s.first.Match(structName) && s.second.Match(structName)
}

func NewStructMatcherDecorator(first, second StructMatcher) StructMatcher {
	return &structMatcherDecorator{
		first:  first,
		second: second,
	}
}
