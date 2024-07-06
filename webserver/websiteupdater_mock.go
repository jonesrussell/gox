package webserver

import (
	"github.com/stretchr/testify/mock"
	"golang.org/x/net/html"
)

type MockWebsiteUpdater struct {
	mock.Mock
}

func (m *MockWebsiteUpdater) UpdateWebsite(title, body, templatePath string) (string, error) {
	args := m.Called(title, body, templatePath)
	return args.String(0), args.Error(1)
}

func (m *MockWebsiteUpdater) ChangeTitle(doc *html.Node, title string) {
	m.Called(doc, title)
}

func (m *MockWebsiteUpdater) ChangeBody(doc *html.Node, body string) {
	m.Called(doc, body)
}
