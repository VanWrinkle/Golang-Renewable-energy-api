# PROG2005-assignment-2

[TOC]

## Project description



## Endpoints



## Deployment
### Configuration

### Dependencies



## Directory Structure

````
root
│   .gitlab-ci.yml
│   go.mod
│   go.sum
│   README.md
│
├───.github
│   └───workflows
│           deployment.yaml
│
├───caching
│       cache_worker.go
│       caching_structs.go
│       caching_util.go
│
├───cmd
│       server.go
│
├───consts
│       consts.go
│
├───Documentation_Internal
│       conventions.md
│       team-work.md
│       Workflow.drawio
│
├───fsutils
│       fsutils.go
│
├───handlers
│       notification.go
│       renewables.go
│       status.go
│
├───internal
│   ├───assets
│   │       codes=CHN.json
│   │       codes=FIN.json
│   │       codes=INV.json
│   │       ...
│   │       renewable-share-energy.csv
│   │
│   ├───stubbing
│   │       stubbing.go
│   │
│   └───testing
│       │   caching_test.go
│       │   fsutils_test.go
│       │   renewables_test.go
│       │   stubbing_test.go
│       │   util_test.go
│       │
│       └───internal 
│           └───assets ... copy of assets for testing
│
└───util
        util.go
````