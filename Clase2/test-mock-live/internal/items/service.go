package items

import (
	"testmock/internal/domain"
	"time"
)

// constructor
func NewService(st Storage) Service {
	return &service{st: st}
}

// controller
const (
	YYYY_MM_DD = "2006-01-02"
)

type service struct {
	st Storage
}
// read
func (sv *service) GetTotalByName(name string, quantity int) (total float64, err error) {
	var i domain.Item
	i, err = sv.st.GetByName(name)
	if err != nil {
		return
	}

	total += i.Price * float64(quantity)

	return
}
// write
func (sv *service) UpdateByName(name string, weight, price *float64, release *string) (i domain.Item, err error) {
	// get
	i, err = sv.st.GetByName(name)
	if err != nil {
		return
	}

	// check
	if weight != nil {
		i.Weight = *weight
	}
	if price != nil {
		i.Price = *price
	}
	if release != nil {
		var dt time.Time
		dt, err = time.Parse(YYYY_MM_DD, *release)
		if err != nil {
			err = ErrServiceInternal
			return
		}

		i.Release = dt
	}

	// validate
	if !i.Valid() {
		err = ErrServiceInvalidDomain
		return
	}

	// update
	err = sv.st.UpdateByName(name, i)
	return
}