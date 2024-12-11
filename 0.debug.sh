#!/bin/bash

if [ "$1" == "q" ] ; then   
    while [ -z $PODNAME ]
    do
        PODNAME=`kubectl get po -o=name --field-selector=status.phase=Running -n keti-opencsd| grep query-engine-instance`
        PODNAME="${PODNAME:4}"
    done
    kubectl logs $PODNAME -n keti-opencsd -f 
elif [ "$1" == "i" ] ; then  
    while [ -z $PODNAME ]
    do
        PODNAME=`kubectl get po -o=name --field-selector=status.phase=Running -n keti-opencsd | grep storage-engine-instance`
        echo $PODNAME
        PODNAME="${PODNAME:4}"
    done
    kubectl logs $PODNAME -n keti-opencsd -f interface-module
elif [ "$1" == "o" ] ; then  
    while [ -z $PODNAME ]
    do
        PODNAME=`kubectl get po -o=name --field-selector=status.phase=Running -n keti-opencsd | grep storage-engine-instance`
        PODNAME="${PODNAME:4}"
    done
    kubectl logs $PODNAME -n keti-opencsd -f offloading-module
elif [ "$1" == "m" ] ; then  
    while [ -z $PODNAME ]
    do
        PODNAME=`kubectl get po -o=name --field-selector=status.phase=Running -n keti-opencsd | grep storage-engine-instance`
        PODNAME="${PODNAME:4}"
    done
    kubectl logs $PODNAME -n keti-opencsd -f merging-module
elif [ "$1" == "mo" ] ; then  
    while [ -z $PODNAME ]
    do
        PODNAME=`kubectl get po -o=name --field-selector=status.phase=Running | grep storage-engine-instance`
        PODNAME="${PODNAME:4}"
    done
    kubectl logs $PODNAME -n keti-opencsd -f monitoring-module
else 
    echo arg error
fi



