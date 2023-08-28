# syntax=docker/dockerfile:1
FROM golang:1.19 as builder

WORKDIR /app
COPY entrypoint.go .
RUN CGO_ENABLED=0 go build entrypoint.go

FROM scratch
COPY --from=builder /app/entrypoint /entrypoint 

EXPOSE 8080
ENTRYPOINT ["/entrypoint"]
