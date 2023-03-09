package model

type Deployment struct {
	ApiVersion string         `yaml:"apiVersion"`
	Kind       string         `yaml:"kind"`
	Metadata   Metadata       `yaml:"metadata"`
	Spec       DeploymentSpec `yaml:"spec"`
}

type DeploymentSpec struct {
	Replicas int                    `yaml:"replicas"`
	Selector SpecSelector           `yaml:"selector"`
	Template DeploymentSpecTemplate `yaml:"template"`
}

type DeploymentSpecTemplate struct {
	Metadata Metadata           `yaml:"metadata"`
	Spec     DeploymentSpecSpec `yaml:"spec"`
}

type DeploymentSpecSpec struct {
	Volumes          SpecVolumes        `yaml:"volumes"`
	ImagePullSecrets ImagePullSecrets   `yaml:"imagePullSecrets"`
	Containers       SpecContainers     `yaml:"containers"`
	Strategy         SpecUpdateStrategy `yaml:"strategy"`
}
