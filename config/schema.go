package config

type Project struct {
	Version  string             `yaml:"version"`
	Name     string             `yaml:"name"`
	Services map[string]Service `yaml:"services"`
}

type Service struct {
	Image string `yaml:"image"`
	Build Build  `yaml:"build"`
}

type Build struct {
	Context     string `yaml:"context"`
	Destination string `yaml:"destination"`
}
