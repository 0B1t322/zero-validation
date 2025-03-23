package config

type PackagesConfig map[string]PackageConfig

type PackageConfig struct {
	Structs PackageStructsConfig `yaml:"structs"`
	Dst     string               `yaml:"dst"`
}

type PackageStructsConfig struct {
	Include []string `yaml:"include"`
	Exclude []string `yaml:"exclude"`
}
