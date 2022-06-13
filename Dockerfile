# syntax=docker/dockerfile:1

FROM golang:1.18-alpine

WORKDIR /go/src

COPY ./controller ./controller
COPY ./domain ./domain
COPY ./interfaces ./interfaces
COPY ./usecases ./usecases

COPY go.mod ./
COPY go.sum ./
COPY main.go ./

RUN go mod download

# Build
RUN go build -o main



