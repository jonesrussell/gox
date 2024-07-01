package httpserver

import (
	"log"
	"net/http"
	"sync"
)

// Define the Server interface
type Server interface {
	Start() error
	Stop() error
	UpdateTitle(title string)
}

// serverImpl is the actual implementation of the Server interface
type serverImpl struct {
	mux  *http.ServeMux
	srv  *http.Server
	wg   sync.WaitGroup
	page *Page
}

// NewServer returns a new Server
func NewServer() Server {
	return &serverImpl{
		mux:  http.NewServeMux(),
		srv:  &http.Server{Addr: ":3000"},
		page: NewPage(""),
	}
}

func (s *serverImpl) Start() error {
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

func (s *serverImpl) Stop() error {
	if s.srv != nil {
		err := s.srv.Close()
		if err != nil {
			return err
		}
		s.wg.Wait() // Wait for the server to shutdown
	}
	return nil
}

func (s *serverImpl) UpdateTitle(title string) {
	s.page.Title = title
}
