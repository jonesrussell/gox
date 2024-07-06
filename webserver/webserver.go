package webserver

import (
	"html/template"
	"jonesrussell/gocreate/logger"
	"jonesrussell/gocreate/utils"
	"net/http"
	"sync"

	"github.com/tmaxmax/go-sse"
)

// Define the WebServerInterface interface
type WebServerInterface interface {
	Start() error
	Stop() error
	UpdateTitle(title string)
	UpdateBody(body string)
	GetHTML() string
	GetURL() string
	Logger() logger.LoggerInterface
	GetUpdateChan() <-chan struct{}
}

// webServer is the actual implementation of the Server interface
type webServer struct {
	logger    logger.LoggerInterface
	mux       *http.ServeMux
	srv       *http.Server
	wg        sync.WaitGroup
	page      *Page
	sseServer *sse.Server
}

// NewServer returns a new Server
func NewServer(logger logger.LoggerInterface) WebServerInterface {
	logger.Debug("Creating a new web server...")

	// Create a new WebsiteUpdater
	updater := NewPageUpdater(logger)

	body := "<h1>My Heading</h1>"

	// Explicitly use the FileReader interface when creating a new Page instance
	page := NewPage("My Title", template.HTML(body), utils.OSFileReader{}, updater, "static/index.html", logger)

	return &webServer{
		logger:    logger,
		mux:       http.NewServeMux(),
		srv:       &http.Server{Addr: ":3000"},
		page:      page,
		sseServer: &sse.Server{}, // Initialize the SSE server here
	}
}

func (s *webServer) Logger() logger.LoggerInterface {
	return s.logger
}

func (s *webServer) setupRoutes() {
	s.mux.HandleFunc("/", s.handleRootRequest)
	s.mux.HandleFunc("/updates", s.handleUpdatesRequest)
}

func (s *webServer) Start() error {
	s.setupRoutes()
	s.srv.Handler = s.mux

	s.wg.Add(1)
	go s.startServer()

	return nil
}

func (s *webServer) startServer() {
	defer s.wg.Done()

	if err := s.srv.ListenAndServe(); err != http.ErrServerClosed {
		s.logger.Error("Server stopped with error: ", err)
	} else {
		s.logger.Debug("Server stopped normally")
	}
}

func (s *webServer) Stop() error {
	s.logger.Debug("Stopping server")
	if s.srv != nil {
		err := s.srv.Close()
		if err != nil {
			return err
		}
		s.wg.Wait() // Wait for the server to shutdown
	}
	return nil
}

func (s *webServer) UpdateTitle(content string) {
	s.logger.Debug("UpdateTitle called with content: " + content)
	s.page.SetTitle(content)
	s.logger.Debug("Title updated, sending update signal")
	select {
	case s.page.updateChan <- struct{}{}:
		s.logger.Debug("Update signal sent successfully")
	default:
		s.logger.Debug("Update channel is full, skipping signal")
	}
}

func (s *webServer) UpdateBody(content string) {
	s.logger.Debug("UpdateBody called with content: " + content)
	s.page.SetBody(content)
	s.logger.Debug("Body updated, sending update signal")
	select {
	case s.page.updateChan <- struct{}{}:
		s.logger.Debug("Update signal sent successfully")
	default:
		s.logger.Debug("Update channel is full, skipping signal")
	}
}

func (s *webServer) GetHTML() string {
	return s.page.GetHTML()
}

func (s *webServer) GetURL() string {
	addr := s.srv.Addr
	if addr == ":3000" {
		addr = "localhost:3000"
	} else if addr == "0.0.0.0:3000" {
		addr = "localhost:3000"
	}
	return "http://" + addr
}

func (s *webServer) GetUpdateChan() <-chan struct{} {
	return s.page.updateChan
}
