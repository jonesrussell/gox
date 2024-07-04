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
	HTML    []byte
	updater *WebsiteUpdater
}

// Modify NewPage to accept a FileReader, a WebsiteUpdater, and a filename as arguments
func NewPage(
	title string,
	body string,
	fr utils.FileReader,
	updater *WebsiteUpdater,
	filename string,
) *Page {
	html, err := fr.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	return &Page{
		Title:   title,
		Body:    body,
		HTML:    html,
		updater: updater,
	}
}

func (p *Page) Render() ([]byte, error) {
	doc, err := html.Parse(bytes.NewReader(p.HTML))
	if err != nil {
		return nil, err
	}

	p.updater.ChangeTitle(doc, p.Title)

	// Unescape the body before updating it
	unescapedBody := html.UnescapeString(p.Body)
	p.updater.ChangeBody(doc, unescapedBody)

	var buf bytes.Buffer
	err = html.Render(&buf, doc)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (p *Page) GetHTML() string {
	rendered, err := p.Render()
	if err != nil {
		log.Println("Error rendering page:", err)
		return ""
	}
	return string(rendered)
}

func (p *Page) SetTitle(content string) {
	p.Title = content
	// Add any additional logic here.
}

func (p *Page) SetBody(content string) {
	p.Title = content
	// Add any additional logic here.
}
