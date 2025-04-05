package tags_adder

import (
	"github.com/0B1t322/zero-validation/codegen/matcher"
	"github.com/0B1t322/zero-validation/codegen/parser"
	"github.com/0B1t322/zero-validation/internal/slices"
)

type TagsAdder struct {
	structMatcher        matcher.StructMatcher
	fieldTagsByFieldName map[string]FieldTags
}

type FieldTags struct {
	FieldName  string
	ValueByTag map[string]string
}

func NewTagsAdder(structMatcher matcher.StructMatcher, fieldsTags []FieldTags) *TagsAdder {
	return &TagsAdder{
		structMatcher: structMatcher,
		fieldTagsByFieldName: slices.Associate(fieldsTags, func(item FieldTags) (string, FieldTags) {
			return item.FieldName, item
		}),
	}
}

type tagsAdderConfiguration interface {
	GetStructMatches() []string
	GetFieldTagsByField() map[string]map[string]string
}

func TagsAdderFromConfiguration(config tagsAdderConfiguration) *TagsAdder {
	fieldTagsByField := config.GetFieldTagsByField()

	fieldTags := make([]FieldTags, 0, len(fieldTagsByField))

	for fieldName, tags := range fieldTagsByField {
		fieldTags = append(fieldTags, FieldTags{
			FieldName:  fieldName,
			ValueByTag: tags,
		})
	}

	return NewTagsAdder(
		matcher.NewBuilder().AddRegexpMatches(config.GetStructMatches()...).Build(),
		fieldTags,
	)
}

func (tags *TagsAdder) IsMatchStruct(structName string) bool {
	return tags.structMatcher.Match(structName)
}

func (tags *TagsAdder) AddTags(structs []parser.Struct) []parser.Struct {
	for idx, st := range structs {
		if !tags.IsMatchStruct(st.Name) {
			continue
		}

		structs[idx] = tags.addTagsToStruct(st)
	}

	return structs
}

func (tags *TagsAdder) addTagsToStruct(st parser.Struct) parser.Struct {
	for idx, field := range st.Fields {
		fieldTags, isFind := tags.fieldTagsByFieldName[field.Name]
		if !isFind {
			continue
		}

		for key, value := range fieldTags.ValueByTag {
			field.Tags[key] = value
		}

		st.Fields[idx] = field
	}

	return st
}
