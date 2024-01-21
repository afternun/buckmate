package util

import (
	"log"
	"os"
)

func RandomDirectory() string {
	dirName, err := os.MkdirTemp("", "")
	if err != nil {
		log.Fatalln("Could not create temporary directory.")
	}
	return dirName
}

func RemoveDirectory(path string) {
	if err := os.RemoveAll(path); err != nil {
		log.Fatalln("Could not remove tmp directory.")
	}
}
