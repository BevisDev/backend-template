# stage 1: build app
FROM golang:1.23.4-alpine AS builder

WORKDIR /build

RUN apk add --no-cache make ca-certificates

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN make swag_install && make swag_gen \
    make wire_install && make wire_gen \
    && make build

## stage 2: run app
FROM alpine:3.21.2

WORKDIR /app

COPY ./src/resources ./src/resources

COPY --from=builder /build/gobe ./

EXPOSE 8089

ENTRYPOINT [ "./gobe" ]