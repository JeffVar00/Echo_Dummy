package settings

import (
	_ "embed"

	"gopkg.in/yaml.v3" //go get gopkg.in/yaml.v3 paquete yaml
)

//go:embed settings.yaml
var settingsFile []byte

type Settings struct {
	Port string `yaml:"port"`
}

func New() (*Settings, error) { // Returns tje secundary variables
	var s Settings
	err := yaml.Unmarshal(settingsFile, &s)

	if err != nil {
		return nil, err
	}

	return &s, nil
}
