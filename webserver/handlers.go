package webserver

import (
	"bytes"
	"fmt"
	"net/http"
	"time"
)

func (s *webServer) handleRequest(w http.ResponseWriter, r *http.Request) error {
	content := s.page.GetHTML()

	http.ServeContent(
		w,
		r,
		"index.html",
		time.Now(),
		bytes.NewReader([]byte(content)),
	)

	return nil
}

func (s *webServer) handleRootRequest(w http.ResponseWriter, r *http.Request) {
	s.logger.Debug(fmt.Sprintf("Received request at %s endpoint", r.RequestURI))
	err := s.handleRequest(w, r)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		s.logger.Error("Error handling request: ", err)
	}
}

// func (s *webServer) handleUpdatesRequest(w http.ResponseWriter, r *http.Request) {
// 	s.logger.Debug("Received request at /updates endpoint")

// 	// Log the details of the request
// 	s.logger.Debug(fmt.Sprintf("Request Method: %s, URL: %s", r.Method, r.URL))

// 	// Use the sseServer to handle the SSE
// 	s.logger.Debug("Handling the SSE request")
// 	s.sseServer.ServeHTTP(w, r)
// 	s.logger.Debug("Finished handling the SSE request")
// }
