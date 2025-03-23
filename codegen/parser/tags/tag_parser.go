package tags

import (
	"reflect"
	"strings"
)

// tagOptions is the string following a comma in a struct field's "json"
// tag, or the empty string. It does not include the leading comma.
type tagOptions string

// parseTag splits a struct field's json tag into its name and
// comma-separated options.
func parseTag(tag string) (string, tagOptions) {
	tag, opt, _ := strings.Cut(tag, ",")
	return tag, tagOptions(opt)
}

// Contains reports whether a comma-separated list of options
// contains a particular substr flag. substr must be surrounded by a
// string boundary or commas.
func (o tagOptions) Contains(optionName string) bool {
	if len(o) == 0 {
		return false
	}
	s := string(o)
	for s != "" {
		var name string
		name, s, _ = strings.Cut(s, ",")
		if name == optionName {
			return true
		}
	}
	return false
}

type Parser struct {
	knownTags []string
}

func NewParser(tags ...string) *Parser {
	return &Parser{
		knownTags: tags,
	}
}

func (p *Parser) ParseTag(tag string) map[string]string {
	parsed := make(map[string]string, len(p.knownTags))
	structTag := reflect.StructTag(tag)

	for _, knownTag := range p.knownTags {
		tagValue := structTag.Get(knownTag)
		if tagValue == "" {
			continue
		}

		value, _ := parseTag(tagValue)
		parsed[knownTag] = value
	}

	if len(parsed) == 0 {
		return nil
	}

	return parsed
}
