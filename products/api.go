package products

type Api struct {
	Service *Service
}

func NewApi(service *Service) *Api {
	return &Api{Service: service}
}

func (this Api) DoGet() {
	this.Service.DoGet()
}
