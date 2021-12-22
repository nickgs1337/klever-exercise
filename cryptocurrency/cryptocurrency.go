package cryptocurrency

import (
	"golang.org/x/net/context"
	"log"
	"sync"
)

type Server struct {
	connections []*Connection
}

type Connection struct {
	stream CryptocurrencyService_CreateVoteStreamServer
	symbol string
	err    chan error
}

func (server *Server) Create(context context.Context, message *CryptocurrencyMessage) (*CryptocurrencyMessage, error) {
	//validate request
	err := validateCryptocurrencyMessage(message)
	if err != nil {
		return nil, err
	}

	//perform creation
	newCurrency, err := Create(fromMessage(message))
	if err != nil {
		return nil, err
	}

	return toMessage(newCurrency), nil
}

func (server *Server) Get(context context.Context, symbol *CryptocurrencySymbol) (*CryptocurrencyMessage, error) {
	currency, err := GetBySymbol(symbol.Symbol)

	if err != nil {
		return nil, err
	}

	return toMessage(currency), nil
}

func (server *Server) List(context context.Context, message *EmptyMessage) (*CryptocurrencyListMessage, error) {
	currencies, err := List()
	if err != nil {
		return nil, err
	}

	var cryptoMessages []*CryptocurrencyMessage
	for _, currency := range currencies {
		cryptoMessages = append(cryptoMessages, toMessage(&currency))
	}

	return &CryptocurrencyListMessage{Cryptocurrency: cryptoMessages}, nil
}

func (server *Server) Update(context context.Context, request *UpdateCryptocurrencyRequest) (*CryptocurrencyMessage, error) {
	//validate request
	err := validateCryptocurrencyMessage(request.NewCryptocurrency)
	if err != nil {
		return nil, err
	}

	//perform update
	currency, err := Update(request.OldSymbol, fromMessage(request.NewCryptocurrency))

	if err != nil {
		return nil, err
	}

	return toMessage(currency), nil
}

func (server *Server) Delete(context context.Context, symbol *CryptocurrencySymbol) (*EmptyMessage, error) {
	err := Delete(symbol.Symbol)

	if err != nil {
		return nil, err
	}

	return &EmptyMessage{}, nil
}

func (server *Server) UpVote(context context.Context, symbol *CryptocurrencySymbol) (*CryptocurrencyMessage, error) {
	currency, err := UpVote(symbol.Symbol)

	if err != nil {
		return nil, err
	}
	notifyVoteListeners(server, currency)

	return toMessage(currency), nil
}

func (server *Server) DownVote(context context.Context, symbol *CryptocurrencySymbol) (*CryptocurrencyMessage, error) {
	currency, err := DownVote(symbol.Symbol)

	if err != nil {
		return nil, err
	}
	notifyVoteListeners(server, currency)

	return toMessage(currency), nil
}

func (server *Server) CreateVoteStream(symbol *CryptocurrencySymbol, stream CryptocurrencyService_CreateVoteStreamServer) error {
	// validate if currency exists
	_, err := GetBySymbol(symbol.Symbol)
	if err != nil {
		return err
	}

	conn := &Connection{
		stream: stream,
		symbol: symbol.Symbol,
		err:    make(chan error),
	}
	server.connections = append(server.connections, conn)
	log.Printf("New vote listener for %s registered!", symbol.Symbol)
	return <-conn.err
}

func (server *Server) mustEmbedUnimplementedCryptocurrencyServiceServer() {
	//TODO implement me
	panic("implement me")
}

func notifyVoteListeners(server *Server, cryptocurrency *Cryptocurrency) {
	wait := sync.WaitGroup{}
	done := make(chan int)

	for _, conn := range server.connections {
		wait.Add(1)
		go func(cryptocurrency *Cryptocurrency, conn *Connection) {
			defer wait.Done()

			if conn.symbol == cryptocurrency.Symbol {
				err := conn.stream.Send(&CryptocurrencyNewVoteNotification{Votes: cryptocurrency.Votes})
				log.Printf("Sending vote notification of %s", conn.symbol)
				if err != nil {
					conn.err <- err
					log.Fatalf("Error with stream %v, Error: %v", conn.stream, err)
				}
			}
		}(cryptocurrency, conn)
	}

	go func() {
		wait.Wait()
		close(done)
	}()

	<-done
}

func fromMessage(message *CryptocurrencyMessage) *Cryptocurrency {
	return &Cryptocurrency{
		Symbol:  message.Symbol,
		Name:    message.Name,
		IconURL: message.IconURL,
		Votes:   0,
	}
}

func toMessage(cryptocurrency *Cryptocurrency) *CryptocurrencyMessage {
	return &CryptocurrencyMessage{
		Symbol:  cryptocurrency.Symbol,
		Name:    cryptocurrency.Name,
		IconURL: cryptocurrency.IconURL,
		Votes:   cryptocurrency.Votes,
	}
}
