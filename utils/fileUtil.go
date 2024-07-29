package utils

import (
	"os"
	"os/exec"
	"path/filepath"
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

func ReadFile(filename string) ([]byte, error) {
	return os.ReadFile(filename)
}

func InitConfigPath() {
	configPath := GetConfigPath()
	os.MkdirAll(configPath, 0755)
}

func GetConfigPath() string {
	osType := runtime.GOOS
	configPath := ""
	switch osType {
	case "windows":
		homeDir, _ := os.UserHomeDir()
		configPath = filepath.Join(homeDir, "AppData", "Roaming", "AiGui", "config")
	case "linux":
		homeDir, _ := os.UserHomeDir()
		configPath = filepath.Join(homeDir, ".config", "AiGui", "config")
	case "darwin":
		homeDir, _ := os.UserHomeDir()
		configPath = filepath.Join(homeDir, "Library", "Application Support", "AiGui", "config")
	default:
		dir, _ := os.Getwd()
		configPath = filepath.Join(dir, "AiGui", "config")
	}
	return configPath
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
