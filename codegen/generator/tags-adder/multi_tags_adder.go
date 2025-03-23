package tags_adder

import "github.com/0B1t322/zero-validaton/codegen/parser"

type MultiTagsAdder struct {
	TagsAdders []*TagsAdder
}

func NewMultiTagsAdder(tagsAdders []*TagsAdder) *MultiTagsAdder {
	return &MultiTagsAdder{
		TagsAdders: tagsAdders,
	}
}

func (m *MultiTagsAdder) AddTags(structs []parser.Struct) []parser.Struct {
	for _, tagsAdder := range m.TagsAdders {
		structs = tagsAdder.AddTags(structs)
	}

	return structs
}
