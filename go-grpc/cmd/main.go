package main

import (
	"log"
	"net"

	"go-grpc/cmd/config"
	"go-grpc/cmd/services"
	productPb "go-grpc/pb/product"

	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

func main() {

	netListen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen %v", err.Error())
	}

	db := config.ConnectDatabase()

	grpcServer := grpc.NewServer()
	productService := services.ProductService{DB: db}
	productPb.RegisterProductServiceServer(grpcServer, &productService)

	if err := grpcServer.Serve(netListen); err != nil {
		log.Fatalf("Failed to server %v", err.Error())
	}
}
