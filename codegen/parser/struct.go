package parser

import (
	field_type "github.com/0B1t322/zero-validaton/codegen/parser/field-type"
	"github.com/0B1t322/zero-validaton/internal/set"
	"path"
)

type Package struct {
	Imports []Import
	Structs []Struct
}

// Struct ...
type Struct struct {
	Name   string
	Fields []Field
}

type Field struct {
	Name string
	Type field_type.FieldTyper
	Tags map[string]string
}

type Import struct {
	Path  string
	Alias string
}

// GetUsedPackageName return alias if present. If not return
// alias as package name in import.
func (i Import) GetUsedPackageName() string {
	if i.Alias != "" {
		return i.Alias
	}

	return path.Base(i.Path)
}

type Structs []Struct

func (s Structs) GetUsedImports() []Import {
	usedImportsSet := set.New[Import]()

	var visitorFunc = getImportsFromFieldTypeVisitorFunc(func(pkgName string, pkgPath string) {
		usedImportsSet.Add(Import{Alias: pkgName, Path: pkgPath})
	})

	for _, st := range s {
		for _, field := range st.Fields {
			field.Type.Accept(visitorFunc)
		}
	}

	return usedImportsSet.Values()
}

type getImportsFromFieldTypeVisitorFunc func(pkgName string, pkgPath string)

func (g getImportsFromFieldTypeVisitorFunc) VisitBasic(_ field_type.Basic) {}

func (g getImportsFromFieldTypeVisitorFunc) VisitCustom(custom field_type.Custom) {
	if custom.PkgName != "" || custom.PkgPath != "" {
		g(custom.PkgName, custom.PkgPath)
	}
}

func (g getImportsFromFieldTypeVisitorFunc) VisitGeneric(_ field_type.Generic) {

}

func (g getImportsFromFieldTypeVisitorFunc) VisitPtr(_ field_type.Ptr) {

}

func (g getImportsFromFieldTypeVisitorFunc) VisitSlice(_ field_type.Slice) {

}

// TypeAlias ...
type TypeAlias struct {
	Name string
	To   string
}
