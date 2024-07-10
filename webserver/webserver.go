package webserver

import (
	"fmt"
	"html/template"
	"jonesrussell/gocreate/htmlservice"
	"jonesrussell/gocreate/logger"
	"net/http"
	"sync"

	sse "github.com/r3labs/sse/v2"
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

// NewWebServer returns a new Server
func NewWebServer(logger logger.LoggerInterface) WebServerInterface {
	logger.Debug("Creating a new web server...")

	// Create a new WebsiteUpdater
	updater := NewPageUpdater(logger)

	body := "<h1>My Heading</h1>"

	// Create a new HTMLService
	htmlService := htmlservice.NewHTMLService()

	// Explicitly use the FileReader interface when creating a new Page instance
	page, err := NewPage("", template.HTML(body), updater, "static/index.html", logger, htmlService)
	if err != nil {
		logger.Error("Error creating page: ", err)
		return nil
	}

	s := &webServer{
		logger:    logger,
		mux:       http.NewServeMux(),
		srv:       &http.Server{Addr: ":3000"},
		page:      page,
		sseServer: sse.New(),
	}

	// Create the SSE server
	s.sseServer.CreateStream("messages")

	return s
}

func (s *webServer) Logger() logger.LoggerInterface {
	return s.logger
}

func (s *webServer) setupRoutes() {
	// Serve the SSE endpoint
	s.mux.HandleFunc("/updates", s.sseServer.ServeHTTP)

	// Serve static files from the "static/js" directory
	s.mux.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("static/js"))))

	// Serve static html files from the "static" directory
	s.mux.HandleFunc("/", s.handleRootRequest)
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
			s.logger.Error("Error stopping server: ", err)
			return err
		}
		s.wg.Wait() // Wait for the server to shutdown
	}
	return nil
}

func (s *webServer) UpdateTitle(content string) {
	s.logger.Debug("UpdateTitle called with content: '" + content + "'")
	s.page.Template.SetTitle(content)
	s.ssePublishUpdate(content, s.page.Template.GetBody()) // pass the title and body to the function
	s.logger.Debug("Title updated, sending update signal")
	select {
	case s.page.UpdateChan <- struct{}{}:
		s.logger.Debug("Update signal sent successfully")
	default:
		s.logger.Debug("Update channel is full, skipping signal")
	}
}

func (s *webServer) UpdateBody(content string) {
	s.logger.Debug("UpdateBody called with content: " + content)
	s.page.Template.SetBody(content)
	s.ssePublishUpdate(s.page.Template.GetTitle(), content) // pass the title and body to the function
	s.logger.Debug("Body updated, sending update signal")
	select {
	case s.page.UpdateChan <- struct{}{}:
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
	return s.page.UpdateChan
}

func (s *webServer) ssePublishUpdate(title string, body string) {
	s.logger.Debug("ssePublishUpdate called")

	// Create a new event
	event := &sse.Event{
		Data: []byte(fmt.Sprintf("Title: %s, Body: %s", title, body)),
	}

	// Publish the event to the "messages" channel
	s.sseServer.Publish("messages", event)
}
