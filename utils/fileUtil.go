package utils

import (
	"os"
	"os/exec"
	"runtime"
)

func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func WriteFile(filename string, data []byte) error {
	return os.WriteFile(filename, data, 0644)
}

func GetCurrentPath() string {
	dir, err := os.Getwd()
	if err != nil {
		return ""
	}
	return dir
}

func OpenFolder(path string) error {
	osType := runtime.GOOS
	switch osType {
	case "windows":
		return exec.Command("explorer", path).Start()
	case "linux":
		return exec.Command("xdg-open", path).Start()
	case "darwin":
		return exec.Command("open", path).Start()
	default:
		return nil
	}
}
