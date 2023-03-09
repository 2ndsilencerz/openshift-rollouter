package model

type Route struct {
	ApiVersion string    `yaml:"apiVersion"`
	Kind       string    `yaml:"kind"`
	Metadata   Metadata  `yaml:"metadata"`
	Spec       RouteSpec `yaml:"spec"`
}

type RouteSpec struct {
	Host           string        `yaml:"host"`
	To             RouteSpecTo   `yaml:"to"`
	Port           RouteSpecPort `yaml:"port"`
	Tls            RouteSpecTls  `yaml:"tls"`
	WildcardPolicy string        `yaml:"wildcardPolicy"`
}

type RouteSpecTo struct {
	Kind   string `yaml:"kind"`
	Name   string `yaml:"name"`
	Weight string `yaml:"weight"`
}

type RouteSpecPort struct {
	TargetPort string `yaml:"targetPort"`
}

type RouteSpecTls struct {
	Termination                   string `yaml:"termination"`
	Certificate                   string `yaml:"certificate"`
	Key                           string `yaml:"key"`
	CaCertificate                 string `yaml:"caCertificate"`
	InsecureEdgeTerminationPolicy string `yaml:"insecureEdgeTerminationPolicy"`
}
