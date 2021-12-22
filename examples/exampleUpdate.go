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

	// update a currency
	newCurrency := cryptocurrency.UpdateCryptocurrencyRequest{
		OldSymbol: "BTC",
		NewCryptocurrency: &cryptocurrency.CryptocurrencyMessage{
			Symbol:  "BTC",
			Name:    "Bitcoin 2",
			IconURL: "https://s2.coinmarketcap.com/static/img/coins/64x64/1.png",
		},
	}

	response, err := c.Update(context.Background(), &newCurrency)
	if err != nil {
		log.Fatalf("Could not connect: %s", err)
	}
	log.Printf("%s updated!", newCurrency.OldSymbol)
	log.Printf("Symbol %s\n", response.Symbol)
	log.Printf("Name %s\n", response.Name)
	log.Printf("Icon %s\n", response.IconURL)
	log.Printf("Votes %d\n", response.Votes)
}
