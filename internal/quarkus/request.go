package quarkus

type ProjectInitializr struct {
	Group       string
	Artifact    string
	BuildTool   string
	Version     string
	JavaVersion string
	StarterCode bool
	Extension   []string
}
