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

	// delete a currency
	newCurrency := cryptocurrency.CryptocurrencySymbol{
		Symbol: "BTC",
	}

	_, err = c.Delete(context.Background(), &newCurrency)
	if err != nil {
		log.Fatalf("Could not connect: %s", err)
	}
	log.Printf("Deleted!\n")
}
