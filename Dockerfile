##
## Build
##
FROM golang:alpine3.10 AS builder

ENV GO111MODULE=on
WORKDIR /app
RUN apk add --no-cache git make bash
COPY go.mod ./
RUN go mod download

COPY . .
RUN ["make", "build"]
CMD ["./Go_Project-linux", "file.txt"]