package main

import (
	"context"
	pb "golang-training-restaurant-grpc/restaurant/proto/proto"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:8282", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	grpcMux := runtime.NewServeMux()
	err = pb.RegisterOrderServiceHandler(context.Background(), grpcMux, conn)
	if err != nil {
		log.Fatal(err)
	}

	log.Fatal(http.ListenAndServe("localhost:8383", grpcMux))
}
