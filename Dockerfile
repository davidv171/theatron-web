FROM golang:1.8

WORKDIR /go/src/theatron

RUN go build -o theatron