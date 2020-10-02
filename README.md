# Gin Project Structure
Golang APIs for Example service.

this project contains backend APIs and logic for Example Service powered by Go,Gin,Redis,MinIO,CouchBase and Dgraph
> Application runs on Go 1.14 And redis,couchbase,MinIO,dgraph and scylladb.
  
## Table of Contents  
  
- [Project Structure](#Project Structure)  
- [Install application](#Install application)  
- [Configuration](#configuration)  
- [Features](#Features)  
- [Services and Database dependencies](#Services and Database dependencies)  
- [Testing](#testing)  
  


## Project Structure 
```
.
├── config                      # all config files is here
    ├── file                    # develop yml config file is here
    └── test_conf               # unit test yml config file is here
├── console                     # console commands dicretory
    └── cmd                     # all console commands is here 
├── constants                   # some static values is here
├── docs                        # any documentation is here swagger,flowchat,gif , ....
├── logic                       # business logic layer directories. all server business logics is here
    └── tests                   # we create logics unit test in here
├── model                       # we save project struct/models in here
    └── proto                   # also if we need proto files we save it here
├── repository                  # repository layer directories. all server repository logics is here
    ├── rediskeys               # redis keys names is here
    ├── couchbaseQuery          # we save static quere is here
    └── tests                   # we create repository unit test in here.
├── router                      # router layer directories. all server router logics is here 
    ├── grpc                    # a directory for grpc routers
    └── http                    # a directory for http routers
        └── tests               # we create controllers unit test in here.
├── services                    # request and connections to therd party servers is here
├── utils                       # some local utils functions saved in here
├── .gitignore                  # realy this need explain ?!?!?
├── .gitlab-ci.yml              # gitlab ci stages and commands
├── go.mod                      # go modules file
├── go.sum                      # go modules file
├── main.go                     # project started in this file
└── Readme.md                   # you read me here :)
```

## Install application
###  Get application to run locally for development or deployment

```bash
https://github.com/RezaOptic/gin-project-structure
cd gin-project-structure
```
### installation dependencies
To installation dependencies use ```go mod vendor``` command.

## Configuration  
by default app server in running read config file from `./file/cofigs.yml` if you want to change some service or database address or ports must edit config file
### Setup your own config
to setup config you can set your settings into yml file.

### Run Application
To run application on development config use
```go run main.go``` 

## Features
### consoles
some time we need some worker to do some jobs periodically we use `cobra` to write console commands to do some jobs.
### API's
#### grpc
we have multiple grpc methods :
#### http
we have multiple http route :
- `/v1/` `POST`‍: api for ...
- `/v1/` `POST`: api for ...
- `/v1/` `POST`: api for ...


## Services and Database dependencies
this server need `redis`,`couchbase`,`MinIO`, `dgraph` and `scylladb` servers to start running.

## Testing  
TODO: Add dependency injection structure
TODO: Additional instructions for testing the application.
TODO: maybe better configs structure