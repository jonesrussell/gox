package websiteserver

import (
	"net/http"
	"sync"
)

type MockServer struct {
	mux  *http.ServeMux
	srv  *http.Server
	wg   sync.WaitGroup
	page *Page
}

// Ensure MockServer implements Server interface
var _ WebsiteServerInterface = &MockServer{}

func (ms *MockServer) Start() error { return nil }

func (ms *MockServer) Stop() error {
	// Assuming ms.srv is the server you want to stop
	ms.srv.Close()
	// Wait for any goroutine to finish
	ms.wg.Wait()
	return nil
}

func (ms *MockServer) UpdateTitle(title string) {
	ms.page.Title = title
}

func NewMockServer() WebsiteServerInterface {
	return &MockServer{
		mux: http.NewServeMux(),
		srv: &http.Server{},
	}
}
