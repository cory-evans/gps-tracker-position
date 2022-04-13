package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/cory-evans/gps-tracker-authentication/pkg/auth"
	"github.com/cory-evans/gps-tracker-position/internal/database"
	"github.com/cory-evans/gps-tracker-position/internal/service"
	"github.com/cory-evans/gps-tracker-position/pkg/position"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

func myAuthFunc(ctx context.Context) (context.Context, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, fmt.Errorf("could not get metadata")
	}

	log.Println(md)

	return ctx, nil
}

func main() {
	listen, err := net.Listen("tcp", ":"+os.Getenv("GRPC_PORT"))

	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_auth.UnaryServerInterceptor(myAuthFunc)),
		grpc.StreamInterceptor(grpc_auth.StreamServerInterceptor(myAuthFunc)),
	)

	authConn, err := grpc.Dial(os.Getenv("AUTH_SERVICE"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	authServiceClient := auth.NewAuthServiceClient(authConn)

	mongoCtx := context.Background()
	db, err := database.NewDatabaseClient(os.Getenv("MONGO_URI"), mongoCtx)
	if err != nil {
		log.Fatalln(err)
	}

	positionDB := db.Database("position")

	position.RegisterPositionServiceServer(grpcServer, &service.PositionService{
		AuthServiceClient: authServiceClient,
		DB:                positionDB,
	})

	grpcServer.Serve(listen)
}
