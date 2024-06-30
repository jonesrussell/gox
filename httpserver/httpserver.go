package httpserver

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"golang.org/x/net/html"
)

type Page struct {
	Title string
}

type Server struct {
	mux  *http.ServeMux
	srv  *http.Server
	wg   sync.WaitGroup
	page Page
}

func NewServer() *Server {
	return &Server{
		mux: http.NewServeMux(),
		srv: &http.Server{Addr: ":3000"},
	}
}

func (s *Server) Start() {
	s.mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handleRequest(w, r, s.page)
	})

	s.srv.Handler = s.mux

	s.wg.Add(1)
	go func() {
		defer s.wg.Done()

		log.Println("Listening on :3000...")
		if err := s.srv.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()
}

func (s *Server) Stop() {
	if s.srv != nil {
		if err := s.srv.Close(); err != nil {
			log.Fatal(err)
		}
		s.wg.Wait() // Wait for the server to shutdown
	}
}

func changeTitle(n *html.Node, newTitle string) {
	if n.Type == html.ElementNode && n.Data == "title" {
		if n.FirstChild != nil {
			n.FirstChild.Data = newTitle
		}
		return
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		changeTitle(c, newTitle)
	}
}

func (s *Server) UpdateTitle(title string) {
	s.page.Title = title
}

func readFile(path string) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return io.ReadAll(file)
}

func handleRequest(w http.ResponseWriter, r *http.Request, page Page) {
	content, err := readFile("static/index.html")
	if err != nil {
		log.Fatal(err)
	}

	doc, err := html.Parse(bytes.NewReader(content))
	if err != nil {
		log.Fatal(err)
	}

	if page.Title != "" {
		changeTitle(doc, page.Title)
	}

	var buf bytes.Buffer
	err = html.Render(&buf, doc)
	if err != nil {
		log.Fatal(err)
	}

	http.ServeContent(w, r, "index.html", time.Now(), bytes.NewReader(buf.Bytes()))
}
