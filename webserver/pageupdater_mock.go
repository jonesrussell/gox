package webserver

import (
	"github.com/stretchr/testify/mock"
	"golang.org/x/net/html"
)

type MockPageUpdater struct {
	mock.Mock
}

func (m *MockPageUpdater) UpdatePage(title, body, templatePath string) (string, error) {
	args := m.Called(title, body, templatePath)
	return args.String(0), args.Error(1)
}

func (m *MockPageUpdater) ChangeTitle(doc *html.Node, title string) {
	m.Called(doc, title)
}

func (m *MockPageUpdater) ChangeBody(doc *html.Node, body string) {
	m.Called(doc, body)
}
