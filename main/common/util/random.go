package util

import (
	"os"

	log "github.com/sirupsen/logrus"
)

func RandomDirectory() string {
	dirName, err := os.MkdirTemp("", "")
	if err != nil {
		log.Fatal("Could not create temporary directory.")
	}
	return dirName
}

func RemoveDirectory(path string) {
	if err := os.RemoveAll(path); err != nil {
		log.Fatal("Could not remove tmp directory.")
	}
}
