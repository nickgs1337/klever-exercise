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
	newCurrency := cryptocurrency.CryptocurrencyMessage{
		Symbol:  "BTC",
		Name:    "Bitcoin",
		IconURL: "https://s2.coinmarketcap.com/static/img/coins/64x64/1.png",
	}

	response, err := c.Create(context.Background(), &newCurrency)
	if err != nil {
		log.Fatalf("Could not connect: %s", err)
	}
	log.Printf("%s created!", response.Symbol)
}
