FROM golang:1.20.5-alpine AS builder

RUN mkdir /configuratorapp

COPY . /configuratorapp

WORKDIR /configuratorapp

RUN CGO_ENABLED=0 go build -o configurator .main.go

RUN chmod +x /configuratorapp/configurator

FROM alpine:latest

RUN mkdir /app

COPY --from=builder /configuratorapp/configurator /configuratorapp/configurator

CMD ["/configuratorapp/configurator"]