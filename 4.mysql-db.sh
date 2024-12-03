#!/usr/bin/env bash
#$1 create/c or delete/d

if [ "$1" == "delete" ] || [ "$1" == "d" ]; then   
    echo kubectl delete -f deployments/db/mysql-db.yaml
    kubectl delete -f deployments/db/mysql-db.yaml
else
    echo kubectl create -f deployments/db/mysql-db.yaml
    kubectl create -f deployments/db/mysql-db.yaml
fi