# base go image
FROM golang:1.19-alpine as builder

RUN mkdir /app


COPY . /app

WORKDIR /app

RUN go build -o graphql-service .

RUN chmod +x /app/graphql-service

# build a tiny docker image
FROM alpine:latest

RUN mkdir /app


COPY --from=builder /app/graphql-service /app

CMD [ "/app/graphql-service" ]

