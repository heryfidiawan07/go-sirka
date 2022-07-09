FROM golang:1.14.9-alpine
RUN mkdir /bin/app
ADD go.mod go.sum server.go /bin/app/
WORKDIR /bin/app
RUN go build