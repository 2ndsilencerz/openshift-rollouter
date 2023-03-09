package openshift

import (
	"testing"
)

const yamlExample = `
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
        <p style="text-align: center" >© Multipolar Technology Tbk.</p>
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

---
apiVersion: apps/v1
kind: StatefulSet
metadata:
  name: openshift-rollouter
spec:
  selector:
    matchLabels:
      name: openshift-rollouter
  serviceName: openshift-rollouter
  template:
    metadata:
      name: openshift-rollouter
      labels:
        name: openshift-rollouter
    spec:
      volumes:
        - name: conf
          configMap:
            name: openshift-rollouter
            items:
              - key: nginx.conf
                path: nginx.conf
              - key: hosts
                path: hosts
              - key: config.yaml
                path: config.yaml
        - name: html
          configMap:
            name: openshift-rollouter
            items:
              - key: index.html
                path: index.html
      imagePullSecrets:
        - name: regcred
      containers:
        - name: openshift-rollouter
          image: docker.io/visiondgmlpt/openshift-rollouter:latest
          imagePullPolicy: Always
          ports:
            - containerPort: 8080
              protocol: TCP
              name: main-port
            - containerPort: 8443
              protocol: TCP
              name: ssl-port
          volumeMounts:
            - mountPath: /opt/bitnami/nginx/conf/
              name: conf
            - mountPath: /app/index.html
              subPath: index.html
              name: html
            - mountPath: /etc/hosts
              name: conf
              subPath: hosts
            - mountPath: /app/config.yaml
              name: conf
              subPath: config.yaml
          startupProbe:
            tcpSocket:
              port: 8080
            timeoutSeconds: 10
          livenessProbe:
            tcpSocket:
              port: 8080
            timeoutSeconds: 10
          readinessProbe:
            tcpSocket:
              port: 8080
            timeoutSeconds: 10
---
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
      name: main-port
---
`

func TestGetApiVersion(t *testing.T) {
	contents := ReadYaml([]byte(yamlExample))
	resources := Categorize(contents)
	correct := false
	if resources.ConfigMap[0].ApiVersion == "v1" {
		correct = true
	}
	//log.Println(ApiVersion(resources.ConfigMap[0]))
	if resources.ConfigMap[0].Kind == "apps/v1" {
		correct = true
	}
	//log.Println(ApiVersion(resources.StatefulSet[0]))

	if !correct {
		t.Error("Failed")
	}
}
