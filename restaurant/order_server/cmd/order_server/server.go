package main

import (
	"context"
	"log"
	"net"
	"os"
	"time"

	"github.com/miladouski/golang-training-restaurant-grpc/restaurant/db"
	"github.com/miladouski/golang-training-restaurant-grpc/restaurant/order_server/pkg/api"
	pb "github.com/miladouski/golang-training-restaurant-grpc/restaurant/proto/proto"

	"google.golang.org/grpc"
	"gorm.io/gorm"
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
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	ctx, cancel = context.WithTimeout(ctx, time.Second*30)
	defer cancel()

	conn, err := connectToDbWithTimeout(ctx)
	if err != nil {
		log.Fatalf("can't connect to database, error: %v", err)
	}

	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}

	server := grpc.NewServer()
	pb.RegisterOrderServiceServer(server, api.NewOrderServer(conn))
	err = server.Serve(listener)
	if err != nil {
		log.Fatal(err)
	}
}

func connectToDbWithTimeout(ctx context.Context) (*gorm.DB, error) {
	for {
		time.Sleep(2 * time.Second)
		conn, err := db.GetConnection(host, dbport, user, dbname, password, sslmode)
		if err == nil {
			return conn, nil
		}
		select {
		case <-ctx.Done():
			return nil, err
		default:
			continue
		}
	}
}
