package util

import (
	"os"
)

func RandomDirectory() (string, error) {
	dirName, err := os.MkdirTemp("", "")
	if err != nil {
		return "", err
	}
	return dirName, nil
}

func RemoveDirectory(path string) error {
	return os.RemoveAll(path)
}
