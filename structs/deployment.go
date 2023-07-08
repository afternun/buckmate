package structs

type defaultTags struct {
	Version bool `yaml:"version"`
}

type source struct {
	Path string `yaml:"path"`
}

type target struct {
	Path        string            `yaml:"path"`
	Tags        map[string]string `yaml:"tags"`
	DefaultTags defaultTags       `yaml:"defaultTags"`
}

type Deployment struct {
	Source source `yaml:"source"`
	Target target `yaml:"target"`
}
