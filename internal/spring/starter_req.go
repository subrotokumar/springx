package spring

type ProjectInitializr struct {
	Project           string
	Language          string
	SpringBootVersion string
	ProjectMetadata   ProjectMetadata
	Dependencies      []string
}

type ProjectMetadata struct {
	GroupID       string
	ArtifactID    string
	Name          string
	Description   string
	PackageName   string
	Packaging     string
	Configuration string
	JavaVersion   string
}
