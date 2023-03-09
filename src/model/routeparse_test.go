package model

import (
	"gopkg.in/yaml.v3"
	"log"
	"testing"
)

const routeexample = `
kind: Route
apiVersion: route.openshift.io/v1
metadata:
  name: main-service
  namespace: qris
  uid: 960e39ea-3636-40d9-beb5-f6b7df08f871
  resourceVersion: '254031959'
  creationTimestamp: '2023-02-13T04:27:11Z'
  labels:
    name: merchant-portal
  managedFields:
    - manager: openshift-router
      operation: Update
      apiVersion: route.openshift.io/v1
      time: '2023-02-13T04:27:11Z'
      fieldsType: FieldsV1
      fieldsV1:
        'f:status':
          'f:ingress': {}
    - manager: Mozilla
      operation: Update
      apiVersion: route.openshift.io/v1
      time: '2023-02-13T06:02:16Z'
      fieldsType: FieldsV1
      fieldsV1:
        'f:metadata':
          'f:labels':
            .: {}
            'f:name': {}
        'f:spec':
          'f:host': {}
          'f:port':
            .: {}
            'f:targetPort': {}
          'f:tls':
            .: {}
            'f:caCertificate': {}
            'f:certificate': {}
            'f:insecureEdgeTerminationPolicy': {}
            'f:key': {}
            'f:termination': {}
          'f:to':
            'f:kind': {}
            'f:name': {}
            'f:weight': {}
          'f:wildcardPolicy': {}
spec:
  host: dev1.domain
  to:
    kind: Service
    name: merchant-portal
    weight: 100
  port:
    targetPort: nginx
  tls:
    termination: edge
    certificate: |
      -----BEGIN CERTIFICATE-----
      -----END CERTIFICATE-----
    key: |
      -----BEGIN PRIVATE KEY-----
      -----END PRIVATE KEY-----
    caCertificate: |+
      -----BEGIN CERTIFICATE-----
      -----END CERTIFICATE-----
      -----BEGIN CERTIFICATE-----
      -----END CERTIFICATE-----

    insecureEdgeTerminationPolicy: Allow
  wildcardPolicy: None
status:
  ingress:
    - host: dev1.domain
      routerName: default
      conditions:
        - type: Admitted
          status: 'True'
          lastTransitionTime: '2023-02-13T04:27:11Z'
      wildcardPolicy: None
      routerCanonicalHostname: router-default.apps.ocp-dev.domain
`

func TestParseRoute(t *testing.T) {
	yamlResult := Route{}
	err := yaml.Unmarshal([]byte(routeexample), &yamlResult)
	if err != nil {
		t.Error(err)
		return
	}
	log.Println(yamlResult)
}
