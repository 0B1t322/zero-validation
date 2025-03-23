package parser

import (
	"errors"
	"fmt"
	"github.com/0B1t322/zero-validaton/codegen/matcher"
	"github.com/0B1t322/zero-validaton/codegen/parser"
	field_type "github.com/0B1t322/zero-validaton/codegen/parser/field-type"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
	"iter"
	"strings"
)

var (
	errSkip = errors.New("skip")
)

type Parser struct {
	structs []parser.Struct

	gen           *protogen.Plugin
	structMatcher matcher.StructMatcher

	parsedFilePackagePath string
}

func NewParser(opts ...Option) *Parser {
	opt := newOptions(opts...)

	p := &Parser{
		structMatcher: opt.structMatcher(),
	}

	return p
}

func (p *Parser) Structs() []parser.Struct {
	return p.structs
}

func (p *Parser) Reset() {
	p.structs = p.structs[:0]
	p.parsedFilePackagePath = ""
}

func (p *Parser) ParseFiles(files []*protogen.File) error {
	for _, f := range files {
		if err := p.ParseFile(f); err != nil {
			return err
		}
	}

	return nil
}

func (p *Parser) ParseFile(file *protogen.File) error {
	p.parsedFilePackagePath = file.GoDescriptorIdent.GoImportPath.String()
	messagesSeq := p.getProvidedMessagesToParse(file)

	for message := range messagesSeq {
		if err := p.parseMessage(message); err != nil {
			return err
		}
	}

	return nil
}

func (p *Parser) getProvidedMessagesToParse(file *protogen.File) iter.Seq[*protogen.Message] {
	//	TODO: add struct matcher
	return func(yield func(*protogen.Message) bool) {
		for _, m := range file.Messages {
			if !p.recursiveMessageTravesal(m, yield) {
				return
			}
		}
	}
}

func (p *Parser) recursiveMessageTravesal(m *protogen.Message, yield func(*protogen.Message) bool) bool {
	for _, m := range m.Messages {
		if !p.recursiveMessageTravesal(m, yield) {
			return false
		}
	}

	if !p.structMatcher.Match(m.GoIdent.GoName) {
		return true
	}

	if !yield(m) {
		return false
	}

	return true
}

func (p *Parser) parseMessage(message *protogen.Message) error {
	parsedStruct := parser.Struct{
		Name: message.GoIdent.GoName,
	}

	// oneOf special case
	parsedOneofFields := make(map[*protogen.Oneof]struct{})
	for _, field := range message.Fields {
		var (
			parsedField parser.Field
			err         error
		)
		switch {
		case field.Oneof != nil && !field.Oneof.Desc.IsSynthetic():
			if _, isFind := parsedOneofFields[field.Oneof]; isFind {
				continue
			}
			parsedField, err = p.parseOneOfField(field)
			parsedOneofFields[field.Oneof] = struct{}{}
		default:
			parsedField, err = p.parseField(field)
		}

		if err != nil {
			switch {
			case errors.Is(err, errSkip):
				continue
			}
			return err
		}
		parsedStruct.Fields = append(parsedStruct.Fields, parsedField)
	}

	p.structs = append(p.structs, parsedStruct)
	if err := p.parseMessageOneOfs(message); err != nil {
		return err
	}

	return nil
}

func (p *Parser) parseMessageOneOfs(message *protogen.Message) error {
	if len(message.Oneofs) == 0 {
		return nil
	}

	for _, oneof := range message.Oneofs {
		if oneof.Desc.IsSynthetic() {
			continue
		}

		oneOfStructs, err := p.messageOneOfToStructs(message, oneof)
		if err != nil {
			return err
		}

		p.structs = append(p.structs, oneOfStructs...)
	}

	return nil
}

func (p *Parser) messageOneOfToStructs(_ *protogen.Message, oneOf *protogen.Oneof) ([]parser.Struct, error) {
	oneOfStructs := make([]parser.Struct, 0, len(oneOf.Fields))
	for _, field := range oneOf.Fields {
		oneOfField, err := p.parseField(field)
		if err != nil {
			return nil, err
		}

		oneOfStruct := parser.Struct{
			Name:   field.GoIdent.GoName,
			Fields: []parser.Field{oneOfField},
		}

		oneOfStructs = append(oneOfStructs, oneOfStruct)
	}

	return oneOfStructs, nil
}

