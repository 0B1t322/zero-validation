package generator

import (
	_ "embed"
	"fmt"
	"github.com/0B1t322/zero-validaton/codegen/parser"
	"strings"
	"text/template"
	"unicode"
	"unicode/utf8"
)

//go:embed validation_file.go.tmpl
var validationFileTemplate string

type templateFunctions struct {
	isGenerateInParsedPackage bool
	parsedPackageAlias        string
	forceFromPtr              bool
}

func (t templateFunctions) FuncMap() template.FuncMap {
	return map[string]any{
		"extractorTypeName":     t.extractorTypeName,
		"formatExtractorFunc":   t.formatExtractorFunc,
		"formatField":           t.formatField,
		"formatAdditionalNames": t.formatAdditionalNames,
	}
}

func (templateFunctions) extractorTypeName(str string) string {
	r, size := utf8.DecodeRuneInString(str)
	if r == utf8.RuneError && size <= 1 {
		return str
	}
	lc := unicode.ToLower(r)
	if r == lc {
		return str
	}
	return string(lc) + str[size:]

}

func (t templateFunctions) formatExtractorFunc(s parser.Struct, f parser.Field) string {
	parentType := t.formatParentType(s)
	var returnType string
	switch t.isGenerateInParsedPackage {
	case true:
		returnType = f.Type.GoTypeString()
	case false:
		returnType = f.Type.GoTypeStringWithAlias(t.parsedPackageAlias)
	}

	return fmt.Sprintf(
		"func(from %s) %s { return from.%s }",
		parentType,
		returnType,
		f.Name,
	)
}

func (t templateFunctions) formatParentType(s parser.Struct) string {
	typeName := s.Name

	if !t.isGenerateInParsedPackage {
		typeName = t.parsedPackageAlias + "." + typeName
	}

	if t.forceFromPtr {
		typeName = "*" + typeName
	}

	return typeName
}

func (t templateFunctions) formatChildFieldType(field parser.Field) string {
	if t.isGenerateInParsedPackage {
		return field.Type.GoTypeString()
	}

	//fmt.Fprintf(os.Stderr,,"%+v\n", field.Type.GoTypeStringWithAlias(t.parsedPackageAlias))
	return field.Type.GoTypeStringWithAlias(t.parsedPackageAlias)
}

func (t templateFunctions) formatField(s parser.Struct, f parser.Field) string {
	parentType := t.formatParentType(s)
	childType := t.formatChildFieldType(f)

	return fmt.Sprintf("%s field.StructField[%s,%s]", f.Name, parentType, childType)
}

func (templateFunctions) formatAdditionalNames(f parser.Field) string {
	if len(f.Tags) == 0 {
		return "nil"
	}

	mapBuilder := strings.Builder{}

	mapBuilder.WriteString("map[string]string{")
	idx := 0
	for key, value := range f.Tags {
		mapBuilder.WriteString("\"")
		mapBuilder.WriteString(key)
		mapBuilder.WriteString("\":")
		mapBuilder.WriteString("\"")
		mapBuilder.WriteString(value)
		mapBuilder.WriteString("\"")
		if idx != len(f.Tags)-1 {
			mapBuilder.WriteString(",")
		}
		idx++
	}
	mapBuilder.WriteString("}")

	return mapBuilder.String()
}
