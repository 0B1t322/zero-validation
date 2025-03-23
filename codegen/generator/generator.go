package generator

import (
	"fmt"
	"github.com/0B1t322/zero-validaton/codegen/parser"
	"io"
	"net/url"
	"os"
	"path"
	"text/template"
)

type Generator struct {
	baseProjectPath string
	baseProjectName string

	tagsAdder tagsAdder
}

func NewGenerator(opts ...Option) *Generator {
	baseProjectPath, err := basePathFromOS()
	if err != nil {
		panic(fmt.Errorf("failed to get base project path: %w", err))
	}

	projectName, ok := getModuleName()
	if !ok {
		panic(fmt.Errorf("failed to get module name"))
	}

	opt := newOptions(opts...)

	g := &Generator{
		baseProjectPath: baseProjectPath,
		baseProjectName: projectName,
		tagsAdder:       opt.tagsAdder,
	}

	return g
}

// GenerateCommand ...
type GenerateCommand struct {
	GenerateToCommand
	// PackagePath path of parsed package
	PackagePath string
	// DestinationPath path of destination for generated files
	DestinationPath string
	// GeneratedFileName name of generated file
	GeneratedFileName string
}

type GenerateToCommand struct {
	Structs                   []parser.Struct
	PackageName               string
	IsGenerateInParsedPackage bool
	ForceExtractFromPtr       bool
	// PackageImportPath is import path to package
	PackageImportPath string
}

func (g *Generator) GenerateTo(cmd GenerateToCommand, w io.Writer) error {
	if len(cmd.Structs) == 0 {
		return nil
	}

	if g.tagsAdder != nil {
		cmd.Structs = g.tagsAdder.AddTags(cmd.Structs)
	}

	data := generateFileData{
		PkgName:                   path.Base(cmd.PackageName),
		Imports:                   parser.Structs.GetUsedImports(cmd.Structs),
		Structs:                   cmd.Structs,
		IsGenerateInParsedPackage: cmd.IsGenerateInParsedPackage,
	}

	if !cmd.IsGenerateInParsedPackage {
		data.Imports = append(data.Imports, parser.Import{
			Path:  cmd.PackageImportPath,
			Alias: cmd.parsedPackageAlias(),
		})
	}

	tmpl, err := g.createTemplate(cmd)
	if err != nil {
		return err
	}

	if err := tmpl.Execute(w, data); err != nil {
		return fmt.Errorf("failed execute template: %w", err)
	}

	return nil
}

func (g *Generator) Generate(cmd GenerateCommand) error {
	if len(cmd.Structs) == 0 {
		return nil
	}

	fileURL, err := url.JoinPath(cmd.DestinationPath, cmd.GeneratedFileName)
	if err != nil {
		return fmt.Errorf("failed url.JoinPath: %w", err)
	}

	file, err := createFileByURL(fileURL)
	if err != nil {
		return fmt.Errorf("failed create destination file: %w", err)
	}
	defer file.Close()

	if err := file.Truncate(0); err != nil {
		return err
	}

	return g.GenerateTo(cmd.GenerateToCommand, file)
}

func (g *Generator) createTemplate(cmd GenerateToCommand) (*template.Template, error) {
	tmpl, err := template.
		New("validationFile").
		Funcs(templateFunctions{
			isGenerateInParsedPackage: cmd.IsGenerateInParsedPackage,
			parsedPackageAlias:        cmd.parsedPackageAlias(),
			forceFromPtr:              cmd.ForceExtractFromPtr,
		}.FuncMap()).
		Parse(validationFileTemplate)

	if err != nil {
		return nil, err
	}

	return tmpl, nil
}

func (cmd GenerateToCommand) parsedPackageAlias() string {
	return "parsedPackage"
}

func createFileByURL(fileURL string) (*os.File, error) {
	err := os.MkdirAll(path.Dir(fileURL), os.ModePerm)
	if err != nil {
		switch {
		case os.IsExist(err):
		//	skip
		default:
			return nil, fmt.Errorf("failed create dir: %w", err)
		}
	}
	file, err := os.Create(fileURL)
	if err != nil {
		return nil, err
	}

	return file, nil
}
