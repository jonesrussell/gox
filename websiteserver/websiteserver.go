package websiteserver

import (
	"jonesrussell/gocreate/utils"
	"log"
	"net/http"
	"sync"
)

// Define the WebsiteServerInterface interface
type WebsiteServerInterface interface {
	Start() error
	Stop() error
	UpdateTitle(title string)
}

// websiteServerImpl is the actual implementation of the Server interface
type websiteServerImpl struct {
	mux  *http.ServeMux
	srv  *http.Server
	wg   sync.WaitGroup
	page *Page
}

// NewServer returns a new Server
func NewServer() WebsiteServerInterface {
	// Explicitly use the FileReader interface when creating a new Page instance
	page := NewPage("", utils.OSFileReader{}) // utils.OSFileReader{} is of type utils.FileReader

	return &websiteServerImpl{
		mux:  http.NewServeMux(),
		srv:  &http.Server{Addr: ":3000"},
		page: page,
	}
}

func (s *websiteServerImpl) Start() error {
	s.mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := handleRequest(w, r, s.page)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			log.Println(err)
		}
	})

	s.srv.Handler = s.mux

	s.wg.Add(1)
	go func() {
		defer s.wg.Done()

		log.Println("Listening on :3000...")
		if err := s.srv.ListenAndServe(); err != http.ErrServerClosed {
			log.Println(err)
		}
	}()

	return nil
}

func (s *websiteServerImpl) Stop() error {
	if s.srv != nil {
		err := s.srv.Close()
		if err != nil {
			return err
		}
		s.wg.Wait() // Wait for the server to shutdown
	}
	return nil
}

func (s *websiteServerImpl) UpdateTitle(title string) {
	s.page.Title = title
}