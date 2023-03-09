package model

type Service struct {
	ApiVersion string      `yaml:"apiVersion"`
	Kind       string      `yaml:"kind"`
	Metadata   Metadata    `yaml:"metadata"`
	Spec       ServiceSpec `yaml:"spec"`
}

type ServiceSpec struct {
	// Type is either "ExternalName", "ClusterIP", "NodePort", and "LoadBalancer"
	Type        string            `yaml:"type"`
	Selector    map[string]string `yaml:"selector"`
	Ports       ServiceSpecPorts  `yaml:"ports"`
	ExternalIPs []string          `yaml:"externalIPs"`
}

type ServiceSpecPorts []ServiceSpecPort
type ServiceSpecPort struct {
	Port       int    `yaml:"port"`
	TargetPort string `yaml:"targetPort"`
	Name       string `yaml:"name"`
	Protocol   string `yaml:"protocol"`
}
