package main

import (
	"client/proto"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

func main() {
	//dialing connection
	conn, err := grpc.Dial(":5000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	//registering the client
	c := proto.NewLogServiceClient(conn)

	l := proto.Log{
		Name: "testing log",
		Data: "logging an error",
	}

	//constructing the logRequest
	lr := &proto.WriteLogRequest{LogEntry: &l}

	//setting timeouts
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	//calling of the remote server
	res, err := c.WriteLog(ctx, lr)

	if err != nil {
		log.Fatalln("something went wrong ", err)
	}

	log.Println("we did some gRPC stuff", res)
}
