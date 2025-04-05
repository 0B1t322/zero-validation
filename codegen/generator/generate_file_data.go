package generator

import "github.com/0B1t322/zero-validaton/codegen/parser"

type generateFileData struct {
	PkgName                   string
	Imports                   []parser.Import
	Structs                   []parser.Struct
	IsGenerateInParsedPackage bool
	TypeAlias                 []parser.TypeAlias
}
