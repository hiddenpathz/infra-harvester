FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY main.go .
RUN go build -o tcp-proxy main.go

FROM alpine:latest
RUN apk add --no-cache ca-certificates
COPY --from=builder /app/tcp-proxy /usr/local/bin/tcp-proxy
ENTRYPOINT ["/usr/local/bin/tcp-proxy"]