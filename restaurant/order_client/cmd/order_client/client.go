package main

import (
	"context"
	"log"
	"net/http"
	"os"

	pb "github.com/miladouski/golang-training-restaurant-grpc/restaurant/proto/proto"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
)

var (
	server = os.Getenv("SERVER")
	port   = os.Getenv("PORT")
)

func init() {
	if server == "" {
		server = "localhost:8282"
	}
	if port == "" {
		port = "localhost:8383"
	}
}

func main() {
	conn, err := grpc.Dial(server, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	grpcMux := runtime.NewServeMux()
	err = pb.RegisterOrderServiceHandler(context.Background(), grpcMux, conn)
	if err != nil {
		log.Fatal(err)
	}

	log.Fatal(http.ListenAndServe(port, grpcMux))
}
