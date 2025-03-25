#!/usr/bin/env bash
#$1 create/c or delete/d

if [ "$1" == "delete" ] || [ "$1" == "d" ]; then   
    echo kubectl delete -f deployments/keti-mysql/query-engine-instance.yaml
    kubectl delete -f deployments/keti-mysql/query-engine-instance.yaml
else
    echo kubectl create -f deployments/keti-mysql/query-engine-instance.yaml
    kubectl create -f deployments/keti-mysql/query-engine-instance.yaml
fi