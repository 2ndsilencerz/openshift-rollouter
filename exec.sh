#! /bin/bash

cp /root/.kube/config-copy /root/.kube/config
echo $password | /usr/bin/oc login -u $username --insecure-skip-tls-verify $host
exec /app/app
