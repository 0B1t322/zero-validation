package main

import (
	"fmt"
	"github.com/0B1t322/zero-validaton/codegen/config"
)

func main() {
	cfg := config.GetDefaultConfig()

	fmt.Printf("%+v\n", cfg)

	_ = cfg

	//const path = "/Users/danademin/Projects/Golang/zero-validation/internal/_testdata"
	//
	//p := parser.NewParser(
	//	parser.WithStructMatcherBuilder(
	//		matcher.
	//			NewBuilder().
	//			AddFullMatches("ToDo"),
	//	),
	//)
	//err := p.ParsePackage(path)
	//if err != nil {
	//	panic(err)
	//}
	//
	//g := generator.NewGenerator()
	//
	//err = g.Generate(generator.GenerateCommand{
	//	Structs:           p.ParsedStructs(),
	//	PackagePath:       "",
	//	DestinationPath:   "/Users/danademin/Projects/Golang/zero-validation/internal/_testdata/generated",
	//	GeneratedFileName: "validation_gen.go",
	//})
	//if err != nil {
	//	panic(err)
	//}
}
