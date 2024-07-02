package websiteserver

import (
	"bytes"
	"jonesrussell/gocreate/utils"
	"log"

	"golang.org/x/net/html"
)

type Page struct {
	Title   string
	Body    string
	content []byte
	updater *WebsiteUpdater
}

// Modify NewPage to accept a FileReader and a WebsiteUpdater as arguments
func NewPage(title string, body string, fr utils.FileReader, updater *WebsiteUpdater) *Page {
	content, err := fr.ReadFile("static/index.html")
	if err != nil {
		log.Fatal(err)
	}

	return &Page{
		Title:   title,
		Body:    body,
		content: content,
		updater: updater,
	}
}

func (p *Page) Render() ([]byte, error) {
	doc, err := html.Parse(bytes.NewReader(p.content))
	if err != nil {
		return nil, err
	}

	if p.Title != "" {
		p.updater.ChangeTitle(doc, p.Title)
	}

	if p.Body != "" {
		p.updater.ChangeBody(doc, p.Body)
	}

	var buf bytes.Buffer
	err = html.Render(&buf, doc)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
