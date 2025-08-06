FROM golang:1.23.1-alpine AS builder

RUN apk add --no-cache git ca-certificates

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
    -ldflags="-s -w" \
    -o crypto-price-service \
    cmd/service/main.go

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
    -ldflags="-s -w" \
    -o migrate \
    cmd/migrate/main.go

FROM alpine:latest AS final

WORKDIR /root/

COPY --from=builder /app/crypto-price-service .
COPY --from=builder /app/migrations migrations
COPY --from=builder /app/config config
COPY --from=builder /app/migrate .
COPY --from=builder /app/entrypoint.sh .

# Сделай скрипт исполняемым
RUN chmod +x ./entrypoint.sh

EXPOSE 8080

CMD ["./entrypoint.sh"]