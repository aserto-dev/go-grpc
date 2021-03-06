//go:build mage
// +build mage

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/aserto-dev/mage-loot/buf"
	"github.com/aserto-dev/mage-loot/deps"
	"github.com/aserto-dev/mage-loot/fsutil"
	"github.com/aserto-dev/mage-loot/mage"
)

func All() error {
	Deps()
	err := Clean()
	if err != nil {
		return err
	}
	return Generate()
}

// install required dependencies.
func Deps() {
	deps.GetAllDeps()
}

// Generate go code
func Generate() error {
	bufImage := "buf.build/aserto-dev/aserto"

	tag, err := buf.GetLatestTag(bufImage)
	if err != nil {
		fmt.Println("Could not retrieve tags, using latest")
	} else {
		bufImage = fmt.Sprintf("%s:%s", bufImage, tag.Name)
	}

	return gen(bufImage, bufImage)
}

// Generates from a dev build.
func GenerateDev() error {
	err := BuildDev()
	if err != nil {
		return err
	}

	bufImage := filepath.Join(getProtoRepo(), "bin", "aserto.bin#format=bin")
	fileSources := filepath.Join(getProtoRepo(), "public#format=dir")

	return gen(bufImage, fileSources)
}

// Builds the aserto proto image
func BuildDev() error {
	return mage.RunDir(getProtoRepo(), mage.AddArg("build"))
}

func getProtoRepo() string {
	protoRepo := os.Getenv("PROTO_REPO")
	if protoRepo == "" {
		protoRepo = "../proto"
	}
	return protoRepo
}

func gen(bufImage, fileSources string) error {

	files, err := getClientFiles(fileSources)
	if err != nil {
		return err
	}

	oldPath := os.Getenv("PATH")

	pathSeparator := string(os.PathListSeparator)
	path := oldPath +
		pathSeparator +
		filepath.Dir(deps.GoBinPath("protoc-gen-go-grpc")) +
		pathSeparator +
		filepath.Dir(deps.GoBinPath("protoc-gen-grpc-gateway")) +
		pathSeparator +
		filepath.Dir(deps.GoBinPath("protoc-gen-go"))

	return buf.RunWithEnv(map[string]string{
		"PATH": path,
	},
		buf.AddArg("generate"),
		buf.AddArg("--template"),
		buf.AddArg(filepath.Join("buf", "buf.gen.yaml")),
		buf.AddArg(bufImage),
		buf.AddPaths(files),
	)
}

func getClientFiles(fileSources string) ([]string, error) {
	var clientFiles []string

	bufExportDir, err := ioutil.TempDir("", "bufimage")
	if err != nil {
		return clientFiles, err
	}
	bufExportDir = filepath.Join(bufExportDir, "")

	defer os.RemoveAll(bufExportDir)
	err = buf.Run(
		buf.AddArg("export"),
		buf.AddArg(fileSources),
		buf.AddArg("--exclude-imports"),
		buf.AddArg("-o"),
		buf.AddArg(bufExportDir),
	)
	if err != nil {
		return clientFiles, err
	}
	excludePattern := filepath.Join(bufExportDir, "aserto", "authorizer", "authorizer", "**", "*.proto")

	protoFiles, err := fsutil.Glob(filepath.Join(bufExportDir, "aserto", "**", "*.proto"), excludePattern)
	if err != nil {
		return clientFiles, err
	}

	for _, protoFile := range protoFiles {
		clientFiles = append(clientFiles, strings.TrimPrefix(protoFile, bufExportDir+string(filepath.Separator)))
	}

	return clientFiles, nil
}

// Removes generated files
func Clean() error {
	return os.RemoveAll("aserto")
}
