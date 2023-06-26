# syntax=docker/dockerfile:1
FROM golang:1.19

WORKDIR /app

COPY entrypoint.go .
RUN go build entrypoint.go
EXPOSE 8080

ENTRYPOINT ["./entrypoint"]
