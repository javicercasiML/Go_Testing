package mocks

// constructor
func NewSimulatorMock(mttc float64) *simulatorMock {
	return &simulatorMock{maxTimeToCatch: mttc}
}

// controller
type simulatorMock struct {
	ResultOnCan bool
	ResultOnGet float64
	Spy         bool

	maxTimeToCatch float64
}

func (st *simulatorMock) CanCatch(distance float64, speed float64, catchSpeed float64) bool {
	timeToCatch := distance / (speed - catchSpeed)
	st.ResultOnCan = timeToCatch > 0 && timeToCatch <= st.maxTimeToCatch
	return st.ResultOnCan
}

func (st *simulatorMock) GetLinearDistance(position [2]float64) float64 {
	st.Spy = true
	st.ResultOnGet = position[0] * position[1]
	return st.ResultOnGet
}
