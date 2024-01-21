package util

import (
	"os"

	"gopkg.in/yaml.v2"
)

func LoadYaml(path string) ([]byte, error) {
	b, err := os.ReadFile(path)
	return b, err
}

func YamlToStruct(data []byte, dest interface{}) error {
	return yaml.UnmarshalStrict([]byte(data), dest)
}
