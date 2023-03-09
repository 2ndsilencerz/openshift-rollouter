package model

import (
	"gopkg.in/yaml.v3"
	"log"
	"testing"
)

const serviceexample = `
apiVersion: v1
kind: Service
metadata:
  name: openshift-rollouter
  labels:
    name: openshift-rollouter
spec:
  type: ClusterIP
  selector:
    name: openshift-rollouter
  ports:
    - port: 443
      protocol: TCP
      targetPort: ssl-port
      name: ssl-port
    - port: 80
      protocol: TCP
      targetPort: main-port
      name: main-port`

func TestParseService(t *testing.T) {
	yamlResult := Service{}
	err := yaml.Unmarshal([]byte(serviceexample), &yamlResult)
	if err != nil {
		t.Error(err)
		return
	}
	log.Println(yamlResult)
}
