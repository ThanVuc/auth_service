# build stage
FROM golang:alpine AS builder
RUN apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o auth_service ./cmd/server

# stage 2
FROM scratch
COPY --from=builder /app/auth_service /app/auth_service
COPY config/dev.yaml /app/config/dev.yaml
WORKDIR /app
ENTRYPOINT ["./auth_service"]