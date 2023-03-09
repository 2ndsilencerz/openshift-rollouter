package model

import (
	"gopkg.in/yaml.v3"
	"log"
	"testing"
)

const statefulsetexample = `
kind: StatefulSet
apiVersion: apps/v1
metadata:
  name: openshift-rollouter
  namespace: qris
  uid: e7877649-c464-4a07-85a4-e4efacc0a051
  resourceVersion: '312664293'
  generation: 18
  creationTimestamp: '2023-03-03T04:13:14Z'
  managedFields:
    - manager: Mozilla
      operation: Update
      apiVersion: apps/v1
      time: '2023-03-08T11:43:43Z'
      fieldsType: FieldsV1
      fieldsV1:
        'f:spec':
          'f:podManagementPolicy': {}
          'f:revisionHistoryLimit': {}
          'f:selector': {}
          'f:serviceName': {}
          'f:template':
            'f:metadata':
              'f:labels':
                .: {}
                'f:name': {}
              'f:name': {}
            'f:spec':
              'f:containers':
                'k:{"name":"openshift-rollouter"}':
                  'f:image': {}
                  'f:startupProbe':
                    .: {}
                    'f:failureThreshold': {}
                    'f:periodSeconds': {}
                    'f:successThreshold': {}
                    'f:tcpSocket':
                      .: {}
                      'f:port': {}
                    'f:timeoutSeconds': {}
                  'f:volumeMounts':
                    .: {}
                    'k:{"mountPath":"/app/index.html"}':
                      .: {}
                      'f:mountPath': {}
                      'f:name': {}
                      'f:subPath': {}
                    'k:{"mountPath":"/etc/hosts"}':
                      .: {}
                      'f:mountPath': {}
                      'f:name': {}
                      'f:subPath': {}
                    'k:{"mountPath":"/opt/bitnami/nginx/conf/"}':
                      .: {}
                      'f:mountPath': {}
                      'f:name': {}
                  'f:terminationMessagePolicy': {}
                  .: {}
                  'f:resources': {}
                  'f:livenessProbe':
                    .: {}
                    'f:failureThreshold': {}
                    'f:periodSeconds': {}
                    'f:successThreshold': {}
                    'f:tcpSocket':
                      .: {}
                      'f:port': {}
                    'f:timeoutSeconds': {}
                  'f:readinessProbe':
                    .: {}
                    'f:failureThreshold': {}
                    'f:periodSeconds': {}
                    'f:successThreshold': {}
                    'f:tcpSocket':
                      .: {}
                      'f:port': {}
                    'f:timeoutSeconds': {}
                  'f:terminationMessagePath': {}
                  'f:imagePullPolicy': {}
                  'f:ports':
                    .: {}
                    'k:{"containerPort":8080,"protocol":"TCP"}':
                      .: {}
                      'f:containerPort': {}
                      'f:name': {}
                      'f:protocol': {}
                    'k:{"containerPort":8443,"protocol":"TCP"}':
                      .: {}
                      'f:containerPort': {}
                      'f:name': {}
                      'f:protocol': {}
                  'f:name': {}
              'f:dnsPolicy': {}
              'f:imagePullSecrets':
                .: {}
                'k:{"name":"regcred"}': {}
              'f:restartPolicy': {}
              'f:schedulerName': {}
              'f:securityContext': {}
              'f:terminationGracePeriodSeconds': {}
              'f:volumes':
                .: {}
                'k:{"name":"conf"}':
                  .: {}
                  'f:configMap':
                    .: {}
                    'f:defaultMode': {}
                    'f:name': {}
                  'f:name': {}
                'k:{"name":"html"}':
                  .: {}
                  'f:configMap':
                    .: {}
                    'f:defaultMode': {}
                    'f:items': {}
                    'f:name': {}
                  'f:name': {}
          'f:updateStrategy':
            'f:rollingUpdate':
              .: {}
              'f:partition': {}
            'f:type': {}
    - manager: Go-http-client
      operation: Update
      apiVersion: apps/v1
      time: '2023-03-08T12:52:29Z'
      fieldsType: FieldsV1
      fieldsV1:
        'f:spec':
          'f:template':
            'f:spec':
              'f:containers':
                'k:{"name":"openshift-rollouter"}':
                  'f:volumeMounts':
                    'k:{"mountPath":"/app/config.yaml"}':
                      .: {}
                      'f:mountPath': {}
                      'f:name': {}
                      'f:subPath': {}
              'f:volumes':
                'k:{"name":"conf"}':
                  'f:configMap':
                    'f:items': {}
    - manager: kube-controller-manager
      operation: Update
      apiVersion: apps/v1
      time: '2023-03-08T13:41:04Z'
      fieldsType: FieldsV1
      fieldsV1:
        'f:status':
          'f:currentReplicas': {}
          'f:currentRevision': {}
          'f:updatedReplicas': {}
          'f:readyReplicas': {}
          'f:replicas': {}
          'f:availableReplicas': {}
          'f:collisionCount': {}
          'f:observedGeneration': {}
          'f:updateRevision': {}
      subresource: status
spec:
  replicas: 1
  selector:
    matchLabels:
      name: openshift-rollouter
  template:
    metadata:
      name: openshift-rollouter
      creationTimestamp: null
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
            defaultMode: 420
        - name: html
          configMap:
            name: openshift-rollouter
            items:
              - key: index.html
                path: index.html
            defaultMode: 420
      containers:
        - resources: {}
          readinessProbe:
            tcpSocket:
              port: 8080
            timeoutSeconds: 10
            periodSeconds: 10
            successThreshold: 1
            failureThreshold: 3
          terminationMessagePath: /dev/termination-log
          name: openshift-rollouter
          livenessProbe:
            tcpSocket:
              port: 8080
            timeoutSeconds: 10
            periodSeconds: 10
            successThreshold: 1
            failureThreshold: 3
          ports:
            - name: main-port
              containerPort: 8080
              protocol: TCP
            - name: ssl-port
              containerPort: 8443
              protocol: TCP
          imagePullPolicy: Always
          startupProbe:
            tcpSocket:
              port: 8080
            timeoutSeconds: 10
            periodSeconds: 10
            successThreshold: 1
            failureThreshold: 3
          volumeMounts:
            - name: conf
              mountPath: /opt/bitnami/nginx/conf/
            - name: html
              mountPath: /app/index.html
              subPath: index.html
            - name: conf
              mountPath: /etc/hosts
              subPath: hosts
            - name: conf
              mountPath: /app/config.yaml
              subPath: config.yaml
          terminationMessagePolicy: File
          image: 'openshift-rollouter:latest'
      restartPolicy: Always
      terminationGracePeriodSeconds: 30
      dnsPolicy: ClusterFirst
      securityContext: {}
      imagePullSecrets:
        - name: regcred
      schedulerName: default-scheduler
  serviceName: openshift-rollouter
  podManagementPolicy: OrderedReady
  updateStrategy:
    type: RollingUpdate
    rollingUpdate:
      partition: 0
  revisionHistoryLimit: 10
status:
  observedGeneration: 18
  availableReplicas: 1
  updateRevision: openshift-rollouter-65bfd66999
  currentRevision: openshift-rollouter-65bfd66999
  currentReplicas: 1
  updatedReplicas: 1
  replicas: 1
  collisionCount: 0
  readyReplicas: 1
`

func TestParseStatefulSet(t *testing.T) {
	yamlResult := StatefulSet{}
	err := yaml.Unmarshal([]byte(statefulsetexample), &yamlResult)
	if err != nil {
		t.Error(err)
		return
	}
	log.Println(yamlResult)
	log.Println()
}
