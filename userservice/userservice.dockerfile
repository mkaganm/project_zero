FROM golang:1.20.5-alpine AS builder

RUN mkdir /userapp

COPY . /userapp

WORKDIR /userapp

RUN CGO_ENABLED=0 go build -o userservice ./cmd/main.go

RUN chmod +x /userapp/userservice

FROM alpine:latest

RUN mkdir /userapp

COPY --from=builder /userapp/userservice /userapp
COPY userservice.env ./userservice.env

EXPOSE 3001:3001

CMD ["/userapp/userservice"]