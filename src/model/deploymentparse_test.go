package model

import (
	"gopkg.in/yaml.v3"
	"log"
	"testing"
)

const deploymentexample = `
kind: Deployment
apiVersion: apps/v1
metadata:
  annotations:
    deployment.kubernetes.io/revision: '37'
  resourceVersion: '312766041'
  name: ababil
  uid: 6842db6c-ab8c-4374-b235-bca7b82913e6
  creationTimestamp: '2023-02-22T10:08:03Z'
  generation: 49
  managedFields:
    - manager: kubectl-rollout
      operation: Update
      apiVersion: apps/v1
      time: '2023-02-22T13:42:39Z'
      fieldsType: FieldsV1
      fieldsV1:
        'f:spec':
          'f:template':
            'f:metadata':
              'f:annotations':
                .: {}
                'f:kubectl.kubernetes.io/restartedAt': {}
    - manager: Mozilla
      operation: Update
      apiVersion: apps/v1
      time: '2023-03-08T08:02:28Z'
      fieldsType: FieldsV1
      fieldsV1:
        'f:metadata':
          'f:labels':
            .: {}
            'f:name': {}
        'f:spec':
          'f:progressDeadlineSeconds': {}
          'f:revisionHistoryLimit': {}
          'f:selector': {}
          'f:strategy':
            'f:rollingUpdate':
              .: {}
              'f:maxSurge': {}
              'f:maxUnavailable': {}
            'f:type': {}
          'f:template':
            'f:metadata':
              'f:labels':
                .: {}
                'f:name': {}
              'f:name': {}
            'f:spec':
              'f:containers':
                'k:{"name":"ababil"}':
                  'f:image': {}
                  'f:startupProbe':
                    .: {}
                    'f:failureThreshold': {}
                    'f:initialDelaySeconds': {}
                    'f:periodSeconds': {}
                    'f:successThreshold': {}
                    'f:tcpSocket':
                      .: {}
                      'f:port': {}
                    'f:timeoutSeconds': {}
                  'f:volumeMounts':
                    .: {}
                    'k:{"mountPath":"/config/application.yml"}':
                      .: {}
                      'f:mountPath': {}
                      'f:name': {}
                      'f:subPath': {}
                  'f:terminationMessagePolicy': {}
                  .: {}
                  'f:resources': {}
                  'f:env':
                    .: {}
                    'k:{"name":"PIGEON_QRIS_URI"}':
                      .: {}
                      'f:name': {}
                      'f:value': {}
                    'k:{"name":"bypassAsliri"}':
                      .: {}
                      'f:name': {}
                      'f:value': {}
                    'k:{"name":"bypassDhn"}':
                      .: {}
                      'f:name': {}
                      'f:value': {}
                    'k:{"name":"bypassRiskProfile"}':
                      .: {}
                      'f:name': {}
                      'f:value': {}
                  'f:terminationMessagePath': {}
                  'f:imagePullPolicy': {}
                  'f:ports':
                    .: {}
                    'k:{"containerPort":8081,"protocol":"TCP"}':
                      .: {}
                      'f:containerPort': {}
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
                'k:{"name":"config"}':
                  .: {}
                  'f:configMap':
                    .: {}
                    'f:defaultMode': {}
                    'f:items': {}
                    'f:name': {}
                  'f:name': {}
    - manager: kube-controller-manager
      operation: Update
      apiVersion: apps/v1
      time: '2023-03-08T14:38:50Z'
      fieldsType: FieldsV1
      fieldsV1:
        'f:metadata':
          'f:annotations':
            .: {}
            'f:deployment.kubernetes.io/revision': {}
        'f:status':
          'f:availableReplicas': {}
          'f:conditions':
            .: {}
            'k:{"type":"Available"}':
              .: {}
              'f:lastTransitionTime': {}
              'f:lastUpdateTime': {}
              'f:message': {}
              'f:reason': {}
              'f:status': {}
              'f:type': {}
            'k:{"type":"Progressing"}':
              .: {}
              'f:lastTransitionTime': {}
              'f:lastUpdateTime': {}
              'f:message': {}
              'f:reason': {}
              'f:status': {}
              'f:type': {}
          'f:observedGeneration': {}
          'f:readyReplicas': {}
          'f:replicas': {}
          'f:updatedReplicas': {}
      subresource: status
  namespace: qris
  labels:
    name: ababil
spec:
  replicas: 1
  selector:
    matchLabels:
      name: ababil
  template:
    metadata:
      name: ababil
      creationTimestamp: null
      labels:
        name: ababil
      annotations:
        kubectl.kubernetes.io/restartedAt: '2023-03-08T18:41:50+07:00'
    spec:
      volumes:
        - name: config
          configMap:
            name: ababil
            items:
              - key: application.yml
                path: application.yml
            defaultMode: 420
      containers:
        - resources: {}
          terminationMessagePath: /dev/termination-log
          name: ababil
          env:
            - name: bypassDhn
              value: 'true'
            - name: bypassRiskProfile
              value: 'true'
            - name: bypassAsliri
              value: 'true'
            - name: PIGEON_QRIS_URI
              value: 'http://pigeon-vlink.qris.svc:8080/pigeon/apiv1'
          ports:
            - containerPort: 8081
              protocol: TCP
          imagePullPolicy: Always
          startupProbe:
            tcpSocket:
              port: 8081
            initialDelaySeconds: 30
            timeoutSeconds: 1
            periodSeconds: 10
            successThreshold: 1
            failureThreshold: 3
          volumeMounts:
            - name: config
              mountPath: /config/application.yml
              subPath: application.yml
          terminationMessagePolicy: File
          image: 'ababil-bsb:staging'
      restartPolicy: Always
      terminationGracePeriodSeconds: 30
      dnsPolicy: ClusterFirst
      securityContext: {}
      imagePullSecrets:
        - name: regcred
      schedulerName: default-scheduler
  strategy:
    type: RollingUpdate
    rollingUpdate:
      maxUnavailable: 25%
      maxSurge: 25%
  revisionHistoryLimit: 10
  progressDeadlineSeconds: 600
status:
  observedGeneration: 49
  replicas: 1
  updatedReplicas: 1
  readyReplicas: 1
  availableReplicas: 1
  conditions:
    - type: Progressing
      status: 'True'
      lastUpdateTime: '2023-03-08T11:43:30Z'
      lastTransitionTime: '2023-02-22T10:08:03Z'
      reason: NewReplicaSetAvailable
      message: ReplicaSet "ababil-5dddfd6668" has successfully progressed.
    - type: Available
      status: 'True'
      lastUpdateTime: '2023-03-08T14:38:50Z'
      lastTransitionTime: '2023-03-08T14:38:50Z'
      reason: MinimumReplicasAvailable
      message: Deployment has minimum availability.
`

func TestParseDeployment(t *testing.T) {
	yamlResult := Deployment{}
	err := yaml.Unmarshal([]byte(deploymentexample), &yamlResult)
	if err != nil {
		t.Error(err)
		return
	}
	log.Println(yamlResult)
}
