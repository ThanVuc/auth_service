# build stage
FROM golang:alpine AS builder
RUN apk add --no-cache

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o auth_service ./main.go

# stage 2
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/auth_service .
COPY --from=builder /app/sql/schema /app/sql/schema
RUN chmod +x /app/auth_service
ENTRYPOINT ["./auth_service"]
