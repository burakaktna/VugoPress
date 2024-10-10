FROM golang:1.20-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o main ./cmd/vugopress

FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/main /app/main
COPY --from=builder /app/.env /app/.env

EXPOSE 3000
CMD ["/app/main"]
