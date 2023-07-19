# mailerservice

Before getting started, please update the MAILER_SENDER_ADDRESS and MAILER_SENDER_PASSWORD fields in the dev.env file.

```
MAILER_SENDER_ADDRESS='your-mail@gmail.com'
MAILER_SENDER_PASSWORD='app password'
```

```bash
docker build -t mailerservice .
```

```bash
docker run -p 3002:3002 mailerservice
```

```bash
make mailer
```
