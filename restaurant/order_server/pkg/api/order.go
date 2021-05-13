package api

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strconv"
	"time"

	"github.com/miladouski/golang-training-restaurant-grpc/restaurant/order_server/pkg/data"
	pb "github.com/miladouski/golang-training-restaurant-grpc/restaurant/proto/proto"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
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
		s := status.Newf(codes.Internal, "got an error when tried to read orders")
		errWithDetails, err := s.WithDetails(request)
		if err != nil {
			return nil, status.Errorf(codes.Unknown, "can't covert status to status with details %v", s)
		}
		return nil, errWithDetails.Err()
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
		log.WithFields(log.Fields{"Id": orders[i].Id,
			"date":    orders[i].Date,
			"table":   orders[i].TableNumber,
			"waiter":  orders[i].WaiterId,
			"price":   orders[i].Price,
			"payment": orders[i].Payment}).
			Info("Read all orders")
		pbOrders = append(pbOrders, order)
	}
	return &pb.ReadAllResponse{Order: pbOrders}, nil
}

func (o orderServer) Read(ctx context.Context, request *pb.ReadRequest) (*pb.ReadResponse, error) {
	order, err := o.data.Read(request.Id)
	if err != nil {
		s := status.Newf(codes.InvalidArgument, "got an error when tried to read order with id: %d", request.Id)
		errWithDetails, err := s.WithDetails(request)
		if err != nil {
			return nil, status.Errorf(codes.Unknown, "can't covert status to status with details %v", s)
		}
		return nil, errWithDetails.Err()
	}
	pbOrder := pb.Order{
		Id:          order.Id,
		Date:        order.Date.String(),
		TableNumber: order.TableNumber,
		WaiterId:    order.WaiterId,
		Price:       order.Price,
		Payment:     order.Payment,
	}
	log.WithFields(log.Fields{"Id": order.Id,
		"date":    order.Date,
		"table":   order.TableNumber,
		"waiter":  order.WaiterId,
		"price":   order.Price,
		"payment": order.Payment}).
		Info("Read order")
	return &pb.ReadResponse{Order: &pbOrder}, nil
}

func (o orderServer) Create(ctx context.Context, request *pb.CreateRequest) (*pb.CreateResponse, error) {
	date, err := time.Parse("2006-01-02", request.Order.Date)
	if err != nil {
		s := status.Newf(codes.InvalidArgument, "wrong date format: %s", request.Order.Date)
		errWithDetails, err := s.WithDetails(request)
		if err != nil {
			return nil, status.Errorf(codes.Unknown, "can't covert status to status with details %v", s)
		}
		return nil, errWithDetails.Err()
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
		s := status.Newf(codes.InvalidArgument, "got an error when tried to create order with id: %d", request.Order.Id)
		errWithDetails, err := s.WithDetails(request)
		if err != nil {
			return nil, status.Errorf(codes.Unknown, "can't covert status to status with details %v", s)
		}
		return nil, errWithDetails.Err()
	}
	log.WithFields(log.Fields{"Id": order.Id,
		"date":    order.Date,
		"table":   order.TableNumber,
		"waiter":  order.WaiterId,
		"price":   order.Price,
		"payment": order.Payment}).
		Info("Create order")
	return &pb.CreateResponse{}, nil
}

func (o orderServer) Update(ctx context.Context, request *pb.UpdateRequest) (*pb.UpdateResponse, error) {
	payment, _ := strconv.ParseBool(request.Payment)
	err := o.data.Update(request.Id, payment)
	if err != nil {
		s := status.Newf(codes.InvalidArgument, "got an error when tried to update order with id: %d", request.Id)
		errWithDetails, err := s.WithDetails(request)
		if err != nil {
			return nil, status.Errorf(codes.Unknown, "can't covert status to status with details %v", s)
		}
		return nil, errWithDetails.Err()
	}
	log.WithFields(log.Fields{"Id": request.Id, "payment": payment}).
		Info("Update order")
	return &pb.UpdateResponse{}, nil
}

func (o orderServer) Delete(ctx context.Context, request *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	err := o.data.Delete(request.Id)
	if err != nil {
		s := status.Newf(codes.InvalidArgument, "got an error when tried to delete order with id: %d", request.Id)
		errWithDetails, err := s.WithDetails(request)
		if err != nil {
			return nil, status.Errorf(codes.Unknown, "can't covert status to status with details %v", s)
		}
		return nil, errWithDetails.Err()
	}
	log.WithFields(log.Fields{"Id": request.Id}).
		Info("Update order")
	return &pb.DeleteResponse{}, nil
}
