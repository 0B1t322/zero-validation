//go:build gqlgen_custom

package main

import (
	"fmt"
	"github.com/0B1t322/zero-validaton/codegen"
	zconfig "github.com/0B1t322/zero-validaton/codegen/config"
	"github.com/99designs/gqlgen/api"
	"github.com/99designs/gqlgen/codegen/config"
	"os"
)

func main() {
	cfg, err := config.LoadConfigFromDefaultLocations()
	if err != nil {
		fmt.Fprintln(os.Stderr, "failed to load config", err.Error())
		os.Exit(2)
	}

	zeroValidCfg := zconfig.GetDefaultConfig()

	gqlGenPlugin := codegen.NewGqlGen(zeroValidCfg, cfg)

	err = api.Generate(cfg,
		api.AddPlugin(gqlGenPlugin),
	)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(3)
	}
}
