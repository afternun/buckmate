package util

import (
	"bytes"
	"os"
)

func ReplaceInFiles(path string, boundary string, configMap map[string]string) error {
	files, err := os.ReadDir(path)
	if err != nil {
		return err
	}
	for _, file := range files {
		if !file.IsDir() {
			err := ReplaceInFile(path+"/"+file.Name(), boundary, configMap)
			if err != nil {
				return err
			}
		} else {
			err := ReplaceInFiles(path+"/"+file.Name(), boundary, configMap)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func ReplaceInFile(path string, boundary string, configMap map[string]string) error {
	read, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	for key, value := range configMap {
		read = bytes.ReplaceAll(read, []byte(boundary+key+boundary), []byte(value))
	}
	return os.WriteFile(path, read, os.FileMode(int(0644)))
}
