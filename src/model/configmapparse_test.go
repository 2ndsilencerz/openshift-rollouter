package model

import (
	"gopkg.in/yaml.v3"
	"log"
	"testing"
)

const configmapexample = `
apiVersion: v1
kind: ConfigMap
metadata:
  name: openshift-rollouter
data:
  nginx.conf: |
    events {
      worker_connections  1024;
    }
    http {
      upstream default {
        server https://bastion:6443;
      }
      server {
        listen       8080;
          location / {
    #         return 200;
            proxy_pass default;
          }
      }
    }
  index.html: |
    <header>
        <title>Nginx Test</title>
    </header>
    <body>
        <h1 style="text-align: center">Landing Page</h1>
    </body>
    <footer>
        <p style="text-align: center" >©</p>
    </footer>
  hosts: |
    127.0.0.1       localhost
    ::1     localhost ip6-localhost ip6-loopback
    fe00::0 ip6-localnet
    fe00::0 ip6-mcastprefix
    fe00::1 ip6-allnodes
    fe00::2 ip6-allrouters
    172.17.28.101 api-int.ocp-dev.domain
    172.17.28.101 api.ocp-dev.domain
    172.17.28.101 oauth-openshift.apps.ocp-dev.domain
    172.17.28.101 bastion
  config.yaml: |
    server:
      port: 8080
    openshift:
      auth:
        uri: "https://oauth-openshift.apps.ocp-dev.domain/oauth/authorize?client_id=openshift-challenging-client&response_type=token"
        username: "admin"
        password: "P4ssw0rd1!"
      api:
        apply:
          #uri: "https://api-int.ocp-dev.domain:6443/api/v1/namespaces/<namespace>/apply"
          #uri: https://api-int.ocp-dev.domain:6443/api/v1/namespaces/<namespace>/<resource>
          uri: "https://api-int.ocp-dev.domain:6443/api/<apiVersion>/namespaces/<namespace>/<resource>"
`

func TestParseConfigMap(t *testing.T) {
	yamlResult := ConfigMap{}
	err := yaml.Unmarshal([]byte(configmapexample), &yamlResult)
	if err != nil {
		t.Error(err)
		return
	}
	log.Println(yamlResult)
	if len(yamlResult.Data["config.yaml"]) <= 0 {
		t.Error("Failed")
	}
}
