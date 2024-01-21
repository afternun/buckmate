package structs

type Deployment struct {
	Source struct {
		Path string
	}
	Target struct {
		Path string
		Tags map[string]string
		DefaultTags bool
	}
	Artifact struct {
		Name string
	}
}