package cryptocurrency

type Cryptocurrency struct {
	Symbol  string `bson:"_id"`
	Name    string `bson:"name"`
	IconURL string `bson:"icon_url"`
	Votes   int32  `bson:"votes"`
}
