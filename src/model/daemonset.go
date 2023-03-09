package model

type DaemonSet struct {
	ApiVersion string        `yaml:"apiVersion"`
	Kind       string        `yaml:"kind"`
	Metadata   Metadata      `yaml:"metadata"`
	Spec       DaemonSetSpec `yaml:"spec"`
}

type DaemonSetSpec struct {
	Selector       SpecSelector           `yaml:"selector"`
	Template       DeploymentSpecTemplate `yaml:"template"`
	UpdateStrategy SpecUpdateStrategy     `yaml:"updateStrategy"`
}
