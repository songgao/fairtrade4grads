package main

import (
	"os/exec"
	"path"
	"strings"
)

const pkgImportPath = "github.com/songgao/sign_petition"

func getRootPath() (string, error) {
	out, err := exec.Command("go", "list", "-f", "{{.Dir}}", pkgImportPath).Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(out)), nil
}

func getAssetsPath() (string, error) {
	root, err := getRootPath()
	if err != nil {
		return "", err
	}
	return path.Join(root, "assets"), nil
}
