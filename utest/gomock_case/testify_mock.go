package gomock_case

import (
	"github.com/pkg/errors"
	"github.com/stretchr/testify/mock"
)

// mock代码
type TestDD struct {
	mock.Mock
}

func (m *TestDD) Send(key string, value []byte) error {
	// args := m.Called(key, value)
	// switch key {
	// case "no_value":
	// 	return errors.New("empty values")
	// case "h1":
	// 	return nil
	// default:
	// 	return errors.New("exception ky")
	// }
	return nil
}

func (m *TestDD) Get(key string) ([]byte, error) {
	switch key {
	case "no_value":
		return nil, errors.New("aaaaaa empty values")
	case "h1":
		return []byte("exist h1"), nil
	default:
		return nil, errors.New("exception key")
	}
}
