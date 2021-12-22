package main

import (
	"fmt"
	"google.golang.org/grpc"
	"klever/cryptocurrency"
	"klever/utils"
	"log"
	"net"
)

func main() {
	fmt.Println("Starting cryptocurrency voting app...")
	fmt.Printf("%s environment detected!\n", utils.GetAppEnvironment())

	// Setup listener
	listener, err := net.Listen("tcp", ":"+utils.GetConfig().ServerPort)
	if err != nil {
		log.Fatalf("Failed to listen on port " + utils.GetConfig().ServerPort)
	}

	// Setup gRPC server
	grpcServer := grpc.NewServer()

	server := cryptocurrency.Server{}

	cryptocurrency.RegisterCryptocurrencyServiceServer(grpcServer, &server)

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve gRPC server")
	}

}
