package webserver

import (
	"html/template"
	"jonesrussell/gocreate/logger"
	"jonesrussell/gocreate/utils"
	"log"
	"net/http"
	"sync"
)

type MockServer struct {
	srv  *http.Server
	wg   sync.WaitGroup
	page *Page
	log  logger.LoggerInterface
}

var _ WebServerInterface = (*MockServer)(nil)

func NewMockServer(page *Page) WebServerInterface {
	if page == nil {
		page = NewPage(
			"Mock Title",
			template.HTML("<h1>Mock Body</h1>"),
			utils.OSFileReader{},
			NewWebsiteUpdater(nil), "../static/index.html",
		)
	}

	var err error
	wslog, err := logger.NewLogger("/tmp/gocreate-test.go")
	if err != nil {
		log.Fatalf("Error initializing logger: %v", err)
	}

	return &MockServer{
		srv:  &http.Server{Addr: ":3000"},
		page: page,
		log:  wslog,
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

func (ms *MockServer) Logger() logger.LoggerInterface {
	return ms.log
}

func (ms *MockServer) GetURL() string {
	return "http://127.0.0.1:3000"
}
