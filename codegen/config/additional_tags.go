package config

type AdditionalTags struct {
	Matches   []string   `yaml:"matches"`
	FieldTags []FieldTag `yaml:"field_tags"`
}

type FieldTag struct {
	FieldName   string            `yaml:"field_name"`
	ValueByTags map[string]string `yaml:"tags"`
}

func (a AdditionalTags) GetStructMatches() []string {
	return a.Matches
}

func (a AdditionalTags) GetFieldTagsByField() map[string]map[string]string {
	fieldTagsByField := make(map[string]map[string]string, len(a.FieldTags))

	for _, fieldTag := range a.FieldTags {
		fieldTagsByField[fieldTag.FieldName] = fieldTag.ValueByTags
	}

	return fieldTagsByField
}
