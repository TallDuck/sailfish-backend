# Development docker image, leverages gin for automatic compilation
FROM golang:1.8-alpine
MAINTAINER <mike@tallduck.com> Mike George

# This Dockerfile will likely be used in multiple services.
# Once copied into a new service the PROJECT var must be updated
# to match the folder (package) name.
ENV PROJECT router
ENV PORT 8080
ENV CODE_DIR /go/src/github.com/tallduck/sailfish-backend
ENV PROJECT_DIR ${CODE_DIR}/${PROJECT}

RUN mkdir -p ${CODE_DIR}{protobuf, ${PROJECT}}
WORKDIR ${PROJECT_DIR}

COPY ./${PROJECT} ${PROJECT_DIR}

# This service is dependent on the protobuf package, it
# gets copied in so it doesn't have to be downloaded
COPY ./protobuf ${CODE_DIR}/protobuf

# Git is needed in order to fetch gin
RUN apk add --update git

# Gin is used for auto compilation, the service will be
# recompiled on code change so long as the volume is mounted.
RUN go get github.com/codegangsta/gin

EXPOSE 3000

CMD gin run