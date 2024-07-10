package webserver

import (
	"html/template"
	"jonesrussell/gocreate/htmlservice"
	"jonesrussell/gocreate/logger"
	"log"
	"net/http"
	"sync"

	sse "github.com/r3labs/sse/v2"
	"github.com/stretchr/testify/mock"
	"golang.org/x/net/html"
)

type MockServer struct {
	srv        *http.Server
	wg         sync.WaitGroup
	page       *Page
	log        logger.LoggerInterface
	updateChan chan struct{}
	sseServer  *sse.Server
}

var _ WebServerInterface = (*MockServer)(nil)

func NewMockServer(page *Page) WebServerInterface {
	var err error
	wslog, err := logger.NewLogger("/tmp/gocreate-test.go")
	if err != nil {
		log.Fatalf("Error initializing logger: %v", err)
	}

	// Create a new MockHTMLService
	mockHTMLService := new(htmlservice.MockHTMLService)

	// Set up the ParseHTML method of the MockHTMLService to return a specific value
	mockHTMLService.On("ParseHTML", mock.Anything).Return(&html.Node{}, nil)

	if page == nil {
		page, err = NewPage(
			"Mock Title",
			template.HTML("<h1>Mock Body</h1>"),
			NewPageUpdater(wslog),
			"../static/index.html",
			wslog,
			mockHTMLService, // Pass the mock HTMLService to NewPage
		)
		if err != nil {
			wslog.Error("Error creating page: ", err)
			return nil
		}
	}

	return &MockServer{
		srv:       &http.Server{Addr: ":3000"},
		wg:        sync.WaitGroup{},
		page:      page,
		sseServer: sse.New(),
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

func (ms *MockServer) GetUpdateChan() <-chan struct{} {
	return ms.updateChan
}

func (ms *MockServer) UpdateTitle(title string) {
	log.Printf("Updating title to: %s\n", title)
	ms.page.Template.Title = title
	ms.notifyUpdate()
}

func (ms *MockServer) UpdateBody(content string) {
	log.Printf("Updating body to: %s\n", content)
	ms.page.Template.Body = template.HTML(content)
	ms.notifyUpdate()
}

func (ms *MockServer) GetHTML() string {
	return string(ms.page.Template.HTML)
}

func (ms *MockServer) notifyUpdate() {
	select {
	case ms.updateChan <- struct{}{}:
	default:
	}
}

func (ms *MockServer) Logger() logger.LoggerInterface {
	return ms.log
}

func (ms *MockServer) GetURL() string {
	return "http://127.0.0.1:3000"
}
