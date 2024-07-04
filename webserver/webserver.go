package webserver

import (
	"html/template"
	"jonesrussell/gocreate/logger"
	"jonesrussell/gocreate/utils"
	"net"
	"net/http"
	"sync"
)

// Define the WebServerInterface interface
type WebServerInterface interface {
	Start() error
	Stop() error
	UpdateTitle(title string)
	UpdateBody(body string)
	GetHTML() string
	GetURL() string
}

// webServer is the actual implementation of the Server interface
type webServer struct {
	logger logger.LoggerInterface
	mux    *http.ServeMux
	srv    *http.Server
	wg     sync.WaitGroup
	page   *Page
}

// NewServer returns a new Server
func NewServer(logger logger.LoggerInterface) WebServerInterface {
	logger.Debug("Creating a new web server...")

	// Create a new WebsiteUpdater
	updater := NewWebsiteUpdater(logger)

	// Explicitly use the FileReader interface when creating a new Page instance
	body := "<h1>My Heading</h1>"
	page := NewPage("My Title", template.HTML(body), utils.OSFileReader{}, updater, "static/index.html") // utils.OSFileReader{} is of type utils.FileReader

	return &webServer{
		logger: logger,
		mux:    http.NewServeMux(),
		srv:    &http.Server{Addr: ":3000"},
		page:   page,
	}
}

func (s *webServer) Start() error {
	// Check for errors before starting the goroutine
	listener, err := net.Listen("tcp", ":3000")
	if err != nil {
		s.logger.Error("Error starting server: ", err)
		return err
	}

	s.mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := handleRequest(w, r, s.page)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			s.logger.Error("Error handling request: ", err)
		}
	})

	s.srv.Handler = s.mux

	s.wg.Add(1)
	go func() {
		defer s.wg.Done()

		if err := s.srv.Serve(listener); err != http.ErrServerClosed {
			s.logger.Error("Server stopped with error: ", err)
		} else {
			s.logger.Debug("Server stopped normally")
		}
	}()

	return nil
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
	s.page.SetTitle(content)
}

func (s *webServer) UpdateBody(content string) {
	s.page.SetBody(content)
}

func (s *webServer) GetHTML() string {
	return s.page.GetHTML()
}

func (s *webServer) GetURL() string {
	addr := s.srv.Addr
	if addr == ":3000" {
		addr = "127.0.0.1:3000"
	} else if addr == "0.0.0.0:3000" {
		addr = "127.0.0.1:3000"
	}
	return "http://" + addr
}
