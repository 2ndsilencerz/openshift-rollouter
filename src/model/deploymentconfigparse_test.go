package model

import (
	"gopkg.in/yaml.v3"
	"log"
	"testing"
)

const deploymentconfigexample = `
apiVersion: apps.openshift.io/v1
kind: DeploymentConfig
metadata:
  name: example
  namespace: qris
spec:
  selector:
    app: httpd
  replicas: 3
  template:
    metadata:
      labels:
        app: httpd
    spec:
      containers:
        - name: httpd
          image: >-
            image-registry.openshift-image-registry.svc:5000/openshift/httpd:latest
          ports:
            - containerPort: 8080
`

func TestParseDeploymentConfig(t *testing.T) {
	yamlResult := DeploymentConfig{}
	err := yaml.Unmarshal([]byte(deploymentconfigexample), &yamlResult)
	if err != nil {
		t.Error(err)
		return
	}
	log.Println(yamlResult)
}
