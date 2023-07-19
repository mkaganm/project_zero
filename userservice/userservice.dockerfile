FROM golang:1.20.5-alpine as builder

RUN mkdir /userapp

COPY . /userapp

WORKDIR /userapp

RUN CGO_ENABLED=0 go build -o mailerservice ./cmd/main.go

RUN chmod +x /userapp/userservice

FROM alpine:latest

RUN mkdir /app1

COPY --from=builder /userapp/userservice /userapp
COPY userservice.env ./userservice.env

EXPOSE 3001:3001

CMD ["/userapp/userservice"]