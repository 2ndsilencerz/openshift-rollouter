package model

import (
	"gopkg.in/yaml.v3"
	"log"
	"testing"
)

const ingressexample = `
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  annotations:
    cert-manager.io/acme-challenge-type: dns01
    cert-manager.io/issue-temporary-certificate: "true"
    cert-manager.io/issuer: example-issuer-prod
    kubernetes.io/ingress.allow-http: "true"
    kubernetes.io/ingress.class: nginx
  creationTimestamp: "2023-01-24T07:31:43Z"
  generation: 1
  labels:
    name: bitnami-nginx
  name: bitnami-nginx
  namespace: merchant-bsb
  resourceVersion: "16817969"
  uid: 030ee56a-00bc-4851-a951-576858f5d35d
spec:
  defaultBackend:
    service:
      name: bitnami-nginx
      port:
        number: 80
  rules:
  - host: nginx.domain
    http:
      paths:
      - backend:
          service:
            name: bitnami-nginx
            port:
              number: 80
        path: /
        pathType: ImplementationSpecific
  tls:
  - hosts:
    - nginx.domain
    secretName: nginx-tls-new
`

func TestParseIngress(t *testing.T) {
	yamlResult := Ingress{}
	err := yaml.Unmarshal([]byte(ingressexample), &yamlResult)
	if err != nil {
		t.Error(err)
		return
	}
	log.Println(yamlResult)
}
