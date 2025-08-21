FROM golang:1.24.4-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git make gcc musl-dev

COPY  go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o gateway ./cmd/main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/gateway .

CMD ["./gateway"]