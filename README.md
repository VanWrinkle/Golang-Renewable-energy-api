# PROG2005-assignment-2

[TOC]

## Project description

### Brief overview

This is the project "Assignment 2" in the course PROG2005 at NTNU Gjøvik.

The end product of the project will be a REST web application, where users can retrieve information about the percentage of renewable energy for (most of) the countries in the world.

The application depends data about countries from the "REST Countries API"; for this assignment hosted locally at NTNU.

Furthermore, a "Renewable Energy Dataset" from https://ourworldindata.org/energy, in the form of a .csv document will be used to find renewable energy percentages.

There will be endpoints for retrieving information about renewable energy, both current and historic. Registering webhooks for getting notifications about renewable energy in countries of interest will also be implemented. All responses will be in .json format.

The service will be deployed as a Docker container in NTNUs OpenStack.



## Endpoints

The service has four main endpoints:

| Endpoint                      | Purpose/functionality                   |
| ----------------------------- | --------------------------------------- |
| /energy/v1/renewables/current | Current percentage of renwable energy   |
| /energy/v1/renewables/history | Historic percentage of renewable energy |
| /energy/v1/notifications/     | Register/view/delete webhooks           |
| /energy/v1/status/            | View status of the service              |

### Current

The "current" endpoint provides renewable energy percentage for one or more countries for the most recent year in the dataset.

General form of request:

```
Method: GET
Path: /energy/v1/renewables/current/{country?}{?neighbours=bool?}
```

Where...

* {country?} refers to a country by 3-letter name code (optional)
* {?neighbours=bool?} determines if energy data for neighbouring countries should be retrieved (optional; requires preceding country code)

Example request 1; country code:

```
/energy/v1/renewables/current/nor
```

Corresponding response:

```
[
    {
        "name": "Norway",
        "isocode": "NOR",
        "year": 2021,
        "percentage": 71.55836486816406
    }
]
```

Example request 2; country and neighbours:

```
/energy/v1/renewables/current/nor?neighbours=true
```

Corresponding response:

```
[
    {
        "name": "Norway",
        "isocode": "NOR",
        "year": 2021,
        "percentage": 71.55836486816406
    },
    {
        "name": "Finland",
        "isocode": "FIN",
        "year": 2021,
        "percentage": 34.611289978027344
    },
    {
        "name": "Sweden",
        "isocode": "SWE",
        "year": 2021,
        "percentage": 50.924007415771484
    },
    {
        "name": "Russia",
        "isocode": "RUS",
        "year": 2021,
        "percentage": 6.620289325714111
    }
]
```

Example request 3; no country code:

```
/energy/v1/renewables/current
```

Corresponding response:

```
[
    {
        "name": "Japan",
        "isocode": "JPN",
        "year": 2021,
        "percentage": 11.428995132446289
    },
    {
        "name": "Spain",
        "isocode": "ESP",
        "year": 2021,
        "percentage": 22.341663360595703
    },
    {
        "name": "Iran",
        "isocode": "IRN",
        "year": 2021,
        "percentage": 1.2903937101364136
    },
    
    .... all countries in dataset
    
]
```



### Historic

### Notifications

### Status





## Deployment
The application is designed to be deployed as a container using Docker in a linux environment.


### Configuration

### Dependencies

The following must be in place in order to deploy the application as a Docker container on a virtual machine.

#### Docker engine

Docker engine must be installed on a virtual or physical machine running Ubuntu.
Follow this instruction to get the latest Docker version:
https://docs.docker.com/engine/install/ubuntu/

We have tested and confirmed that using the "apt repository" install method described works in Ubuntu 22.04 LTS. Other methods and Ubuntu versions may work as well.

#### Docker Compose plugin

The Docker Compose plugin must be installed. Use following commands:

`sudo apt-get update
sudo apt-get install docker-compose-plugin`

Verify the installation using

`docker compose version`

#### Golang

Support for Golang must be  installed in order to compile the source code.

Download the archive: 

`wget https://go.dev/dl/go1.20.3.linux-amd64.tar.gz`

Next, remove (potetntially) existing version and extract:

`rm -rf /usr/local/go && tar -C /usr/local -xzf go1.20.3.linux-amd64.tar.gz`  

Add /usr/local/go/bin to PATH environment variables:

nano $HOME/.profile

Add the following as last line in the .profile file:

`export PATH=$PATH:/usr/local/go/bin`

Lastly, check the installation using:

`go version`



A guide is available for additional support: https://docs.docker.com/language/golang/



#### Required local files

The application relies on a Google Firestore database. To access the database, a service-account certificate must be present on the host machine. 

The certificate must have the filename "sha.json", and be located in /home/ubuntu/.secret. This directory will be mounted as a volume for the container. The certificate is project specific.

#### Network

The host machine must be connected to the Internet, and have an associated floating IP address. Furthermore, these ports must be open to traffic:

| Port | Direction | Purpose |
| ---- | --------- | ------- |
| 22   | ingress   | SSH     |
| 8080 | ingress   | http    |
|      |           |         |

### Build and deploy
The source code for the project must be downloaded to the machine used for deployment.

Clone the repository:

```
git clone git@git.gvk.idi.ntnu.no:course/prog2005/prog2005-2023-workspace/even/assignment2.git 
```

To build and deploy the application, navigate to the project directory (assignment2) and run:
```
docker compose up -d
```

The application should now be running. A message should confirm that the container is started.

The service is set up to restart automatically when the host machine is rebooted.

To manually stop the service, this command can be used:

`docker stop <name-of-container>`

Using "stop" will stop the service until service is started again, or until next reboot of host machine.

To manually start a stopped container:

`docker start <name-of-container>`

To shut down the service and remove the container, this command can be used:

`docker compose down`

The service will not start again at reboot after using the "down" command.



## Directory Structure

````
root
│   .env
│   .gitignore
│   .gitlab-ci.yml
│   compose.yml
│   Dockerfile
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
│       invocation_worker.go
│
├───cmd
│       server.go
│       sha.json
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
│   │   renewables.go
│   │   status.go
│   │
│   └───notifications
│           notification.go
│           notification_structs.go
│
├───internal
│   ├───assets
│   │       codes=CHN.json
│   │       codes=FIN.json
│   │       codes=INV.json
|   |       ...
│   │       renewable-share-energy.csv
│   │
│   ├───stubbing
│   │       stubbing.go
│   │
│   └───testing
│       │   caching_test.go
│       │   fsutils_test.go
│       │   renewables_test.go
│       │   sha.json
│       │   stubbing_test.go
│       │   util_test.go
│       │
│       └───internal
│           └───assets ... copy of assets for testing
│
└───util
        dataset.go
        util.go
````



