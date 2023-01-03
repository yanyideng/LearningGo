package main

import (
	"context"
	"google.golang.org/grpc"
	pb "learning-grpc/client/ecommerce"
	"log"
	"time"
)

const (
	address = "localhost:50051"
)

func main() {
	// 创建与gRPC服务器的链接
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	// 用这个链接创建新的客户端
	c := pb.NewProductInfoClient(conn)

	// 新Product的信息
	name := "Apple 11"
	description := "New iPhone!"
	price := float32(1000.0)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// 客户端调用方法，会通过gRPC发给服务器来处理并响应；客户端调用方法就和在本地调用一样，gRPC抽象化了底层的网络通信细节
	r, err := c.AddProduct(ctx, &pb.Product{Name: name, Description: description, Price: price})
	if err != nil {
		log.Fatalf("Could not add product: %v", err)
	}
	log.Printf("Product ID: %s added successfully", r.Value)

	product, err := c.GetProduct(ctx, &pb.ProductID{Value: r.Value})
	if err != nil {
		log.Fatalf("Could not get product: %v", err)
	}
	log.Printf("Prodct: %s", product.String())
}
