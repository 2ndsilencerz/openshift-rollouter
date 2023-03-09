package model

type Ingress struct {
	ApiVersion string      `yaml:"apiVersion"`
	Kind       string      `yaml:"kind"`
	Metadata   Metadata    `yaml:"metadata"`
	Spec       IngressSpec `yaml:"spec"`
}

type IngressSpec struct {
	DefaultBackend IngressSpecDefaultBackend `yaml:"defaultBackend"`
	Rules          IngressSpecRules          `yaml:"rules"`
	Tls            IngressSpecTlsArr         `yaml:"tls"`
}

type IngressSpecDefaultBackend struct {
	Service IngressService `yaml:"service"`
}

type IngressService struct {
	Name string                               `yaml:"name"`
	Port IngressSpecDefaultBackendServicePort `yaml:"port"`
}

type IngressSpecDefaultBackendServicePort struct {
	Number string `yaml:"number"`
}

type IngressSpecRules []IngressSpecRule
type IngressSpecRule struct {
	Host string              `yaml:"host"`
	Http IngressSpecRuleHttp `yaml:"http"`
}

type IngressSpecRuleHttp struct {
	Paths IngressSpecRuleHttpPaths `yaml:"paths"`
}

type IngressSpecRuleHttpPaths []IngressSpecRuleHttpPath
type IngressSpecRuleHttpPath struct {
	Backend  IngressSpecDefaultBackend `yaml:"backend"`
	Path     string                    `yaml:"path"`
	PathType string                    `yaml:"pathType"`
}

type IngressSpecRuleHttpPathBackend struct {
	Service IngressService `yaml:"service"`
}

type IngressSpecTlsArr []IngressSpecTls
type IngressSpecTls struct {
	Hosts      []string `yaml:"hosts"`
	SecretName string   `yaml:"secretName"`
}
