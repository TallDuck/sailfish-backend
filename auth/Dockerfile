# Development docker image, leverages CompileDaemon for automatic compilation
FROM golang:1.8-alpine
MAINTAINER <mike@tallduck.com> Mike George

# This Dockerfile will likely be used in multiple services.
# Once copied into a new service the PROJECT var must be updated
# to match the folder (package) name.
ENV PROJECT auth
ENV APP_PORT 8080
ENV CODE_DIR /go/src/github.com/tallduck/sailfish-backend
ENV PROJECT_DIR ${CODE_DIR}/${PROJECT}

RUN mkdir -p ${CODE_DIR}{protobuf, helpers, ${PROJECT}}
WORKDIR ${PROJECT_DIR}

COPY ./${PROJECT} ${PROJECT_DIR}

# This service is dependent on the protobuf and helpers packages,
# they get copied in so they don't have to be downloaded
COPY ./protobuf ${CODE_DIR}/protobuf
COPY ./helpers ${CODE_DIR}/helpers

# Copy the vendor directory containing dependencies for all services
COPY ./vendor ${CODE_DIR}/vendor

# Git is needed in order to fetch CompileDaemon
RUN apk add --update git

# CompileDaemon is used for auto compilation, the service will be
# recompiled on code change so long as the volume is mounted.
RUN go get github.com/githubnemo/CompileDaemon

EXPOSE ${APP_PORT}

CMD CompileDaemon -command="./${PROJECT}"