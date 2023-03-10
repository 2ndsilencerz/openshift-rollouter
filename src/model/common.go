package model

//type Common struct {
//	ApiVersion string `yaml:"apiVersion"`
//	Kind       string `yaml:"kind"`
//	Labels     Label  `yaml:"labels"`
//}

type Metadata struct {
	Name            string            `yaml:"name" json:"name"`
	ResourceVersion string            `yaml:"resourceVersion" json:"resourceVersion"`
	Labels          map[string]string `yaml:"labels" json:"labels"`
	Annotations     map[string]string `yaml:"annotations" json:"annotations"`
	Generation      int               `yaml:"generation" json:"generation"`
}

type SpecSelector struct {
	MatchLabels MatchLabels `yaml:"matchLabels"`
}

type MatchLabels map[string]interface{}

type SpecVolumes []SpecVolume
type SpecVolume struct {
	Name      string          `yaml:"name"`
	ConfigMap VolumeConfigMap `yaml:"configMap"`
}

type VolumeConfigMapItems []VolumeConfigMapItem
type VolumeConfigMap struct {
	Name  string               `yaml:"name"`
	Items VolumeConfigMapItems `yaml:"items"`
}

type VolumeConfigMapItem struct {
	Key  string `yaml:"key"`
	Path string `yaml:"path"`
}

type ImagePullSecrets []ImagePullSecret
type ImagePullSecret struct {
	Name string `yaml:"name"`
}

type SpecContainers []SpecContainer
type SpecContainer struct {
	// Name should be unique between SpecContainers
	Name string `yaml:"name"`
	// Image defined for image that should be used to run
	Image string `yaml:"image"`
	// ImagePullPolicy selection is either "Always", "IfNotPresent", and "Never"
	ImagePullPolicy string                 `yaml:"imagePullPolicy"`
	Ports           SpecContainerPorts     `yaml:"ports"`
	StartupProbe    SpecContainerProbe     `yaml:"startupProbe"`
	LivenessProbe   SpecContainerProbe     `yaml:"livenessProbe"`
	ReadinessProbe  SpecContainerProbe     `yaml:"readinessProbe"`
	Resources       SpecContainerResources `yaml:"resources"`
	Env             SpecContainerEnvs      `yaml:"env"`
	EnvFrom         SpecContainerEnvFroms  `yaml:"envFrom"`
	SecurityContext map[string]interface{} `yaml:"securityContext"`
	// DnsPolicy is either "ClusterFirst", "ClusterFirstWithHostNet", "Default", or "None"
	DnsPolicy string `yaml:"dnsPolicy"`
	// RestartPolicy is either "Always", "OnFailure", or "Never"
	RestartPolicy string `yaml:"restartPolicy"`
}

type SpecContainerPorts []SpecContainerPort
type SpecContainerPort struct {
	// ContainerPort should be fill with the exposed port for incoming traffic
	ContainerPort int `yaml:"containerPort"`
	// Protocol selection is either "TCP" or "UDP"
	Protocol string `yaml:"protocol"`
	// Name should be unique between SpecContainerPorts
	Name string `yaml:"name"`
}

type SpecContainerVolumeMount struct {
	// MountPath is defined either a directory location or file location
	MountPath string `yaml:"mountPath"`
	// Name should match with the one listed in SpecVolumes
	Name string `yaml:"name"`
	// SubPath is required if MountPath defined until file location
	SubPath string `yaml:"subPath"`
}

// SpecContainerProbe should select either SpecContainerProbeTcpSocket or SpecContainerProbeHttpGet
// but not both
type SpecContainerProbe struct {
	TcpSocket SpecContainerProbeTcpSocket `yaml:"tcpSocket"`
	HttpGet   SpecContainerProbeHttpGet   `yaml:"httpGet"`
}

// SpecContainerProbeTcpSocket to detect probe by checking tcp socket exposed by container
type SpecContainerProbeTcpSocket struct {
	Port int `yaml:"port"`
}

// SpecContainerProbeHttpGet to detect probe by checking api exposed by container
type SpecContainerProbeHttpGet struct {
	Port int    `yaml:"port"`
	Path string `yaml:"path"`
}

type SpecContainerResources struct {
	// Requests to define what is the minimum resources to allocate to kubernetes.
	Requests SpecContainerResourcesType `yaml:"requests"`
	// Limits to define what is the maximum resources to allocate to kubernetes.
	// The value presented in metric may exceed the defined.
	// Be careful, setting this might slow or crash the application.
	Limits SpecContainerResourcesType `yaml:"limits"`
}

type SpecContainerResourcesType struct {
	// Memory defined with subfix like Ki, Mi, Gi, etc
	Memory string `yaml:"memory"`
	// Cpu defined with subfix m. It's recommended to set it with milli scale, hence the subfix m
	Cpu string `yaml:"cpu"`
}

type SpecContainerEnvs []SpecContainerEnv
type SpecContainerEnv struct {
	Name  string `yaml:"name"`
	Value string `yaml:"value"`
}

type SpecContainerEnvFroms []SpecContainerEnvFrom
type SpecContainerEnvFrom struct {
	ConfigMapRef SpecContainerEnvFromResourceRef `yaml:"configMapRef"`
	SecretRef    SpecContainerEnvFromResourceRef `yaml:"secretRef"`
}

type SpecContainerEnvFromResourceRef struct {
	// Name should match with ConfigMap or Secret
	Name string `yaml:"name"`
}

type SpecUpdateStrategy struct {
	// Type is either "RollingUpdate" (for most deployment), "Recreate", "OnDelete",
	Type          string                          `yaml:"type"`
	RollingUpdate SpecUpdateStrategyRollingUpdate `yaml:"rollingUpdate"`
}

type SpecUpdateStrategyRollingUpdate struct {
	MaxUnavailable int `yaml:"maxUnavailable"`
	MaxSurge       int `yaml:"maxSurge"`
}
