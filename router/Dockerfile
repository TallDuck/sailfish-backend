# Development docker image, leverages gin for automatic compilation
FROM golang:1.8-alpine
MAINTAINER <mike@tallduck.com> Mike George

ENV CODE_DIR /go/src/github.com/tallduck/sailfish-backend
ENV PROJECT router
ENV PROJECT_DIR ${CODE_DIR}/${PROJECT}

RUN mkdir -p ${CODE_DIR}{protobuf, ${PROJECT}}
WORKDIR ${PROJECT_DIR}

COPY ./router ${PROJECT_DIR}
COPY ./protobuf ${CODE_DIR}/protobuf

RUN apk add --update git

RUN go get github.com/codegangsta/gin

EXPOSE 3000

CMD gin run