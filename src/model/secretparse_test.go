package model

import (
	"gopkg.in/yaml.v3"
	"log"
	"testing"
)

const secretexample = `
kind: Secret
apiVersion: v1
metadata:
  name: regcred
  namespace: qris
  uid: 9bbbbf29-c90a-418b-b541-711461958c6f
  resourceVersion: '241206913'
  creationTimestamp: '2023-01-31T09:38:51Z'
  managedFields:
    - manager: Mozilla
      operation: Update
      apiVersion: v1
      time: '2023-01-31T09:38:51Z'
      fieldsType: FieldsV1
      fieldsV1:
        'f:data':
          .: {}
          'f:.dockerconfigjson': {}
        'f:type': {}
data:
  .dockerconfigjson: >-
    eyJhdXRocyI6eyJodHRwczovL2luZGV4LmRvY2tlci5pby92MS8iOnsidXNlcm5hbWUiOiJuYWl2YXRrbyIsInBhc3N3b3JkIjoiMzNmZTljZWEtZTNlOC00YjE1LTkwYzgtMGFhYTcxYjZhMGFhIiwiYXV0aCI6ImJtRnBkbUYwYTI4Nk16Tm1aVGxqWldFdFpUTmxPQzAwWWpFMUxUa3dZemd0TUdGaFlUY3hZalpoTUdGaCJ9fX0=
type: kubernetes.io/dockerconfigjson
`

func TestParseSecret(t *testing.T) {
	yamlResult := Secret{}
	err := yaml.Unmarshal([]byte(secretexample), &yamlResult)
	if err != nil {
		t.Error(err)
		return
	}
	log.Println(yamlResult)
}
