package webserver

import "github.com/stretchr/testify/mock"

type MockFileReader struct {
	mock.Mock
}

func (m *MockFileReader) ReadFile(filename string) ([]byte, error) {
	args := m.Called(filename)
	return args.Get(0).([]byte), args.Error(1)
}
