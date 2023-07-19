FROM golang:1.20.5-alpine AS builder

RUN mkdir /cronitorapp

COPY . /cronitorapp

WORKDIR /cronitorapp

RUN CGO_ENABLED=0 go build -o cronitor ./cmd/main.go

RUN chmod +x /cronitorapp/cronitor

FROM alpine:latest

RUN mkdir /app

COPY --from=builder /cronitorapp/cronitor /cronitorapp/cronitor
COPY cronitor.env ./cronitor.env

CMD ["/cronitorapp/cronitor"]