package codegen

import (
	"errors"
	"github.com/0B1t322/zero-validaton/codegen/config"
	"github.com/0B1t322/zero-validaton/codegen/generator"
	"github.com/0B1t322/zero-validaton/codegen/parser"
	"github.com/0B1t322/zero-validaton/codegen/parser/gqlgen"
	"github.com/0B1t322/zero-validaton/codegen/parser/tags"
	"github.com/99designs/gqlgen/codegen"
	gqlgenconfig "github.com/99designs/gqlgen/codegen/config"
	"go/types"
	"path"
	"strings"
)

type GqlGen struct {
	cfg          config.Config
	gqlGenConfig *gqlgenconfig.Config
	generator    *generator.Generator
}

func NewGqlGen(cfg config.Config, gqlGenConfig *gqlgenconfig.Config) *GqlGen {
	return &GqlGen{
		cfg:          cfg,
		gqlGenConfig: gqlGenConfig,
		generator:    generator.NewGenerator(generationOptionsFromConfig(cfg)...),
	}
}

func (g *GqlGen) Name() string {
	return "zero-validaton-gqlgen"
}

func (g *GqlGen) GenerateCode(data *codegen.Data) error {
	gqlgenParser := gqlgen.NewParser(tags.NewParser("json"))

	parsedStructsByPackageName := make(map[string][]parser.Struct, 1+len(data.Config.AutoBind))
	for _, inputObject := range data.Inputs {
		gqlgenParser.Reset()

		pkgPath, err := getPackagePathForGqlgenObject(inputObject)
		if err != nil {
			return err
		}

		err = gqlgenParser.ParseObject(pkgPath, inputObject)
		if err != nil {
			return err
		}

		parsedStructsByPackageName[pkgPath] = append(parsedStructsByPackageName[pkgPath], gqlgenParser.Structs()...)
	}

	for pkgPath, parsedStructs := range parsedStructsByPackageName {
		err := g.generator.Generate(generator.GenerateCommand{
			GenerateToCommand: generator.GenerateToCommand{
				Structs:                   parsedStructs,
				PackageName:               path.Base(pkgPath),
				IsGenerateInParsedPackage: true,
				ForceExtractFromPtr:       true,
			},
			PackagePath:       pkgPath,
			DestinationPath:   strings.ReplaceAll(pkgPath, g.cfg.GoModulePath(), g.cfg.BastPath()),
			GeneratedFileName: "zero_validation_extractors_gen.go",
		})
		if err != nil {
			return err
		}
	}

	return nil
}

func getPackagePathForGqlgenObject(obj *codegen.Object) (string, error) {
	namedType, ok := obj.Type.(*types.Named)
	if !ok {
		return "", errors.New("could not determine package name")
	}

	return namedType.Obj().Pkg().Path(), nil
}
