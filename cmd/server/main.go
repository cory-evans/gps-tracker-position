package main

import (
	"net"

	"github.com/cory-evans/gps-tracker-position/internal/service"
	"github.com/cory-evans/gps-tracker-position/pkg/position"
	"google.golang.org/grpc"
)

func main() {
	listen, err := net.Listen("tcp", ":8080")

	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()

	position.RegisterPositionServiceServer(grpcServer, &service.PositionService{})

	grpcServer.Serve(listen)
}
