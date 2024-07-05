#!/bin/bash

if [ "$1" == "q" ] ; then   
    while [ -z $PODNAME ]
    do
        PODNAME=`kubectl get po -o=name -A --field-selector=status.phase=Running | grep query-engine-instance`
        PODNAME="${PODNAME:4}"
    done
    kubectl logs $PODNAME -n keti-opencsd -f 
elif [ "$1" == "i" ] ; then  
    while [ -z $PODNAME ]
    do
        PODNAME=`kubectl get po -o=name -A --field-selector=status.phase=Running | grep storage-engine-instance`
        PODNAME="${PODNAME:4}"
    done
    kubectl logs $PODNAME -n keti-opencsd -f interface-container
elif [ "$1" == "o" ] ; then  
    while [ -z $PODNAME ]
    do
        PODNAME=`kubectl get po -o=name -A --field-selector=status.phase=Running | grep storage-engine-instance`
        PODNAME="${PODNAME:4}"
    done
    kubectl logs $PODNAME -n keti-opencsd -f offloading-container
elif [ "$1" == "m" ] ; then  
    while [ -z $PODNAME ]
    do
        PODNAME=`kubectl get po -o=name -A --field-selector=status.phase=Running | grep storage-engine-instance`
        PODNAME="${PODNAME:4}"
    done
    kubectl logs $PODNAME -n keti-opencsd -f merging-container
elif [ "$1" == "mo" ] ; then  
    while [ -z $PODNAME ]
    do
        PODNAME=`kubectl get po -o=name -A --field-selector=status.phase=Running | grep storage-engine-instance`
        PODNAME="${PODNAME:4}"
    done
    kubectl logs $PODNAME -n keti-opencsd -f monitoring-container
else 
    echo arg error
fi



