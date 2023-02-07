package products

type Storage interface {
	Get(id string) (product *Product, err error)
}
