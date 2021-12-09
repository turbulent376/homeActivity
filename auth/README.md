# Auth service

## Description
**auth** service is responsible for user authentication

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
|SPACE_USER|space account. Used for loading go modules from private bitbucket repo ||
|SPACE_PASSWORD|space token. Used for loading go modules from private bitbucket repo ||
|AUTH_KEY_RS256|RSA private key to signing JWT token ||

## Owner
**Name:** izolot

**Email:** ivvzolotarev@gmail.com
