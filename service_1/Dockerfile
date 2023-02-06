# Dockerfile for Adder
FROM golang:alpine AS builder

WORKDIR /app

COPY . .

RUN go mod download
RUN go get google.golang.org/grpc
RUN go get github.com/hashicorp/consul/api
RUN go get github.com/grpc-ecosystem/grpc-gateway/runtime
RUN go build -o adder main.go



FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/adder /app/

CMD ["./adder"]
