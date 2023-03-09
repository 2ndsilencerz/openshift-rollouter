package model

type DeploymentConfig struct {
	ApiVersion string         `yaml:"apiVersion"`
	Kind       string         `yaml:"kind"`
	Metadata   Metadata       `yaml:"metadata"`
	Spec       DeploymentSpec `yaml:"spec"`
}
