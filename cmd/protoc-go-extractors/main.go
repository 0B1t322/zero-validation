package main

import (
	"fmt"
	"github.com/0B1t322/zero-validation/codegen"
	"github.com/0B1t322/zero-validation/codegen/config"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/pluginpb"
	"os"
)

func main() {
	cfg := config.GetDefaultConfig()
	fmt.Fprintf(os.Stderr, "cfg: %+v\n", cfg)
	protogen.Options{}.Run(func(gen *protogen.Plugin) error {
		gen.SupportedFeatures = uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL) | uint64(pluginpb.CodeGeneratorResponse_FEATURE_SUPPORTS_EDITIONS)
		gen.SupportedEditionsMinimum = descriptorpb.Edition_EDITION_PROTO2
		gen.SupportedEditionsMaximum = descriptorpb.Edition_EDITION_2023

		//if gen.Files[0].GoPackageName == "external" {
		//	return nil
		//}
		//
		//externalEnumField := gen.Files[4].Messages[0].Fields[5]
		//enumField := gen.Files[4].Messages[0].Fields[4]
		//fmt.Fprintf(os.Stderr, "%+v\n", externalEnumField.Enum.GoIdent.GoImportPath)
		//fmt.Fprintf(os.Stderr, "%+v\n", enumField.Enum.GoIdent.GoImportPath)
		//fmt.Fprintf(os.Stderr, "%+v\n", enumField.GoName)

		err := codegen.NewGRPC(gen, cfg).Generate()
		if err != nil {
			return err
		}

		return nil
	})
}
