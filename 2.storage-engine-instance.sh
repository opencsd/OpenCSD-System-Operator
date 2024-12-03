#!/usr/bin/env bash
dir=$( pwd )

#$1 create/c or delete/d

if [ "$1" == "delete" ] || [ "$1" == "d" ]; then   
    echo kubectl delete -f deployments/opencsd/storage-engine-instance.yaml
    kubectl delete -f deployments/opencsd/storage-engine-instance.yaml
else
    echo kubectl create -f deployments/opencsd/storage-engine-instance.yaml
    kubectl create -f deployments/opencsd/storage-engine-instance.yaml
fi