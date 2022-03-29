package main

import (
	"net"

	servicev1 "github.com/cory-evans/gps-tracker-position/internal/service/v1"
	positionv1 "github.com/cory-evans/gps-tracker-position/pkg/position/v1"
	"google.golang.org/grpc"
)

func main() {
	listen, err := net.Listen("tcp", ":8080")

	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()

	positionv1.RegisterPositionServiceServer(grpcServer, &servicev1.PositionService{})

	grpcServer.Serve(listen)
}
