package codegen

import (
	"fmt"
	"github.com/0B1t322/zero-validaton/codegen/config"
	"github.com/0B1t322/zero-validaton/codegen/generator"
	model "github.com/0B1t322/zero-validaton/codegen/parser"
	parser "github.com/0B1t322/zero-validaton/codegen/parser/proto"
	"google.golang.org/protobuf/compiler/protogen"
)

type GRPC struct {
	cfg       config.Config
	generator *generator.Generator
	gen       *protogen.Plugin
}

func NewGRPC(gen *protogen.Plugin, cfg config.Config) *GRPC {
	return &GRPC{
		cfg:       cfg,
		generator: generator.NewGenerator(),
		gen:       gen,
	}
}

func (g *GRPC) Generate() error {
	protoParser := parser.NewParser(
		parser.WithExcludes(g.cfg.GrpcConfig.Exclude),
	)
	for _, file := range g.gen.Files {
		if !file.Generate {
			continue
		}

		protoParser.Reset()

		err := protoParser.ParseFile(file)
		if err != nil {
			return fmt.Errorf("parsing proto file: %w", err)
		}

		structs := protoParser.Structs()
		if len(structs) == 0 {
			continue
		}

		imports := model.Structs.GetUsedImports(structs)

		generatedFile := g.gen.NewGeneratedFile(file.GeneratedFilenamePrefix+"_extractors.pb.go", file.GoImportPath)

		for _, imp := range imports {
			impFile, isFind := g.gen.FilesByPath[imp.Path]
			if !isFind {
				continue
			}

			if impFile.GoImportPath == file.GoImportPath {
				continue
			}

			generatedFile.Import(impFile.GoImportPath)
		}

		err = g.generator.GenerateTo(
			generator.GenerateToCommand{
				Structs:                   structs,
				PackageName:               string(file.GoPackageName),
				IsGenerateInParsedPackage: true,
				ForceExtractFromPtr:       true,
			},
			generatedFile,
		)
		if err != nil {
			return fmt.Errorf("generating proto file: %w", err)
		}
	}

	return nil
}
