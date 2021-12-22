package main

import (
	"context"
	"google.golang.org/grpc"
	"klever/cryptocurrency"
	"log"
)

func main() {

	// setup grpc connection
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":8000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %s", err)
	}
	defer conn.Close()

	// use CryptocurrencyClient
	c := cryptocurrency.NewCryptocurrencyServiceClient(conn)

	// create new currency
	response, err := c.List(context.Background(), &cryptocurrency.EmptyMessage{})
	if err != nil {
		log.Fatalf("Could not connect: %s", err)
	}
	for _, currency := range response.Cryptocurrency {
		log.Printf("Symbol %s\n", currency.Symbol)
		log.Printf("Name %s\n", currency.Name)
		log.Printf("Icon %s\n", currency.IconURL)
		log.Printf("Votes %d\n", currency.Votes)
		log.Println("---------------------")
	}
}
