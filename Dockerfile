FROM golang:1.26-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o token-monitoring cmd/server/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/token-monitoring .

EXPOSE 8081

CMD ["./token-monitoring"]
