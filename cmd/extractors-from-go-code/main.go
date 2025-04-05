package main

import (
	"fmt"
	"github.com/0B1t322/zero-validation/codegen"
	"github.com/0B1t322/zero-validation/codegen/config"
	"os"
)

func main() {
	cfg := config.GetDefaultConfig()
	codeGen := codegen.NewPackages(cfg)
	if err := codeGen.Generate(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

}
