FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -mod=readonly -o bin ./cmd/app/main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/bin .

EXPOSE 8080

CMD ["./bin"]