package mocks

// constructor
func NewPreyStub() *PreyStub {
	return &PreyStub{}
}

// controller
type PreyStub struct {
}

func (st *PreyStub) GetSpeed() float64 {
	return 10
}
