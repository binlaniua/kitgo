package kfile

import (
	"os"
	"path/filepath"
	"os/exec"
)

func RuntimePath() string {
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	return path
}
