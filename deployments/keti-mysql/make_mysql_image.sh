#!/bin/bash
docker_id="ketidevit2"
controller_name="custom-mysql"
tag="5.6"

docker build -t $docker_id/$controller_name:$tag . && \
docker push $docker_id/$controller_name:$tag
