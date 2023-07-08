package structs

type Config struct {
	Version   string            `yaml:"version"`
	ConfigMap map[string]string `yaml:"configMap"`
}
