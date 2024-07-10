package htmlservice

import (
	"github.com/stretchr/testify/mock"
	"golang.org/x/net/html"
)

// Ensure MockHTMLService implements HTMLServiceInterface
var _ HTMLServiceInterface = (*MockHTMLService)(nil)

type MockHTMLService struct {
	mock.Mock
}

func (m *MockHTMLService) ParseHTML(htmlData []byte) (*html.Node, error) {
	args := m.Called(htmlData)
	return args.Get(0).(*html.Node), args.Error(1)
}

func (m *MockHTMLService) RenderHTML(doc *html.Node) ([]byte, error) {
	args := m.Called(doc)
	return args.Get(0).([]byte), args.Error(1)
}
