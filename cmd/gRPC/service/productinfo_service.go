package main

import (
	"context"
	"github.com/gofrs/uuid"
	"github.com/golang/protobuf/ptypes/wrappers"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	pb "learning-grpc/service/ecommerce"
)

// 用结构体来作为服务器
type server struct {
	productMap map[string]*pb.Product
	orderMap   map[string]*pb.Order
	pb.UnsafeProductInfoServer
}

// AddProduct 定义服务器要实现的方法，使用Product和ProductID两种消息（Message）来实现业务逻辑
func (s *server) AddProduct(ctx context.Context, in *pb.Product) (*pb.ProductID, error) {
	out, err := uuid.NewV4()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error while generating Product ID", err)
	}
	in.Id = out.String()
	if s.productMap == nil {
		s.productMap = make(map[string]*pb.Product)
	}
	s.productMap[in.Id] = in
	return &pb.ProductID{Value: in.Id}, status.New(codes.OK, "").Err()
}

func (s *server) GetProduct(ctx context.Context, in *pb.ProductID) (*pb.Product, error) {
	value, ok := s.productMap[in.Value]
	if ok {
		return value, status.New(codes.OK, "").Err()
	}
	return nil, status.Errorf(codes.NotFound, "Product does not exist.", in.Value)
}

func (s *server) GetOrder(ctx context.Context, orderId *wrappers.StringValue) (*pb.Order, error) {
	ord := s.orderMap[orderId.Value]
	return ord, nil
}
