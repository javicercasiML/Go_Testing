package mocks

// constructor
func NewStorageMock(data map[string]interface{}) *storageMock {
	return &storageMock{Data: data}
}

// controller
type storageMock struct {
	Data map[string]interface{}
	Err  error
	Spy  bool
}

func (s *storageMock) GetValue(key string) interface{} {
	s.Spy = true
	if val := s.Data[key]; val != nil {
		return val
	}
	return nil
}
