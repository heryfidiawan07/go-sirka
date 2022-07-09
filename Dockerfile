FROM golang:1.14.9-alpine
RUN mkdir /bin
ADD go.mod go.sum server.go /bin/
WORKDIR /bin
RUN go build