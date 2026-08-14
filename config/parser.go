package config

import (
	"fmt"

	"gopkg.in/yaml.v3"
)

func Parse(fileData []byte) (*Project, error) {

	var pmanService Project
	// unmarshal it into the struct
	err := yaml.Unmarshal(fileData, &pmanService)

	if err != nil {
		return nil, err
	}

	// Now we can go straight to the podman API
	for name, service := range pmanService.Services {
		fmt.Println("name", name)
		fmt.Println("service", service)
		fmt.Println("context", service.Build.Context)
	}

	return &pmanService, nil
}
