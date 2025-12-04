# syntax=docker/dockerfile:1
FROM golang:1.21-alpine AS builder

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o go-service

FROM alpine:3.19

RUN apk add --no-cache ffmpeg ca-certificates

WORKDIR /srv
COPY --from=builder /app/go-service ./go-service

EXPOSE 8080

ENV PORT=8080
CMD ["./go-service"]
