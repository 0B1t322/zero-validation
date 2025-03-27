package gqlgen

import (
	"errors"
	"fmt"
	"github.com/0B1t322/zero-validaton/codegen/matcher"
	"github.com/0B1t322/zero-validaton/codegen/parser"
	field_type "github.com/0B1t322/zero-validaton/codegen/parser/field-type"
	"github.com/99designs/gqlgen/codegen"
	"go/types"
)

type tagParser interface {
	ParseTag(tag string) map[string]string
}

type Parser struct {
	structs []parser.Struct

	structMatcher         matcher.StructMatcher
	parsedFilePackagePath string

	tagParser tagParser
}

func NewParser(tagParser tagParser) *Parser {
	return &Parser{
		tagParser: tagParser,
	}
}

func (p *Parser) Reset() {
	p.structs = p.structs[:0]
	p.parsedFilePackagePath = ""
}

func (p *Parser) Structs() []parser.Struct {
	return p.structs
}

func (p *Parser) ParseObject(pkgPath string, object *codegen.Object) error {
	// Обычно здесь named object
	namedType, ok := object.Type.(*types.Named)
	if !ok {
		return fmt.Errorf("object.Type is not types.Named, it's %T type", object.Type)
	}
	p.parsedFilePackagePath = pkgPath

	err := p.parseNamedType(namedType)
	if err != nil {
		return err
	}

	return nil
}

func (p *Parser) parseNamedType(namedType *types.Named) error {
	unerlaying := namedType.Underlying()

	structType, ok := unerlaying.(*types.Struct)
	if !ok {
		return fmt.Errorf("unerlaying type %T is not a types.Struct", unerlaying)
	}

	parsedStruct := parser.Struct{
		Name:   namedType.Obj().Name(),
		Fields: nil,
	}

	var err error
	parsedStruct.Fields, err = p.parseStructFields(structType)
	if err != nil {
		return err
	}

	p.structs = append(p.structs, parsedStruct)

	return nil
}

func (p *Parser) parseStructFields(structType *types.Struct) ([]parser.Field, error) {
	parsedFields := make([]parser.Field, 0, structType.NumFields())

	for i := range structType.NumFields() {
		field := structType.Field(i)
		parsedField := parser.Field{
			Name: field.Name(),
			Type: nil,
			Tags: nil,
		}

		var err error

		parsedField.Type, err = p.parseFieldType(field.Type())
		if err != nil {
			return nil, err
		}

		parsedField.Tags = p.parseFieldTags(structType.Tag(i))

		parsedFields = append(parsedFields, parsedField)
	}

	return parsedFields, nil
}

var (
	errUnknownFieldType = errors.New("unknown field type")
)

func (p *Parser) parseFieldType(fieldType types.Type) (field_type.FieldTyper, error) {
	travesal := func(fieldType types.Type, fieldTyper field_type.FieldTyper) (field_type.FieldTyper, error) {
		if fieldType == nil {
			return fieldTyper, nil
		}

		switch currentType := fieldType.(type) {
		case *types.Pointer:
			innerFieldType, err := p.parseFieldType(currentType.Elem())
			if err != nil {
				return nil, err
			}

			return field_type.PtrField(innerFieldType), nil
		case *types.Basic:
			basicType, ok := basicFieldTyperFromBasicKind[currentType.Kind()]
			if !ok {
				return nil, fmt.Errorf("%w: %d", errUnknownFieldType, currentType.Kind())
			}

			return basicType, nil
		case *types.Slice:
			innerFieldType, err := p.parseFieldType(currentType.Elem())
			if err != nil {
				return nil, err
			}

			return field_type.SliceField(innerFieldType), nil
		case *types.Named:
			if p.parsedFilePackagePath == currentType.Obj().Pkg().Path() {
				return field_type.CustomField(currentType.Obj().Name(), "", ""), nil
			}

			return field_type.CustomField(
				currentType.Obj().Name(),
				"",
				currentType.Obj().Pkg().Path(),
			), nil
		}

		return nil, errUnknownFieldType
	}

	return travesal(fieldType, nil)
}

var basicFieldTyperFromBasicKind = map[types.BasicKind]field_type.FieldTyper{
	types.String:  field_type.BasicString,
	types.Int:     field_type.BasicInt,
	types.Int32:   field_type.BasicInt32,
	types.Int64:   field_type.BasicInt64,
	types.Uint:    field_type.BasicUint,
	types.Uint32:  field_type.BasicUint32,
	types.Uint64:  field_type.BasicUint64,
	types.Bool:    field_type.BasicBool,
	types.Float32: field_type.BasicFloat32,
	types.Float64: field_type.BasicFloat64,
	types.Byte:    field_type.BasicByte,
}

func (p *Parser) parseFieldTags(tag string) map[string]string {
	return p.tagParser.ParseTag(tag)
}
