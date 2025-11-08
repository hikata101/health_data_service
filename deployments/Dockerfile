
FROM golang:1.24 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -o health-data-service ./cmd/health-data-service/main.go


FROM alpine:3.20

RUN apk add --no-cache ca-certificates

WORKDIR /app

COPY --from=builder /app/health-data-service .

EXPOSE 50051

ENV health_LOG_LEVEL=info

ENTRYPOINT ["./health-data-service"]
