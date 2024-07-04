package websiteserver

import (
	"html/template"
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

func NewMockServer(page *Page) WebsiteServerInterface {
	return &MockServer{
		srv:  &http.Server{Addr: ":3000"},
		page: page,
	}
}

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

func (ms *MockServer) UpdateBody(content string) {
	log.Printf("Updating body to: %s\n", content)
	ms.page.Body = template.HTML(content)
}

func (ms *MockServer) GetHTML() string {
	return string(ms.page.HTML)
}

func (ms *MockServer) GetURL() string {
	return ms.srv.Addr
}
