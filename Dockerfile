##
## Build
##
FROM golang:alpine3.10 AS build

WORKDIR /app
RUN apk add --no-cache git make bash
COPY go.mod ./
RUN go mod download

COPY * ./
CMD ["make", "build_run"]