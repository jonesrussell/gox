package httpserver

import (
	"bytes"
	"net/http"
	"time"
)

func handleRequest(w http.ResponseWriter, r *http.Request, page *Page) error {
	content, err := page.Render()
	if err != nil {
		return err
	}

	http.ServeContent(w, r, "index.html", time.Now(), bytes.NewReader(content))
	return nil
}
