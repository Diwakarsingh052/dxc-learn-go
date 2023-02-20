package db

import (
	"context"
	"errors"
	"github.com/redis/go-redis/v9"
	"log"
	"time"
)

func ConnectToRedis(ctx context.Context) (*redis.Client, error) {
	count := 0
	for {
		if count == 5 {
			return nil, errors.New("cannot connect to redis")
		}

		rdb := redis.NewClient(&redis.Options{
			Addr: "redis-follower:6379", // this host name is from the docker file
		})

		//ping redis to make sure connection is successful
		err := rdb.Ping(ctx).Err()
		if err != nil {
			log.Println(err)
			time.Sleep(time.Second) // waiting for one second before making another connection request
			log.Println("retrying connection to redis")
			count++
			continue
		}

		//if success return the connection
		return rdb, nil

	}

}
