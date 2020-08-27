#!/bin/sh

# ./docker-run.sh goapp 37000 PORT=37000

NAME=$1
PORT=$2
ENV=$3 #used to define the por where app would be used
DOCKER_IMAGE=scarrionv/simple-api
DOCKER_IMAGE_VERSION=1.0

echo "Executing docker run with parameters: name:$NAME port:$PORT env:$ENV"
docker rm "$NAME"
docker run -d -p "$PORT:$PORT" -e "PORT=$PORT" --name "$NAME" $DOCKER_IMAGE:$DOCKER_IMAGE_VERSION