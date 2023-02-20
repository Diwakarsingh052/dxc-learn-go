package database

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/redis/go-redis/v9"
	"log"
	"time"
)

const (
	host     = "postgres-article" // this host name is from the docker file
	port     = 5432
	user     = "postgres"
	password = "password"
	dbname   = "users"
)

func ConnectToPostgres() (*sql.DB, error) {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	return sql.Open("postgres", psqlInfo)

}

func ConnectToRedis() (*redis.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	count := 0
	//connecting to redis and retrying 5 times in case redis is not ready to accept connections
	for {
		if count == 5 {
			return nil, errors.New("cannot connect to redis")
		}

		rdb := redis.NewClient(&redis.Options{
			Addr: "redis-event:6379", // this host name is from the docker file
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
