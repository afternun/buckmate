package util

import (
	"bytes"
	"log"
	"os"

	"dario.cat/mergo"
	"gopkg.in/yaml.v2"
)

func MergeStruct(dest interface{}, src interface{}) error {
	return mergo.Merge(dest, src, mergo.WithOverride)
}

func LoadYaml(path string) ([]byte, error) {
	b, err := os.ReadFile(path)
	return b, err
}

func YamlToStruct(data []byte, dest interface{}) error {
	return yaml.UnmarshalStrict([]byte(data), dest)
}

func ReplaceInFiles(path string, boundary string, configMap map[string]string) {
	files, err := os.ReadDir(path)
	if err != nil {
		log.Fatalln("Couldn't list files under " + path)
	}
	for _, file := range files {
		if !file.IsDir() {
			err := ReplaceInFile(path+"/"+file.Name(), boundary, configMap)
			if err != nil {
				log.Fatalln("Couldn't apply config map to " + path)
			}
		} else {
			ReplaceInFiles(path+"/"+file.Name(), boundary, configMap)
		}
	}
}

func ReplaceInFile(path string, boundary string, configMap map[string]string) error {
	read, err := os.ReadFile(path)
	if err != nil {
		log.Fatalln("Could not load file " + path)
	}
	for key, value := range configMap {
		read = bytes.ReplaceAll(read, []byte(boundary+key+boundary), []byte(value))
	}
	return os.WriteFile(path, read, os.FileMode(int(0644)))
}
