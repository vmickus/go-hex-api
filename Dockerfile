FROM golang:lastest as builder


WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build cmd/api/main.go

