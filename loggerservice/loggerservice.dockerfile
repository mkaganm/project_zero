FROM golang:1.20.5-alpine AS builder

RUN mkdir /loggerapp

COPY . /loggerapp

WORKDIR /loggerapp

RUN CGO_ENABLED=0 go build -o loggerservice ./cmd/main.go

RUN chmod +x /loggerapp/loggerservice

FROM alpine:latest

RUN mkdir /loggerapp

COPY --from=builder /loggerapp/loggerservice /loggerapp
COPY loggerservice.env ./loggerservice.env

EXPOSE 3003:3003

CMD ["/loggerapp/loggerservice"]