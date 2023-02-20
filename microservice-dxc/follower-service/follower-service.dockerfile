# base go image
FROM golang:1.19-alpine as builder

RUN mkdir /app

COPY . /app

WORKDIR /app

RUN go build -o follower-service .

RUN chmod +x /app/follower-service

# build a tiny docker image
FROM alpine:latest

RUN mkdir /app

COPY --from=builder /app/follower-service /app

CMD [ "/app/follower-service" ]