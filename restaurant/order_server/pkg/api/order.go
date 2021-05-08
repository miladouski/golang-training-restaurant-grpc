package api

import (
	"context"
	"gorm.io/gorm"
	"log"
	"strconv"
	"time"

	"golang-training-restaurant-grpc/restaurant/order_server/pkg/data"
	pb "golang-training-restaurant-grpc/restaurant/proto/proto"
)

type orderServer struct {
	data *data.OrderData
}

func NewOrderServer(db *gorm.DB) *orderServer {
	return &orderServer{data: data.NewOrderData(db)}
}

func (o orderServer) ReadAll(ctx context.Context, request *pb.ReadAllRequest) (*pb.ReadAllResponse, error) {
	orders, err := o.data.ReadAll()
	if err != nil {
		log.Println("got an error when tried to read orders")
	}
	var pbOrders []*pb.Order
	for i := 0; i < len(orders); i++ {
		order := &pb.Order{
			Id:          orders[i].Id,
			Date:        orders[i].Date.String(),
			TableNumber: orders[i].TableNumber,
			WaiterId:    orders[i].WaiterId,
			Price:       orders[i].Price,
			Payment:     orders[i].Payment,
		}
		pbOrders = append(pbOrders, order)
	}
	return &pb.ReadAllResponse{Order: pbOrders}, nil
}

func (o orderServer) Read(ctx context.Context, request *pb.ReadRequest) (*pb.ReadResponse, error) {
	order, err := o.data.Read(request.Id)
	if err != nil {
		log.Println("got an error when tried to read orders")
	}
	return &pb.ReadResponse{Order: &pb.Order{
		Id:          order.Id,
		Date:        order.Date.String(),
		TableNumber: order.TableNumber,
		WaiterId:    order.WaiterId,
		Price:       order.Price,
		Payment:     order.Payment,
	}}, nil
}

func (o orderServer) Create(ctx context.Context, request *pb.CreateRequest) (*pb.CreateResponse, error) {
	date, err := time.Parse("2006-Jan-02", request.Order.Date)
	if err != nil {
		log.Println("wrong date format")
	}
	order := data.Order{
		Id:          request.Order.Id,
		Date:        date,
		TableNumber: request.Order.TableNumber,
		WaiterId:    request.Order.WaiterId,
		Price:       request.Order.Price,
		Payment:     request.Order.Payment,
	}
	err = o.data.Create(order)
	if err != nil {
		log.Fatal("got an error when tried to create order")
	}
	return &pb.CreateResponse{}, nil
}

func (o orderServer) Update(ctx context.Context, request *pb.UpdateRequest) (*pb.UpdateResponse, error) {
	payment, _ := strconv.ParseBool(request.Payment)
	err := o.data.Update(request.Id, payment)
	if err != nil {
		log.Fatal("got an error when tried to update order")
	}
	return &pb.UpdateResponse{}, nil
}

func (o orderServer) Delete(ctx context.Context, request *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	err := o.data.Delete(request.Id)
	if err != nil {
		log.Fatal("got an error when tried to delete order")
	}
	return &pb.DeleteResponse{}, nil
}
