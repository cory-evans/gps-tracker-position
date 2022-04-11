package main

import (
	"context"
	"log"

	"github.com/cory-evans/gps-tracker-position/pkg/position"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln(err)
	}

	defer conn.Close()

	client := position.NewPositionServiceClient(conn)

	resp, err := client.GetPosition(context.Background(), &position.GetPositionRequest{})

	if err != nil {
		log.Fatalln(err)
	}

	log.Println(resp.Position)
}
