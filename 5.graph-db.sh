#!/usr/bin/env bash
#$1 create/c or delete/d

if [ "$1" == "delete" ] || [ "$1" == "d" ]; then   
    echo kubectl delete -f deployments/db/graph-db.yaml
    kubectl delete -f deployments/db/graph-db.yaml
else
    echo kubectl create -f deployments/db/graph-db.yaml
    kubectl create -f deployments/db/graph-db.yaml
fi