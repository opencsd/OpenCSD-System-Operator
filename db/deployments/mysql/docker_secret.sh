#!/bin/bash

kubectl create secret generic "regcred" \
    --from-file=.dockerconfigjson=/root/.docker/config.json \
    --type=kubernetes.io/dockerconfigjson \
    --namespace=MYSQL_NAMESPACE
