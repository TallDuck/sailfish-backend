# Development docker image, leverages reflex for automatic compilation
FROM golang:1.8-alpine
MAINTAINER <mike@tallduck.com> Mike George

# This Dockerfile will likely be used in multiple services.
# Once copied into a new service the PROJECT var must be updated
# to match the folder (package) name.
ENV PROJECT auth
ENV APP_PORT 3000
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

# Git is needed in order to fetch reflex
RUN apk add --update git

# reflex is used for auto compilation, the service will be
# recompiled on code change so long as the volume is mounted.
RUN go get github.com/cespare/reflex

EXPOSE ${APP_PORT}

CMD cd ../ && \
    reflex -r '(^|/)(helpers|protobuf|vendor|${PROJECT})/' \ 
           -r '\.go$' \
           -s -- sh -c 'go build -o ${PROJECT}/${PROJECT} ./${PROJECT}/ && ./${PROJECT}/${PROJECT}'
