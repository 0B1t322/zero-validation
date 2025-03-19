package generator

import (
	"os"
	"runtime/debug"
)

func basePathFromOS() (string, error) {
	basePath, err := os.Getwd()
	if err != nil {
		return "", err
	}

	return basePath, nil
}

func getModuleName() (string, bool) {
	info, ok := debug.ReadBuildInfo()

	return info.Main.Path, ok
}
