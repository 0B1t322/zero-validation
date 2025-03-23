package config

import (
	"fmt"
	"golang.org/x/mod/modfile"
	"gopkg.in/yaml.v3"
	"net/url"
	"os"
	"path/filepath"
)

const (
	defaultConfigName = ".zerovalid.yaml"
)

type Config struct {
	// path where file go.mod is located
	goModulePath string `yaml:"-"`
	// bastPath base path of project
	bastPath string `yaml:"-"`

	GrpcConfig     GrpcConfig       `yaml:"grpc"`
	AdditionalTags []AdditionalTags `yaml:"additional_tags"`
	PackagesConfig PackagesConfig   `yaml:"packages"`
}

func (c *Config) GoModulePath() string {
	return c.goModulePath
}

func (c *Config) BastPath() string {
	return c.bastPath
}

func GetDefaultConfig() Config {
	cfg, ok := tryReadDefaultConfig()
	if !ok {
		cfg.init()
	}

	return cfg
}

func tryReadDefaultConfig() (Config, bool) {
	path, err := os.Getwd()
	if err != nil {
		return Config{}, false
	}

	pathToConfig, err := url.JoinPath(path, defaultConfigName)
	if err != nil {
		return Config{}, false
	}

	cfg, err := ReadConfig(pathToConfig)
	if err != nil {
		return Config{}, false
	}

	return cfg, true
}

func ReadConfig(path string) (Config, error) {
	cfg := Config{}

	file, err := os.Open(path)
	if err != nil {
		return cfg, err
	}

	if err = yaml.NewDecoder(file).Decode(&cfg); err != nil {
		return cfg, err
	}

	cfg.init()

	return cfg, nil
}

func (c *Config) init() {
	c.setGoModulePath()
}

func (c *Config) setGoModulePath() {
	path, err := os.Getwd()
	if err != nil {
		panic(fmt.Errorf("failed to get current working directory: %w", err))
	}

	path, err = filepath.Abs(path)
	if err != nil {
		panic(err)
	}

	path = filepath.ToSlash(path)

	for {
		goModData, err := os.ReadFile(filepath.Join(path, "go.mod"))
		if os.IsNotExist(err) {
			filepath.Base(path)
			path, _ = filepath.Split(path)
			continue
		}
		if err != nil {
			panic(fmt.Errorf("failed to read go.mod: %w", err))
		}

		if path == "" {
			panic(fmt.Errorf("failed to find go.mod"))
		}

		moduleName := modfile.ModulePath(goModData)
		c.bastPath = path
		c.goModulePath = moduleName
		break
	}

}
