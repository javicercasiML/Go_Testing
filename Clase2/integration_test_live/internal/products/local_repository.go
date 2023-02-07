package products

type LocalRepository struct {
	DB Storage
}

func NewLocalRepository(db Storage) *LocalRepository {
	return &LocalRepository{
		DB: db,
	}
}

func (repository *LocalRepository) Get(id string) (product *Product, err error) {
	product, err = repository.DB.Get(id)
	if err != nil {
		return
	}

	if product == nil {
		err = ErrProductNotFound
		return
	}

	return
}
