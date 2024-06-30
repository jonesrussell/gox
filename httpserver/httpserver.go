package httpserver

import (
	"log"
	"net/http"
	"sync"
)

type Server struct {
	mux  *http.ServeMux
	srv  *http.Server
	wg   sync.WaitGroup
	page *Page
}

func NewServer() *Server {
	return &Server{
		mux:  http.NewServeMux(),
		srv:  &http.Server{Addr: ":3000"},
		page: NewPage(""),
	}
}

func (s *Server) Start() error {
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

func (s *Server) Stop() error {
	if s.srv != nil {
		err := s.srv.Close()
		if err != nil {
			return err
		}
		s.wg.Wait() // Wait for the server to shutdown
	}
	return nil
}

func (s *Server) UpdateTitle(title string) {
	s.page.Title = title
}
