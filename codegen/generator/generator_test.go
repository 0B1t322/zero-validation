package generator

import (
	"github.com/0B1t322/zero-validation/codegen/parser"
	field_type "github.com/0B1t322/zero-validation/codegen/parser/field-type"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGenerator(t *testing.T) {
	const path = "/Users/danademin/Projects/Golang/zero-validation/internal/_testdata"
	g := NewGenerator()
	err := g.Generate(GenerateCommand{
		GenerateToCommand: GenerateToCommand{
			Structs: []parser.Struct{
				{
					Name: "Entity",
					Fields: []parser.Field{
						{
							Name: "Name",
							Type: field_type.BasicString,
						},
						{
							Name: "AnotherEntity",
							Type: field_type.SliceField(field_type.Custom{Name: "AnotherEntity"}),
						},
					},
				},
				{
					Name: "AnotherEntity",
					Fields: []parser.Field{
						{
							Name: "ID",
							Type: field_type.BasicUint64,
						},
					},
				},
			},
			PackageName:               "generated",
			IsGenerateInParsedPackage: true,
		},
		PackagePath:       "",
		DestinationPath:   path,
		GeneratedFileName: "validation_gen.go",
	})
	require.NoError(t, err)
}
