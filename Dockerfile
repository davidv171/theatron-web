FROM golang:1.8

RUN go get github.com/davidv171/theatron-web
WORKDIR /go/src/github.com/davidv171/theatron-web
RUN go build -o theatron
WORKDIR /go/bin/
CMD ["theatron-web"]