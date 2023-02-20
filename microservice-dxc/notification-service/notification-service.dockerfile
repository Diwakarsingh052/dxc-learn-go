# base go image
FROM golang:1.19-alpine as builder

RUN mkdir /app

COPY . /app

WORKDIR /app

RUN go build -o notification-service .

RUN chmod +x /app/notification-service

# build a tiny docker image
FROM alpine:latest

RUN mkdir /app

COPY --from=builder /app/notification-service /app

CMD [ "/app/notification-service" ]