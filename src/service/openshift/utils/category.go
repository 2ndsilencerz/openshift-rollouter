package utils

import (
	"gopkg.in/yaml.v3"
	"openshift-rollouter/model"
	"strings"
)

type Bytes []byte

type Resource struct {
	Deployment       []model.Deployment
	StatefulSet      []model.StatefulSet
	DaemonSet        []model.DaemonSet
	CronJob          []model.CronJob
	DeploymentConfig []model.DeploymentConfig

	ConfigMap []model.ConfigMap
	Secret    []model.Secret

	Service []model.Service
	//Ingress []model.Ingress
	Route []model.Route
}

type ResourceDescription struct {
	ApiVersion string         `yaml:"apiVersion"`
	Kind       string         `yaml:"kind"`
	Metadata   model.Metadata `yaml:"metadata"`
}

func Categorize(contents []string) Resource {
	resource := Resource{
		Deployment:       *new([]model.Deployment),
		StatefulSet:      *new([]model.StatefulSet),
		DaemonSet:        *new([]model.DaemonSet),
		CronJob:          *new([]model.CronJob),
		DeploymentConfig: *new([]model.DeploymentConfig),
		ConfigMap:        *new([]model.ConfigMap),
		Secret:           *new([]model.Secret),
		Service:          *new([]model.Service),
		//Ingress:          *new([]Bytes),
		Route: *new([]model.Route),
	}

	for _, v := range contents {
		if strings.Contains(v, "kind: Deployment") {
			yamlRes := model.Deployment{}
			_ = yaml.Unmarshal([]byte(v), &yamlRes)
			resource.Deployment = append(resource.Deployment, yamlRes)
		} else if strings.Contains(v, "kind: StatefulSet") {
			yamlRes := model.StatefulSet{}
			_ = yaml.Unmarshal([]byte(v), &yamlRes)
			resource.StatefulSet = append(resource.StatefulSet, yamlRes)
		} else if strings.Contains(v, "kind: DaemonSet") {
			yamlRes := model.DaemonSet{}
			_ = yaml.Unmarshal([]byte(v), &yamlRes)
			resource.DaemonSet = append(resource.DaemonSet, yamlRes)
		} else if strings.Contains(v, "kind: CronJob") {
			yamlRes := model.CronJob{}
			_ = yaml.Unmarshal([]byte(v), &yamlRes)
			resource.CronJob = append(resource.CronJob, yamlRes)
		} else if strings.Contains(v, "kind: DeploymentConfig") {
			yamlRes := model.DeploymentConfig{}
			_ = yaml.Unmarshal([]byte(v), &yamlRes)
			resource.DeploymentConfig = append(resource.DeploymentConfig, yamlRes)
		} else if strings.Contains(v, "kind: ConfigMap") {
			yamlRes := model.ConfigMap{}
			_ = yaml.Unmarshal([]byte(v), &yamlRes)
			resource.ConfigMap = append(resource.ConfigMap, yamlRes)
		} else if strings.Contains(v, "kind: Secret") {
			yamlRes := model.Secret{}
			_ = yaml.Unmarshal([]byte(v), &yamlRes)
			resource.Secret = append(resource.Secret, yamlRes)
		} else if strings.Contains(v, "kind: Service") {
			yamlRes := model.Service{}
			_ = yaml.Unmarshal([]byte(v), &yamlRes)
			resource.Service = append(resource.Service, yamlRes)
			//} else if strings.Contains(v, "kind: Ingress") {
			//	resource.Ingress = append(resource.Ingress, Bytes(v))
		} else if strings.Contains(v, "kind: Route") {
			yamlRes := model.Route{}
			_ = yaml.Unmarshal([]byte(v), &yamlRes)
			resource.Route = append(resource.Route, yamlRes)
		}
	}

	return resource
}

//func ApiVersion(content []byte) string {
//	//var yamlResult map[string]string
//	yamlResult := &ResourceDescription{}
//	_ = yaml.Unmarshal(content, &yamlResult)
//	//pattern := regexp.MustCompile("apiVersion: (.*)\n")
//	//res := pattern.FindAllString(content, -1)
//	//for _, v := range res {
//	//	result = v
//	//}
//	////log.Println(result)
//	//result = strings.Replace(result, "apiVersion: ", "", -1)
//	////log.Println(result)
//	//result = strings.TrimSpace(result)
//
//	return yamlResult.ApiVersion
//}
//
//func ResourceName(content []byte) string {
//	yamlResult := ResourceDescription{}
//	_ = yaml.Unmarshal(content, &yamlResult)
//	//log.Println(yamlResult)
//	log.Println(yamlResult.Metadata.Name)
//	return yamlResult.Metadata.Name
//}
//
//func AddResourceVersion(content []byte, resourceVersion string) []byte {
//	var yamlContent map[string]interface{}
//	_ = yaml.Unmarshal(content, &yamlContent)
//	yamlContent["metadata.resourceVersion"] = resourceVersion
//	repack, _ := yaml.Marshal(yamlContent)
//	return repack
//}
