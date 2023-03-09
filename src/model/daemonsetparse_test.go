package model

import (
	"gopkg.in/yaml.v3"
	"log"
	"testing"
)

const daemonsetexample = `
kind: DaemonSet
apiVersion: apps/v1
metadata:
  name: multitool
  namespace: qris
  uid: ceea9a3e-54de-44e5-9efb-930f34fcdbea
  resourceVersion: '296392991'
  generation: 1
  creationTimestamp: '2023-02-06T07:53:30Z'
  annotations:
    deprecated.daemonset.template.generation: '1'
  managedFields:
    - manager: Mozilla
      operation: Update
      apiVersion: apps/v1
      time: '2023-02-06T07:53:30Z'
      fieldsType: FieldsV1
      fieldsV1:
        'f:metadata':
          'f:annotations':
            .: {}
            'f:deprecated.daemonset.template.generation': {}
        'f:spec':
          'f:revisionHistoryLimit': {}
          'f:selector': {}
          'f:template':
            'f:metadata':
              'f:labels':
                .: {}
                'f:app': {}
            'f:spec':
              'f:containers':
                'k:{"name":"multitool"}':
                  .: {}
                  'f:image': {}
                  'f:imagePullPolicy': {}
                  'f:name': {}
                  'f:resources': {}
                  'f:terminationMessagePath': {}
                  'f:terminationMessagePolicy': {}
              'f:dnsPolicy': {}
              'f:restartPolicy': {}
              'f:schedulerName': {}
              'f:securityContext': {}
              'f:terminationGracePeriodSeconds': {}
          'f:updateStrategy':
            'f:rollingUpdate':
              .: {}
              'f:maxSurge': {}
              'f:maxUnavailable': {}
            'f:type': {}
    - manager: kube-controller-manager
      operation: Update
      apiVersion: apps/v1
      time: '2023-03-02T02:37:23Z'
      fieldsType: FieldsV1
      fieldsV1:
        'f:status':
          'f:currentNumberScheduled': {}
          'f:desiredNumberScheduled': {}
          'f:numberAvailable': {}
          'f:numberReady': {}
          'f:observedGeneration': {}
          'f:updatedNumberScheduled': {}
      subresource: status
spec:
  selector:
    matchLabels:
      app: multitool
  template:
    metadata:
      creationTimestamp: null
      labels:
        app: multitool
    spec:
      containers:
        - name: multitool
          image: docker.io/wbitt/network-multitool
          resources: {}
          terminationMessagePath: /dev/termination-log
          terminationMessagePolicy: File
          imagePullPolicy: Always
      restartPolicy: Always
      terminationGracePeriodSeconds: 30
      dnsPolicy: ClusterFirst
      securityContext: {}
      schedulerName: default-scheduler
  updateStrategy:
    type: RollingUpdate
    rollingUpdate:
      maxUnavailable: 1
      maxSurge: 0
  revisionHistoryLimit: 10
status:
  currentNumberScheduled: 5
  numberMisscheduled: 0
  desiredNumberScheduled: 5
  numberReady: 5
  observedGeneration: 1
  updatedNumberScheduled: 5
  numberAvailable: 5
`

func TestParseDaemonSet(t *testing.T) {
	yamlResult := DaemonSet{}
	err := yaml.Unmarshal([]byte(daemonsetexample), &yamlResult)
	if err != nil {
		t.Error(err)
		return
	}
	log.Println(yamlResult)
}
