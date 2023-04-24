# syntax=docker/dockerfile:1

## Build the application from source code:

# using Golang base image:
FROM golang:1.20 AS build-stage

# setting "root" folder of project in image;
# subsequent dirs are relative to this:
WORKDIR /go/src/app

# download dependencies:
COPY go.mod go.sum ./
RUN go mod download

# copy source code into image for compiling:
COPY caching/ ./caching
COPY cmd/ ./cmd
COPY consts/ ./consts
COPY fsutils/ ./fsutils
COPY handlers/ ./handlers
COPY internal/ ./internal
COPY util/ ./util

# switch to cmd (location of main):
WORKDIR cmd/

# compile for linux:
# RUN CGO_ENABLED=0 GOOS=linux go build -o /energy ./
# using flags from example 16:
RUN CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -o /energy ./


## Create smaller image with binaries (no source code) for deployment:

FROM gcr.io/distroless/base-debian11 AS build-release-stage

# set working dir to root of container:
WORKDIR /

#copy executable:
COPY --from=build-stage /energy /energy

# expose port (not strictly necessary):
EXPOSE 8080

#run container as non-root user:
USER nonroot:nonroot

# run container as executable:
ENTRYPOINT ["/energy"]