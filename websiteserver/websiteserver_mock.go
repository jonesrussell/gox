package websiteserver

import (
	"log"
	"net/http"
	"sync"
)

type MockServer struct {
	srv  *http.Server
	wg   sync.WaitGroup
	page *Page
}

var _ WebsiteServerInterface = &MockServer{}

func (ms *MockServer) Start() error {
	log.Println("Starting MockServer")
	// Additional startup logic can be logged here
	return nil
}

func (ms *MockServer) Stop() error {
	log.Println("Stopping MockServer")
	// Log any cleanup actions taken here
	ms.srv.Close()
	ms.wg.Wait()
	return nil
}

func (ms *MockServer) UpdateTitle(title string) {
	log.Printf("Updating title to: %s\n", title)
	ms.page.Title = title
}

func NewMockServer(page *Page) WebsiteServerInterface {
	log.Println("Creating NewMockServer instance")
	return &websiteServerImpl{
		mux:  http.NewServeMux(),
		srv:  &http.Server{Addr: ":3000"},
		page: page,
	}
}
