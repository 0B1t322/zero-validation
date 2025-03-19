package generator

import (
	"github.com/0B1t322/zero-validaton/codegen/matcher"
	"github.com/0B1t322/zero-validaton/codegen/parser/go-file"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGenerator(t *testing.T) {
	const path = "/Users/danademin/Projects/Golang/zero-validation/internal/_testdata"

	p := parser.NewParser(
		parser.WithStructMatcherBuilder(
			matcher.
				NewBuilder().
				AddFullMatches("ToDo"),
		),
	)
	err := p.ParsePackage(path)
	require.NoError(t, err)

	g := NewGenerator()

	err = g.Generate(GenerateCommand{
		Structs:           p.ParsedStructs(),
		PackagePath:       "",
		DestinationPath:   "/Users/danademin/Projects/Golang/zero-validation/internal/_testdata/generated",
		GeneratedFileName: "validation_gen.go",
	})
	require.NoError(t, err)
}
