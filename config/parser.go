package config

import (
	"gopkg.in/yaml.v3"
)

func Parse(fileData []byte) (*Project, error) {

	var pmanService Project
	// unmarshal it into the struct
	err := yaml.Unmarshal(fileData, &pmanService)

	if err != nil {
		return nil, err
	}

	return &pmanService, nil
}
