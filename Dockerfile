FROM golang:1.8

RUN go get github.com/davidv171/theatron-web
WORKDIR /go/src/theatron

RUN go build -o theatron