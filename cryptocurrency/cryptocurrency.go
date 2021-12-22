package cryptocurrency

import "golang.org/x/net/context"

type Server struct {
	connections []*Connection
}

func (server *Server) Create(context context.Context, message *CryptocurrencyMessage) (*CryptocurrencyMessage, error) {
	newCurrency, err := Create(fromMessage(message))

	if err != nil {
		return nil, err
	}

	return toMessage(newCurrency), nil
}

func (server *Server) Get(context context.Context, symbol *CryptocurrencySymbol) (*CryptocurrencyMessage, error) {
	//TODO implement me
	panic("implement me")
}

func (server *Server) List(context context.Context, message *EmptyMessage) (*CryptocurrencyListMessage, error) {
	//TODO implement me
	panic("implement me")
}

func (server *Server) Update(context context.Context, request *UpdateCryptocurrencyRequest) (*CryptocurrencyMessage, error) {
	//TODO implement me
	panic("implement me")
}

func (server *Server) Delete(context context.Context, symbol *CryptocurrencySymbol) (*EmptyMessage, error) {
	//TODO implement me
	panic("implement me")
}

func (server *Server) UpVote(context context.Context, symbol *CryptocurrencySymbol) (*CryptocurrencyMessage, error) {
	//TODO implement me
	panic("implement me")
}

func (server *Server) DownVote(context context.Context, symbol *CryptocurrencySymbol) (*CryptocurrencyMessage, error) {
	//TODO implement me
	panic("implement me")
}

func (server *Server) CreateVoteStream(symbol *CryptocurrencySymbol, stream CryptocurrencyService_CreateVoteStreamServer) error {
	//TODO implement me
	panic("implement me")
}

func (server *Server) mustEmbedUnimplementedCryptocurrencyServiceServer() {
	//TODO implement me
	panic("implement me")
}

type Connection struct {
	stream CryptocurrencyService_CreateVoteStreamServer
	symbol string
	err    chan error
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
