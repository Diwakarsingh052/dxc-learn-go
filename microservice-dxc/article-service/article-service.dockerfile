# base go image
FROM golang:1.19-alpine as builder

RUN mkdir /app

COPY . /app

WORKDIR /app

RUN go build -o article-service .

RUN chmod +x /app/article-service

# build a tiny docker image
FROM alpine:latest

RUN mkdir /app

COPY --from=builder /app/article-service /app

CMD [ "/app/article-service" ]