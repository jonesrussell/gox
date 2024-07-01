package websiteserver

import (
	"bytes"
	"log"

	"golang.org/x/net/html"
)

type Page struct {
	Title   string
	content []byte
}

func NewPage(title string) *Page {
	content, err := readFile("static/index.html")
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
