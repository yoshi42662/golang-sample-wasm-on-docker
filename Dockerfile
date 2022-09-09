FROM golang:1.19

WORKDIR /go/src/golang

COPY . /go/src/golang

RUN go mod download & go mod vendor & go mod tidy
