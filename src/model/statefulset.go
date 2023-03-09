package model

type StatefulSet struct {
	ApiVersion string          `yaml:"apiVersion"`
	Kind       string          `yaml:"kind"`
	Metadata   Metadata        `yaml:"metadata"`
	Spec       StatefulSetSpec `yaml:"spec"`
}

type StatefulSetSpec struct {
	Replicas    int                    `yaml:"replicas"`
	Selector    SpecSelector           `yaml:"selector"`
	ServiceName string                 `yaml:"serviceName"`
	Template    DeploymentSpecTemplate `yaml:"template"`
}
