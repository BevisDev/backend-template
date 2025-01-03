FROM golang:1.23.4-alpine AS builder

WORKDIR /app

COPY . .

COPY go.mod go.sum ./

RUN go mod download && go mod verify


RUN go build -v -o res ./

