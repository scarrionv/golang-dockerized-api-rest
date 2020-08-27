# Golang Dockerized API Rest

Simple Golang API Rest dockerized.


## Build Go app

```
$ golang-dockerized-api-rest/bin/go-build-app.sh
```

## Build docker container

```
$ golang-dockerized-api-rest/docker/docker-build.sh
```

## Run docker container

```
$ golang-dockerized-api-rest/docker/docker-run.sh <container-name> <port> <env-variables>
```

## Test

```
$ curl http://localhost:<port>/users
```

