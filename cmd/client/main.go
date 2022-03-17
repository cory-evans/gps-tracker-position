package main

import (
	"context"
	"log"

	positionv1 "github.com/cory-evans/gps-tracker-position/pkg/position/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln(err)
	}

	defer conn.Close()

	client := positionv1.NewPositionServiceClient(conn)

	resp, err := client.GetPosition(context.Background(), &positionv1.GetPositionRequest{})

	if err != nil {
		log.Fatalln(err)
	}

	log.Println(resp.Position)
}
