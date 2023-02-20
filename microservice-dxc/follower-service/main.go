package main

import (
	"context"
	"follower-service/db"
	"follower-service/proto"
	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

type FollowServer struct {
	//UnimplementedFollowUserServiceServer must be embedded to have forward compatible implementations.
	proto.UnimplementedFollowUserServiceServer

	//add deps for the gRPC server
	*redis.Client
}

func main() {
	gRPCServer()
}

func gRPCServer() {

	lis, err := net.Listen("tcp", ":5000")
	if err != nil {
		log.Fatalln(err)
	}

	//NewServer creates a gRPC server which has no service registered and has not started to accept requests yet
	s := grpc.NewServer()

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	//connecting to redis with timeout of 20 sec
	r, err := db.ConnectToRedis(ctx)

	if err != nil {
		log.Fatalln("redis connection failed")
	}

	// registering the rpc server
	proto.RegisterFollowUserServiceServer(s, &FollowServer{Client: r})

	err = s.Serve(lis)
	if err != nil {
		log.Fatalln(err)
	}

}

func (f *FollowServer) FollowUser(ctx context.Context, req *proto.FollowUserRequest) (*proto.FollowUserResponse, error) {

	// Get method retrieve the data from the request struct
	u := req.GetUser()

	//	 john want to follow bruce
	// following john -> bruce
	// followers bruce -> john

	// creating a key name for redis // in this key we will store user is following whom
	followingKey := "user:following:" + u.UserEmail // user:following:john@email.com

	// saving user is  following whom
	err := f.Client.SAdd(ctx, followingKey, u.TargetUserEmail).Err()

	if err != nil {
		log.Println("failed to add follower in redis", err)
		return nil, err
	}

	// creating key name to store the list of updated followers of the target user
	followerKey := "user:followers:" + u.TargetUserEmail
	// saving followers of the user
	//e.g. when bob follows john then john gets one follower, we should add bob in the follower list of john

	err = f.Client.SAdd(ctx, followerKey, u.UserEmail).Err()
	if err != nil {
		log.Println("failed to add follower in redis", err)
		return nil, err
	}

	//constructing the response
	res := &proto.FollowUserResponse{Result: "follow user success"}

	return res, nil

}

func (f *FollowServer) ListFollowing(ctx context.Context, req *proto.ListFollowingRequest) (*proto.ListFollowingResponse, error) {
	r := req.GetEmail()

	//creating a key to get a list of whom this user is following
	key := "user:following:" + r.Email

	//it will return list of following
	list, err := f.Client.SMembers(ctx, key).Result()

	if err != nil {
		log.Println(err)
		return nil, err
	}

	//constructing the list of following to send as resp
	res := &proto.ListFollowingResponse{FollowingList: list}

	return res, nil

}

func (f *FollowServer) ListFollowers(ctx context.Context, req *proto.ListFollowersRequest) (*proto.ListFollowersResponse, error) {

	// Get method retrieve the data from the request struct
	r := req.GetEmail()

	//constructing the key for the current logged-in user
	key := "user:followers:" + r.Email

	//fetching list of the followers of the current user from redis
	list, err := f.Client.SMembers(ctx, key).Result()

	if err != nil {
		log.Println(err)
		return nil, err
	}

	//constructing the response
	res := &proto.ListFollowersResponse{Followers: list}
	return res, nil
}
