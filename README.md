# Sailfish Backend


[![Codeship Status for tallduck/sailfish-backend](https://img.shields.io/codeship/f0dc8440-0607-0135-edfd-52b395dcacd9/master.svg)](https://app.codeship.com/projects/213698)
[![Go Report Card](https://goreportcard.com/badge/github.com/tallduck/sailfish-backend)](https://goreportcard.com/report/github.com/tallduck/sailfish-backend)

Current Status: WIP (pre-alpha)

## Architecture

The Sailfish Backend will be build with Go, leveraging microservices that communicate via [gRPC](http://www.grpc.io/) with [Protocol Buffers](https://developers.google.com/protocol-buffers/) as the Request/Response definition.

The router package is the API entrypoint, as the name implys, it will route each request to the proper service for processing. This includes marshaling JSON to the proper Protobuf.

## Intentions

Sailfish Backend is the api for a commercial product idea that is being developed in the open. The current plans will keep a permissive license on this codebase even after it is launched.

The business plan for this software includes a Paid / Hosted version as well as a free Open Source Edition.

Support and Installation contracts for the Open Source Edition will also be available by request. The terms of which will be determined at the time of request.