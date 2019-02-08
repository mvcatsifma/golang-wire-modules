package products

type Service struct {
	dbClient IDbClient
}

type IDbClient interface {
	Query()
}

func NewService(client IDbClient) *Service {
	return &Service{dbClient: client}
}

func (this *Service) DoGet() {
	this.dbClient.Query()
}
