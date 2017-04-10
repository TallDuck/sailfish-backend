# Development docker image, leverages gin for automatic compilation
FROM golang:1.8-alpine
MAINTAINER <mike@tallduck.com> Mike George

ENV CODE_DIR /go/src/github.com/tallduck/sailfish-backend/router

RUN mkdir -p ${CODE_DIR}
WORKDIR ${CODE_DIR}

COPY . ${CODE_DIR}

RUN apk add --update git

RUN go get github.com/codegangsta/gin
RUN go-wrapper download
RUN go-wrapper install

ENV PORT 8080

EXPOSE 3000

CMD gin run