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

	// down vote a currency
	newCurrency := cryptocurrency.CryptocurrencySymbol{
		Symbol: "BTC",
	}

	response, err := c.DownVote(context.Background(), &newCurrency)
	if err != nil {
		log.Fatalf("Could not connect: %s", err)
	}
	log.Printf("Symbol %s\n", response.Symbol)
	log.Printf("Name %s\n", response.Name)
	log.Printf("Icon %s\n", response.IconURL)
	log.Printf("Votes %d\n", response.Votes)
}
