package products

type Service struct {
	Repository Repository
}

func NewService(repository Repository) *Service {
	return &Service{Repository: repository}
}

func (service *Service) Get(id string) (product *Product, err error) {
	product, err = service.Repository.Get(id)
	if err != nil {
		switch err {
		case ErrProductNotFound:
			break
		// More cases here...
		default:
			err = ErrUnexpected
		}
		return
	}

	return
}
