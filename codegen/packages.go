package codegen

import (
	"fmt"
	"github.com/0B1t322/zero-validation/codegen/config"
	"github.com/0B1t322/zero-validation/codegen/generator"
	"github.com/0B1t322/zero-validation/codegen/matcher"
	packageparser "github.com/0B1t322/zero-validation/codegen/parser/go-file"
	"golang.org/x/tools/go/packages"
	"path"
	"strings"
)

type Packages struct {
	cfg       config.Config
	generator *generator.Generator
}

func NewPackages(cfg config.Config) *Packages {
	return &Packages{
		cfg: cfg,
		generator: generator.NewGenerator(
			generationOptionsFromConfig(cfg)...,
		),
	}
}

func (p *Packages) Generate() error {
	if len(p.cfg.PackagesConfig) == 0 {
		return nil
	}

	generateCommandByDestinationPath := make(map[string]generator.GenerateCommand, len(p.cfg.PackagesConfig))

	for pkgPath, pkgCfg := range p.cfg.PackagesConfig {
		pkgTruePath, err := p.getPkgTruePath(pkgPath)
		if err != nil {
			return err
		}

		packagesParser := packageparser.NewParser(
			packageparser.WithTagsToParse([]string{"json"}),
			packageparser.WithStructMatcherBuilder(
				structMatcherByPackageConfig(pkgCfg),
			),
		)

		err = packagesParser.ParsePackage(pkgTruePath)
		if err != nil {
			return fmt.Errorf("parsing package: %w", err)
		}

		dstPath := p.getDstPath(pkgCfg.Dst, pkgTruePath)

		cmd, isFind := generateCommandByDestinationPath[dstPath]
		if !isFind {
			cmd = generator.GenerateCommand{
				GenerateToCommand: generator.GenerateToCommand{
					PackageName:               path.Base(pkgPath),
					PackageImportPath:         pkgPath,
					IsGenerateInParsedPackage: pkgTruePath == dstPath,
				},
				PackagePath:       pkgPath,
				DestinationPath:   dstPath,
				GeneratedFileName: "zero_validation_extractors_gen.go",
			}
		}

		cmd.Structs = append(cmd.Structs, packagesParser.ParsedStructs()...)
		generateCommandByDestinationPath[dstPath] = cmd
	}

	for _, generateCommand := range generateCommandByDestinationPath {
		err := p.generator.Generate(generateCommand)
		if err != nil {
			return err
		}
	}

	return nil
}

func (p *Packages) getPkgTruePath(pkgPath string) (string, error) {
	if strings.Contains(pkgPath, p.cfg.GoModulePath()) {
		return strings.ReplaceAll(pkgPath, p.cfg.GoModulePath(), p.cfg.BastPath()), nil
	}

	cfg := &packages.Config{
		Mode: packages.NeedModule | packages.NeedName,
	}
	packages, err := packages.Load(cfg, pkgPath)
	if err != nil {
		return "", err
	}

	if len(packages) != 1 {
		return "", fmt.Errorf("not found package %s", pkgPath)
	}

	pkg := packages[0]

	if pkg.Module == nil {
		return "", fmt.Errorf("not found Module for %s", pkgPath)
	}

	dir := pkg.Module.Dir

	truePath := strings.ReplaceAll(pkgPath, pkg.PkgPath, dir)

	return truePath, nil
}

func splitModulePathAndVersion(modulePath string) (string, string) {
	indexOfVersion := strings.LastIndex(modulePath, "@")
	if indexOfVersion == -1 {
		return modulePath, ""
	}

	return modulePath[:indexOfVersion], modulePath[indexOfVersion+1:]
}

func (p *Packages) getDstPath(dstPath string, parsedPackageTruePath string) string {
	if dstPath == "" {
		return parsedPackageTruePath
	}

	return path.Join(p.cfg.BastPath(), dstPath)
}

func structMatcherByPackageConfig(cfg config.PackageConfig) matcher.StructMatcherBuilder {
	builder := matcher.NewBuilder()

	if len(cfg.Structs.Include) > 0 {
		builder = builder.AddRegexpMatches(cfg.Structs.Include...)
	}

	if len(cfg.Structs.Exclude) > 0 {
		builder = builder.AddRegexpExcludes(cfg.Structs.Exclude...)
	}

	return builder
}
