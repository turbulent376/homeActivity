# Activity service

## Description
**activity** service is responsible for home activity keeping information

## Installation

* make sure you have `make` installed
* make sure you have `golangci-lint`, `goose` installed
* populate env variables according to your environment

###### Local Build
````
# build a service
make build

# check and load dependencies
make dep

# apply lint
make lint
````

###### Docker Build & Push
````
# build an image with a tag specified by $(DOCKER_TAG)
make docker-build

# pash image with a tag specified by $(DOCKER_TAG)
make docker-push
````

###### Database migrations
````
# run an initialization script (create and configure your secrice's schema)
make db-init-schem

# current version
make db-status

# apply migratiosn up to the last version
make db-up

# unapply migrations
make db-down
````

## Environment variables

|name|description|default value|
|----|-----------|-------------|
|FOCROOT|Folder where all your services are cloned||
|BITBUCKET_USER|bitbucket account. Used for loading go modules from private bitbucket repo ||
|BITBUCKET_PASSWORD|bitbucket token. Used for loading go modules from private bitbucket repo ||

## Owner

**Name:** Alexander Moreno

**Email:** s.a.moreno.a.s@gmail.com
