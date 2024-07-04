package websiteserver

import (
	"bytes"
	"net/http"
	"time"
)

func handleRequest(w http.ResponseWriter, r *http.Request, page *Page) error {
	content := page.GetHTML()

	http.ServeContent(
		w,
		r,
		"index.html",
		time.Now(),
		bytes.NewReader([]byte(content)),
	)

	return nil
}
