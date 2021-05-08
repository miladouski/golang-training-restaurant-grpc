package main

import (
	"golang-training-restaurant-grpc/restaurant/db"
	"golang-training-restaurant-grpc/restaurant/order_server/pkg/api"
	"log"
	"net"
	"os"

	pb "golang-training-restaurant-grpc/restaurant/proto/proto"

	"google.golang.org/grpc"
)

var (
	host     = os.Getenv("DB_USERS_HOST")
	dbport   = os.Getenv("DB_USERS_PORT")
	user     = os.Getenv("DB_USERS_USER")
	dbname   = os.Getenv("DB_USERS_DBNAME")
	password = os.Getenv("DB_USERS_PASSWORD")
	sslmode  = os.Getenv("DB_USERS_SSL")
)

func init() {
	if host == "" {
		host = "localhost"
	}
	if dbport == "" {
		dbport = "5432"
	}
	if user == "" {
		user = "postgres"
	}
	if dbname == "" {
		dbname = "restaurant"
	}
	if password == "" {
		password = "1111"
	}
	if sslmode == "" {
		sslmode = "disable"
	}
}

const port = ":8282"

func main() {
	conn, err := db.GetConnection(host, dbport, user, dbname, password, sslmode)
	if err != nil {
		log.Fatalf("can't connect to database, error: %v", err)
	}
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}

	server := grpc.NewServer()

	pb.RegisterOrderServiceServer(server, api.NewOrderServer(conn))

	if err = server.Serve(listener); err != nil {
		log.Fatal(err)
	}
}
