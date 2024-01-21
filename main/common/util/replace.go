package util

import (
	"bytes"
	"os"

	log "github.com/sirupsen/logrus"
)

func ReplaceInFiles(path string, boundary string, configMap map[string]string) {
	files, err := os.ReadDir(path)
	if err != nil {
		log.Fatal("Couldn't list files under " + path)
	}
	for _, file := range files {
		if !file.IsDir() {
			err := ReplaceInFile(path+"/"+file.Name(), boundary, configMap)
			if err != nil {
				log.Fatal("Couldn't apply config map to " + path)
			}
		} else {
			ReplaceInFiles(path+"/"+file.Name(), boundary, configMap)
		}
	}
}

func ReplaceInFile(path string, boundary string, configMap map[string]string) error {
	read, err := os.ReadFile(path)
	if err != nil {
		log.Fatal("Could not load file " + path)
	}
	for key, value := range configMap {
		read = bytes.ReplaceAll(read, []byte(boundary+key+boundary), []byte(value))
	}
	return os.WriteFile(path, read, os.FileMode(int(0644)))
}
