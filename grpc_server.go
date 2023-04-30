package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/rlvgl/2pchess/rpc"
	"google.golang.org/grpc"
)

type grpcGameServer struct {
	moveLog []string
	rpc.UnimplementedGameServer
}

func NewGRPCGameServer() *grpcGameServer {
	return &grpcGameServer{}
}

func (g *grpcGameServer) Move(context context.Context, moveData *rpc.MoveData) (*rpc.MoveData, error) {
	log.Println(moveData)
	return &rpc.MoveData{Move: "e7e5"}, nil
}

func (g *grpcGameServer) Run() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer()
	rpc.RegisterGameServer(grpcServer, NewGRPCGameServer())

	fmt.Println("Starting GRPC Server")
	grpcServer.Serve(lis)
}
