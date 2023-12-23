package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

const version = "0.0.2"

func run() int {
	var (
		ver     string
		showVer bool
	)
	flag.StringVar(&ver, "v", "", "update to this version of Go")
	flag.BoolVar(&showVer, "version", false, "show version")
	flag.Parse()

	if showVer {
		fmt.Printf("go-mod-version %s\n", version)
		return 0
	}

	if ver == "" {
		fmt.Println("Please specify the version of Go with the '-version' flag. Example: -v 1.21")
		return 1
	}

	fmt.Printf("updating version: %s\n", ver)

	err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() || filepath.Base(path) != "go.mod" {
			return nil
		}
		updateGoVersion(path, ver)
		return nil
	})

	if err != nil {
		fmt.Printf("%v\n", err)
		return 1
	}
	return 0
}

func updateGoVersion(modFilePath, version string) error {
	fmt.Println(modFilePath)
	cmd := exec.Command("go", "mod", "tidy", fmt.Sprintf("-go=%s", version))
	cmd.Dir = filepath.Dir(modFilePath)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("コマンド実行エラー: %v\n出力: %s\n", err, string(output))
		return err
	}
	return nil
}

func main() {
	os.Exit(run())
}
