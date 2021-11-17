FROM golang:1.17.2-alpine3.14

WORKDIR /workspace

COPY go.mod go.mod
COPY go.sum go.sum

RUN go install github.com/golang/mock/mockgen
RUN rm -rf *
