# Dockerfile for Printer
FROM golang:alpine AS builder

WORKDIR /app

COPY . .

RUN go mod download
RUN go get github.com/hashicorp/consul/api
RUN go get google.golang.org/grpc

RUN go build -o printer main.go


FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/printer /app/

CMD ["./printer"]