func (p *Parser) parseOneOfField(field *protogen.Field) (parser.Field, error) {
	if field.Oneof.Desc.IsSynthetic() {
		return parser.Field{}, fmt.Errorf("%w: one of for optional field", errSkip)
	}

	return parser.Field{
		Name: field.Oneof.GoName,
		Type: field_type.CustomFiled(
			"is"+field.Oneof.GoIdent.GoName,
			"",
			"",
		),
	}, nil
}

func (p *Parser) parseField(field *protogen.Field) (parser.Field, error) {
	parsedField := parser.Field{
		Name: field.GoName,
		Type: nil,
		Tags: make(map[string]string),
	}

	fieldType, err := p.parseFieldType(field)
	if err != nil {
		return parsedField, err
	}
	parsedField.Type = fieldType

	if field.Desc.HasJSONName() {
		parsedField.Tags["json"] = field.Desc.JSONName()
	}

	parsedField.Tags["proto"] = string(field.Desc.Name())

	return parsedField, nil
}

func (p *Parser) parseFieldType(field *protogen.Field) (field_type.FieldTyper, error) {
	basicType, isBasicType := tryGetBasicTypeFromKind(field.Desc)
	if isBasicType {
		return basicType, nil
	}

	return p.parseCustomType(field)
}

func (p *Parser) parseCustomType(field *protogen.Field) (field_type.FieldTyper, error) {
	customTypeImportPath := getTypeImportPath(field)

	isTypeOfThisFile := customTypeImportPath == p.parsedFilePackagePath

	var pkgPath string
	if !isTypeOfThisFile {
		pkgPath = strings.ReplaceAll(customTypeImportPath, `"`, "")
	}

	var customFieldType field_type.FieldTyper
	customFieldType = field_type.CustomFiled(
		getTypeName(field),
		"",
		pkgPath,
	)

	// in grpc wrap ptr in two cases:
	// 1. Field is message
	// 2. Field is optional. For optional we have one_of, that start from _
	if field.Desc.HasOptionalKeyword() || isMessage(field) {
		customFieldType = field_type.PtrField(customFieldType)
	}

	switch field.Desc.Cardinality() {
	case protoreflect.Repeated:
		customFieldType = field_type.SliceField(customFieldType)
	}

	return customFieldType, nil
}

func getTypeImportPath(field *protogen.Field) string {
	if field.Enum != nil {
		return field.Enum.GoIdent.GoImportPath.String()
	}

	if field.Message != nil {
		return field.Message.GoIdent.GoImportPath.String()
	}

	panic("type is not custom can't get go import path")
}

func isMessage(filed *protogen.Field) bool {
	return filed.Message != nil
}

func getTypeName(field *protogen.Field) string {
	if field.Enum != nil {
		return field.Enum.GoIdent.GoName
	}

	if field.Message != nil {
		return field.Message.GoIdent.GoName
	}

	panic("type is not custom can't get name")
}

var fieldDescriptorKindToFieldType = map[protoreflect.Kind]field_type.FieldTyper{
	protoreflect.BoolKind:     field_type.BasicBool,
	protoreflect.Int32Kind:    field_type.BasicInt32,
	protoreflect.Int64Kind:    field_type.BasicInt64,
	protoreflect.Uint32Kind:   field_type.BasicUint32,
	protoreflect.Uint64Kind:   field_type.BasicUint64,
	protoreflect.StringKind:   field_type.BasicString,
	protoreflect.FloatKind:    field_type.BasicFloat32,
	protoreflect.DoubleKind:   field_type.BasicFloat64,
	protoreflect.Fixed32Kind:  field_type.BasicUint32,
	protoreflect.Fixed64Kind:  field_type.BasicUint64,
	protoreflect.Sfixed32Kind: field_type.BasicInt32,
	protoreflect.Sfixed64Kind: field_type.BasicInt64,
	// special case — because it's array of bytes
	protoreflect.BytesKind: field_type.SliceField(field_type.BasicByte),
}

func tryGetBasicTypeFromKind(desc protoreflect.FieldDescriptor) (field_type.FieldTyper, bool) {
	basicFieldType, isFind := fieldDescriptorKindToFieldType[desc.Kind()]
	if !isFind {
		return nil, false
	}

	// is field optional it have one_of
	if desc.HasOptionalKeyword() {
		basicFieldType = field_type.PtrField(basicFieldType)
	}

	switch desc.Cardinality() {
	case protoreflect.Repeated:
		basicFieldType = field_type.SliceField(basicFieldType)
	}

	return basicFieldType, true
}
