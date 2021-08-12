// +build mage

package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/aserto-dev/mage-loot/buf"
	"github.com/aserto-dev/mage-loot/deps"
	"github.com/aserto-dev/mage-loot/fsutil"
)

var bufImage = "buf.build/mitza/aserto"

// install required dependencies.
func Deps() {
	deps.GetAllDeps()
}

// Generate the code
func Generate() error {
	return generateInternal()
}

// Generate Internal services
func generateInternal() error {

	files, err := getClientFiles()
	if err != nil {
		return err
	}

	oldPath := os.Getenv("PATH")
	currnetDirectory, err := os.Getwd()
	if err != nil {
		return err
	}

	// buf requires protoc in the path
	protocPath := filepath.Join(currnetDirectory, deps.BinPath("protoc"))
	err = os.Setenv("PATH", oldPath+string(os.PathListSeparator)+filepath.Dir(protocPath))
	if err != nil {
		return err
	}

	return buf.Run(
		buf.AddArg("generate"),
		buf.AddArg("--template"),
		buf.AddArg(filepath.Join("buf", "buf.gen.yaml")),
		buf.AddArg(bufImage),
		buf.AddPaths(files),
	)
}

func getClientFiles() ([]string, error) {
	var clientFiles []string

	bufExportDir, err := ioutil.TempDir("", "bufimage")
	if err != nil {
		return clientFiles, err
	}
	bufExportDir = filepath.Join(bufExportDir, "")

	defer os.RemoveAll(bufExportDir)
	err = buf.Run(
		buf.AddArg("export"),
		buf.AddArg(bufImage),
		buf.AddArg("-o"),
		buf.AddArg(bufExportDir),
	)
	if err != nil {
		return clientFiles, err
	}
	excludePattern := filepath.Join(bufExportDir, "aserto", "**", "*_internal.proto")

	protoFiles, err := fsutil.Glob(filepath.Join(bufExportDir, "aserto", "**", "*.proto"), excludePattern)
	if err != nil {
		return clientFiles, err
	}

	authorizerFiles, err := fsutil.Glob(filepath.Join(bufExportDir, "aserto", "authorizer", "authorizer", "**", "*.proto"), excludePattern)
	if err != nil {
		return clientFiles, err
	}

	mapAuthorizerFiles := make(map[string]int, len(authorizerFiles))
	for _, authorizerFile := range authorizerFiles {
		mapAuthorizerFiles[authorizerFile] = 1
	}
	for _, protoFile := range protoFiles {
		if _, found := mapAuthorizerFiles[protoFile]; !found {
			clientFiles = append(clientFiles, strings.TrimPrefix(protoFile, bufExportDir+string(filepath.Separator)))
		}
	}

	return clientFiles, nil
}

// Probably not needed
func Build() error {
	return nil
}
