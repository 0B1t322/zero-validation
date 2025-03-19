package parser

import (
	"errors"
	"fmt"
	"github.com/0B1t322/zero-validaton/codegen/matcher"
	"github.com/0B1t322/zero-validaton/codegen/parser"
	field_type "github.com/0B1t322/zero-validaton/codegen/parser/field-type"
	"strings"

	"go/ast"

	"iter"
	"regexp"
)

type fileParser struct {
	// all imports in file
	imports *importSet

	structs []parser.Struct

	structMatcher matcher.StructMatcher
}

func newFileParser(structMatcher matcher.StructMatcher) *fileParser {
	return &fileParser{
		imports:       newImportSet(),
		structMatcher: structMatcher,
	}
}

func (f *fileParser) Reset() {
	f.structs = f.structs[:0]
	f.imports.Clean()
}

func (f *fileParser) parseAstFile(file *ast.File) error {
	f.addImports(file.Imports)
	structsTypeSpecs := f.getProvidedStructsToParse(file)
	for spec := range structsTypeSpecs {
		err := f.parseTypeSpec(spec)
		if err != nil {
			return err
		}
	}

	return nil
}

// getProvidedStructsToParse return iterator of [ast.TypeSpec] where in Type is [ast.StructType]
func (f *fileParser) getProvidedStructsToParse(file *ast.File) iter.Seq[*ast.TypeSpec] {
	return func(yield func(*ast.TypeSpec) bool) {
		//	TODO: add struct matcher
		for _, decl := range file.Decls {
			genDecl, ok := decl.(*ast.GenDecl)
			if !ok {
				continue
			}

			for _, spec := range genDecl.Specs {
				typeSpec, ok := spec.(*ast.TypeSpec)
				if !ok {
					continue
				}

				_, ok = typeSpec.Type.(*ast.StructType)
				if !ok {
					continue
				}

				if !f.structMatcher.Match(typeSpec.Name.Name) {
					continue
				}

				if !yield(typeSpec) {
					return
				}
			}
		}
	}
}

func (f *fileParser) parseTypeSpec(typeSpec *ast.TypeSpec) error {
	structType := typeSpec.Type.(*ast.StructType)
	parsedStruct := parser.Struct{
		Name:   typeSpec.Name.Name,
		Fields: make([]parser.Field, 0, len(structType.Fields.List)),
	}
	for _, field := range structType.Fields.List {
		parsedField, err := f.parseField(field)
		if err != nil {
			return err
		}
		parsedStruct.Fields = append(parsedStruct.Fields, parsedField)
	}

	f.structs = append(f.structs, parsedStruct)
	return nil
}

func (f *fileParser) addImports(imports []*ast.ImportSpec) {
	for _, spec := range imports {
		path := strings.ReplaceAll(spec.Path.Value, `"`, "")
		imp := parser.Import{
			Path: path,
		}

		if spec.Name != nil {
			imp.Alias = spec.Name.Name
		}

		f.imports.Add(imp)
	}
}

var tagRegexp = regexp.MustCompile(`(?P<name>.*):"(?P<value>.*)"`)

func (f *fileParser) parseField(field *ast.Field) (parser.Field, error) {
	parsedField := parser.Field{
		Name: field.Names[0].Name,
		Type: nil,
	}

	fieldType, err := f.parseFieldType(field.Type)
	if err != nil {
		return parsedField, fmt.Errorf("failed to parse field type: %w", err)
	}
	parsedField.Type = fieldType

	if field.Tag != nil {
		//	TODO: parse tags
	}

	return parsedField, nil
}

var (
	errUnknownFieldType = errors.New("unknown field type")
)

// parseFieldType return type, where is final [field_type.Kind] oneOf: [[field_type.KindBasic], [field_type.KindCustom]]
func (f *fileParser) parseFieldType(fieldTypeExpr ast.Expr) (field_type.FieldTyper, error) {
	traversal := func(fieldTypeExpr ast.Expr, fieldTyper field_type.FieldTyper) (field_type.FieldTyper, error) {
		if fieldTypeExpr == nil {
			return fieldTyper, nil
		}

		switch expr := fieldTypeExpr.(type) {
		case *ast.Ident:
			basic, ok := field_type.ParseBasic(expr.Name)
			if ok {
				return basic, nil
			}

			return field_type.CustomFiled(expr.Name, "", ""), nil
		case *ast.SelectorExpr:
			pkgName := expr.X.(*ast.Ident).Name
			pkgPath, err := f.getPkgPathByAlias(pkgName)
			if err != nil {
				return nil, err
			}

			return field_type.CustomFiled(expr.Sel.Name, pkgName, pkgPath), nil
		case *ast.IndexExpr:
			//	generic
			typeInGeneric, err := f.parseFieldType(expr.Index)
			if err != nil {
				return nil, err
			}

			genericType, err := f.parseFieldType(expr.X)
			if err != nil {
				return nil, err
			}

			return field_type.GenericField(genericType, typeInGeneric), nil
		case *ast.StarExpr:
			innerFieldType, err := f.parseFieldType(expr.X)
			if err != nil {
				return nil, err
			}

			return field_type.PtrField(innerFieldType), nil
		case *ast.ArrayType:
			innerFieldType, err := f.parseFieldType(expr.Elt)
			if err != nil {
				return nil, err
			}
			return field_type.SliceField(innerFieldType), nil
		}

		return nil, errUnknownFieldType
	}
	return traversal(fieldTypeExpr, nil)
}

func (f *fileParser) getPkgPathByAlias(alias string) (string, error) {
	findedImport, isFind := f.imports.importByAlias[alias]
	if !isFind {
		return "", fmt.Errorf("not found import by alias %s", alias)
	}

	return findedImport.Path, nil
}
