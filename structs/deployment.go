package structs

type Location struct {
	Address string `yaml:"address"`
	Prefix string `yaml:"prefix"`
}

type Deployment struct {
	Source         Location            `yaml:"source"`
	Target         Location            `yaml:"target"`
	ConfigBoundary string            `yaml:"configBoundary"`
	ConfigMap      map[string]string `yaml:"configMap"`
}
