package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"
	"notificaiton-service/database"
	"notificaiton-service/proto"
	"time"
)

func main() {
	log.Println("connecting to redis....")
	var rdb *redis.Client

	rdb, err := database.ConnectToRedis()

	if err != nil {
		log.Println("all attempts failed , cannot connect to redis", err)
		return
	}

	log.Println("redis connected")

	http.HandleFunc("/ping", ping)

	go getArticleEvent(rdb)

	http.ListenAndServe(":8080", nil)
}

func ping(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "notification service is running")
}

func getArticleEvent(r *redis.Client) {
	// for loop will make sure we keep listening for new events
	for {
		// xread in redis allows us to listen for new messages that are not processed
		val := redis.XReadArgs{
			Streams: []string{"article:add:events", "$"}, // article:add:event is key name, $ is used to fetch msg after the last delivered id
			Count:   1,                                   // return 1 event at a time
			Block:   0,                                   // it means wait until there is no new event // unlimited time
		}
		//exec the xread command in redis and waiting for the result
		res, err := r.XRead(context.Background(), &val).Result()

		if err != nil {
			log.Println(err)
		}

		//when new event is read from redis, we will send the notification to all the followers that a new post is out
		go sendNotification(res) // running it as go routine, so we can move forward to listen for the new messages

	}

}

// sendNotification sends a notification to all the followers of the publisher
func sendNotification(res []redis.XStream) {

	//ranging over xread result
	for _, r := range res {
		fmt.Println(r.Stream, "key name")

		for _, msg := range r.Messages {
			fmt.Println(msg.ID, "stream id")
			fmt.Println(msg.Values, "printing values of the map")

			//taking the email out of the article publisher from the event
			v, ok := msg.Values["email"]
			if !ok {
				log.Println("email not found, can't send a notification")
				continue
			}
			//making sure email is of correct type using type assertion
			email, ok := v.(string)
			if !ok {
				log.Println("email type can't be identified, can't send a notification")
				continue
			}

			title, ok := msg.Values["title"]
			if !ok {
				log.Println("title not found, can't send a notification")
				continue
			}

			//fetching followers list of the publisher
			followersList, err := fetchFollowers(email)

			if err != nil {
				log.Println(err)
				continue
			}

			//send an email notification to all the followers in a goroutine
			go sendEmail(followersList, title)

		}

	}

}

// fetchFollowers functions fetches the follower of the publisher of the article
// we will contact the follower-service to return a list of followers
func fetchFollowers(email string) ([]string, error) {

	conn, err := grpc.Dial("follower-service:5000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("can't reach follower-service %w", err)

	}

	defer conn.Close()

	//creating a gRPC client
	c := proto.NewFollowUserServiceClient(conn)

	f := &proto.ListFollowersRequest{
		Email: &proto.Follower{Email: email},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// calling the ListFollowers rpc
	res, err := c.ListFollowers(ctx, f)
	if err != nil {
		return nil, fmt.Errorf("fetching follower list failed %w", err)

	}

	return res.Followers, nil

}

func sendEmail(list []string, title interface{}) {
	for _, email := range list {
		fmt.Println("sending email", email)
		fmt.Println("article title", title)
	}

}
