package parser

import (
	"fmt"
	model "github.com/0B1t322/zero-validaton/codegen/parser"
	"go/ast"
	"go/parser"
	"go/token"
	"net/url"
	"os"
)

// Parser for provided pkg
type Parser struct {
	fset *token.FileSet

	structs []model.Struct

	fileParse *fileParser
}

func NewParser(opts ...Option) *Parser {
	opt := newOptions(opts...)
	p := &Parser{
		fset:      token.NewFileSet(),
		fileParse: newFileParser(opt.buildStructMatcher()),
	}

	return p
}

func (p *Parser) Reset() {
	p.structs = p.structs[:0]
	p.fset = token.NewFileSet()
}

func (p *Parser) ParsedStructs() []model.Struct {
	return p.structs
}

func (p *Parser) ParsePackage(path string) error {
	dirEntrys, err := os.ReadDir(path)
	if err != nil {
		return fmt.Errorf("failed to read directory: %w", err)
	}

	for _, entry := range dirEntrys {
		if entry.IsDir() {
			continue
		}

		astFile, err := p.parseFileByParentPathAndFileName(path, entry.Name())
		if err != nil {
			return fmt.Errorf("failed to parseFileByUrl path=%s, fileName=%s: %w", path, entry.Name(), err)
		}

		err = p.ParseAstFile(astFile)
		if err != nil {
			return fmt.Errorf("failed ParseAstFile: %w", err)
		}
	}

	return nil
}

func (p *Parser) parseFileByParentPathAndFileName(parentPath string, fileName string) (*ast.File, error) {
	fileUrl, err := url.JoinPath(parentPath, fileName)
	if err != nil {
		return nil, fmt.Errorf("failed to JoinPath: %w", err)
	}

	file, err := os.Open(fileUrl)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	parsedFile, err := parser.ParseFile(
		p.fset,
		fileName,
		file,
		parser.ParseComments|parser.SkipObjectResolution,
	)
	if err != nil {
		return nil, fmt.Errorf("failed parser.ParseFile: %w", err)
	}

	return parsedFile, nil
}

// ParseAstFile parse ast file directly
func (p *Parser) ParseAstFile(file *ast.File) error {
	defer p.fileParse.Reset()
	err := p.fileParse.parseAstFile(file)
	if err != nil {
		return fmt.Errorf("failed parseAstFile: %w", err)
	}

	p.structs = append(p.structs, p.fileParse.structs...)

	return nil
}
