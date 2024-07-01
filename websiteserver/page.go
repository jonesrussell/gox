package websiteserver

import (
	"bytes"
	"jonesrussell/gocreate/utils"
	"log"

	"golang.org/x/net/html"
)

type Page struct {
	Title   string
	content []byte
}

// Modify NewPage to accept a FileReader as an argument
func NewPage(title string, fr utils.FileReader) *Page {
	content, err := fr.ReadFile("static/index.html") // Use the passed FileReader
	if err != nil {
		log.Fatal(err)
	}

	return &Page{
		Title:   title,
		content: content,
	}
}

func (p *Page) Render() ([]byte, error) {
	doc, err := html.Parse(bytes.NewReader(p.content))
	if err != nil {
		return nil, err
	}

	if p.Title != "" {
		changeTitle(doc, p.Title)
	}

	var buf bytes.Buffer
	err = html.Render(&buf, doc)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
