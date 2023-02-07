package internal

type Service interface {
	GetAll() ([]User, error)
	Store(id int, name, email string, age int) (User, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) GetAll() ([]User, error) {
	return s.repo.GetAll()
}

func (s *service) Store(id int, name, email string, age int) (User, error) {
	return s.repo.Store(id, name, email, age)
}
