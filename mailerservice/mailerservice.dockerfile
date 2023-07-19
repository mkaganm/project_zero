FROM golang:1.20.5-alpine AS builder

RUN mkdir /mailerapp

COPY . /mailerapp

WORKDIR /mailerapp

RUN CGO_ENABLED=0 go build -o mailerservice ./cmd/main.go

RUN chmod +x /mailerapp/mailerservice

FROM alpine:latest

RUN mkdir /mailerapp

COPY --from=builder /mailerapp/mailerservice /mailerapp
COPY mailerservice.env ./mailerservice.env

EXPOSE 3002:3002

CMD ["/mailerapp/mailerservice"]