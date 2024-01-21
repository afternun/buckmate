package util

import (
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
