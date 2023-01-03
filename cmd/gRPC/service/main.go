package main

import (
	"google.golang.org/grpc"
	pb "learning-grpc/service/ecommerce"
	"log"
	"net"
)

const (
	port = ":50051"
)

func main() {
	// 设置监听端口
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// 创建一个新的gRPC服务器
	s := grpc.NewServer()
	// 将ProductInfo这个Service注册到新创建的gRPC服务器中，同时传入一个服务器的实例（server是定义了服务器的结构体）；
	// gRPC服务器将使用实例来提供ProductInfo这个服务（Service）
	pb.RegisterProductInfoServer(s, &server{})
	log.Printf("Starting gRPC listener on port " + port)
	// 服务器开始监听
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
