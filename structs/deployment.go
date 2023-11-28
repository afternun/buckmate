package structs

type bucket struct {
	Bucket string `yaml:"bucket"`
	Prefix string `yaml:"prefix"`
}

type Deployment struct {
	Source         bucket            `yaml:"source"`
	Target         bucket            `yaml:"target"`
	ConfigBoundary string            `yaml:"configBoundary"`
	ConfigMap      map[string]string `yaml:"configMap"`
}
