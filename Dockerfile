FROM golang:1.24.5-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /out/notifier-server ./cmd

FROM alpine:latest
WORKDIR /app
COPY --from=builder /out/notifier-server .

RUN touch .env

ENTRYPOINT ["./notifier-server"]