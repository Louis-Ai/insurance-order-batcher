FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download 
COPY . .
RUN go build -o /bin/order-batcher cmd/main.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /bin/order-batcher /app/order-batcher
COPY .env .

CMD ["/app/order-batcher"]