package utils

import (
	"path/filepath"
	"testing"
)

func TestGetConfigPath(t *testing.T) {
	path := GetConfigPath()
	t.Log(path)
}

func TestWriteConfigFile(t *testing.T) {
	path := GetConfigPath()
	WriteFile(filepath.Join(path, "test.txt"), []byte("test"))
}
