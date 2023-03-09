package model

type Secret struct {
	ApiVersion string                 `yaml:"apiVersion"`
	Kind       string                 `yaml:"kind"`
	Metadata   Metadata               `yaml:"metadata"`
	Data       map[string]interface{} `yaml:"data"`
	Type       string                 `yaml:"type"`
}
