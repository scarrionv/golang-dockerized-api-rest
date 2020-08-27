# Golang Dockerized API Rest

Simple Golang API Rest dockerized


## Build Go app

```
$ simple-api/bin/go-build-app.sh
```

## Build docker container

```
$ simple-api/bin/docker-build.sh
```

## Run docker container

```
$ simple-api/bin/docker-run.sh <container-name> <port> <env-variables>
```

## Test

```
$ curl http://localhost:<port>/users
```

