package main

import (
	"context"
	"database/sql"
	"google.golang.org/grpc"
	"log"
	"net"
	"server/proto"
)

type LogServer struct {
	//add any deps for your gRPC server
	db *sql.DB
	proto.UnimplementedLogServiceServer
}

const gRPCPort = "5000"

func main() {
	// creating a tcp connection that would be used by our server to open up connection
	lis, err := net.Listen("tcp", ":5000")
	if err != nil {
		log.Fatalln(err)
	}
	// NewServer creates a gRPC server which has no service registered and has not
	// started to accept requests yet.
	s := grpc.NewServer()
	db, _ := sql.Open("", "")

	//registering the log server and passing any required deps
	proto.RegisterLogServiceServer(s, &LogServer{db: db})

	log.Printf("gRPC Server started on port %s", gRPCPort)
	err = s.Serve(lis)
	if err != nil {
		log.Fatalln(err)
	}
}

func (l *LogServer) WriteLog(ctx context.Context, req *proto.WriteLogRequest) (*proto.WriteLogResponse, error) {
	//l.db.Query()

	input := req.GetLogEntry()

	// do your business logic
	log.Println(input, "yes we have used gRPC to log a message in our terminal")

	res := proto.WriteLogResponse{Result: "all good here, logging is done"}

	return &res, nil

}
