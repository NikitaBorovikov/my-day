FROM golang:1.23-alpine

WORKDIR /app

COPY go.mod .
COPY go.sum .
COPY cmd/app/main.go .
COPY . .

RUN go build -o bin .

ENTRYPOINT [ "/app/bin" ]