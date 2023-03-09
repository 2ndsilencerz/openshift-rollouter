package model

import (
	"gopkg.in/yaml.v3"
	"log"
	"testing"
)

const cronjobexample = `
kind: CronJob
apiVersion: batch/v1beta1
metadata:
  name: ava-sch
  namespace: cob
  uid: c8ee91db-dc8f-4f55-97ff-0613498d7745
  resourceVersion: '25177129'
  generation: 6
  creationTimestamp: '2022-04-27T08:32:56Z'
  annotations: {}
spec:
  schedule: '*/30 * * * *'
  concurrencyPolicy: Allow
  suspend: true
  jobTemplate:
    metadata:
      creationTimestamp: null
    spec:
      template:
        metadata:
          creationTimestamp: null
        spec:
          containers:
            - name: ava-sch
              image: 'ava-sch:prod-bsb'
              envFrom:
                - configMapRef:
                    name: cob-cm
              resources: {}
              terminationMessagePath: /dev/termination-log
              terminationMessagePolicy: File
              imagePullPolicy: Always
          restartPolicy: Never
          terminationGracePeriodSeconds: 30
          dnsPolicy: ClusterFirst
          securityContext: {}
          imagePullSecrets:
            - name: regcred
          schedulerName: default-scheduler
  successfulJobsHistoryLimit: 1
  failedJobsHistoryLimit: 1
status:
  lastScheduleTime: '2022-05-16T15:30:00Z'
  lastSuccessfulTime: '2022-05-14T13:00:41Z'
`

func TestParseCronJob(t *testing.T) {
	cronJob := CronJob{}
	err := yaml.Unmarshal([]byte(cronjobexample), &cronJob)
	if err != nil {
		t.Error(err)
		return
	}
	log.Println(cronJob)
}
