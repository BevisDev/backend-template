# stage 1: build app
FROM golang:1.23.4-alpine AS builder

WORKDIR /build

COPY go.mod go.sum ./

RUN go mod download

COPY . .

EXPOSE 8089

RUN go build -v -o gobe ./src/main

## stage 2: run app
FROM scratch

COPY ./src/resources /src/resources

COPY --from=builder /build/gobe ./

ENTRYPOINT [ "./gobe" ]