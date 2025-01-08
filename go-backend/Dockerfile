# stage 1: build app

FROM golang:1.23.4-alpine AS builder

WORKDIR /build

COPY . .

COPY go.mod go.sum ./

RUN go mod donwload

RUN go build -v -o bevisdev ./

# stage 2: run app
FROM scratch

COPY . .

COPY --from=builder /build/bevisdev ./


