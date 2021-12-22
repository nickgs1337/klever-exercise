package main

import (
	"context"
	"google.golang.org/grpc"
	"io"
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

	// up vote a currency
	newCurrency := cryptocurrency.CryptocurrencySymbol{
		Symbol: "BTC",
	}

	stream, err := c.CreateVoteStream(context.Background(), &newCurrency)
	if err != nil {
		log.Fatalf("Could not connect: %s", err)
	}

	done := make(chan bool)

	go func() {
		for {
			response, err := stream.Recv()
			if err == io.EOF {
				done <- true //means stream is finished
				return
			}
			if err != nil {
				log.Fatalf("cannot receive %v", err)
			}
			log.Printf("Vote update for %s received: %d", newCurrency.Symbol, response.Votes)
		}
	}()

	<-done
	log.Printf("Finished")
	//log.Printf("Symbol %s\n", response.Symbol)
	//log.Printf("Name %s\n", response.Name)
	//log.Printf("Icon %s\n", response.IconURL)
	//log.Printf("Votes %d\n", response.Votes)
}
