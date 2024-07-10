package webserver

import (
	"bytes"
	"html/template"
	"io"
	"jonesrussell/gocreate/logger"
	"os"

	"jonesrussell/gocreate/htmlservice"

	"github.com/yosssi/gohtml"
	"golang.org/x/net/html"
)

type Page struct {
	title        string
	body         template.HTML
	HTML         []byte
	updater      PageUpdaterInterface
	templatePath string
	updateChan   chan struct{}
	logger       logger.LoggerInterface
	htmlService  *htmlservice.HTMLService
}

func NewPage(
	title string,
	body template.HTML,
	updater PageUpdaterInterface,
	templatePath string,
	logger logger.LoggerInterface,
) (*Page, error) {
	page := &Page{
		title:        title,
		body:         body,
		updater:      updater,
		templatePath: templatePath,
		updateChan:   make(chan struct{}, 1),
		logger:       logger,
		htmlService:  htmlservice.NewHTMLService(),
	}

	tmpl, err := os.ReadFile(templatePath)
	if err != nil {
		logger.Error("Error reading template:", err)
		return nil, err
	}

	doc, err := page.htmlService.ParseHTML(tmpl) // Use HTMLService to parse HTML
	if err != nil {
		logger.Error("Error parsing template:", err)
		return nil, err
	}

	// If title is empty, use the default title from the template
	if title == "" {
		var f func(*html.Node)
		f = func(n *html.Node) {
			if n.Type == html.ElementNode && n.Data == "title" {
				page.title = n.FirstChild.Data
			}
			for c := n.FirstChild; c != nil; c = c.NextSibling {
				f(c)
			}
		}
		f(doc)
	}

	// If body is empty, use the default body from the template
	if body == "" {
		var f func(*html.Node)
		f = func(n *html.Node) {
			if n.Type == html.ElementNode && n.Data == "body" {
				var buf bytes.Buffer
				w := io.Writer(&buf)
				html.Render(w, n)
				page.body = template.HTML(gohtml.Format(buf.String()))
			}
			for c := n.FirstChild; c != nil; c = c.NextSibling {
				f(c)
			}
		}
		f(doc)
	}

	// Update the website using the updater
	html, err := updater.UpdatePage(page.title, string(page.body), templatePath)
	if err != nil {
		logger.Error("Error updating website:", err)
		return nil, err
	}

	page.HTML = []byte(html)

	return page, nil
}

func (p *Page) Render() ([]byte, error) {
	doc, err := p.htmlService.ParseHTML(p.HTML) // Use HTMLService to parse HTML
	if err != nil {
		p.logger.Error("Error parsing HTML: ", err)
		return nil, err
	}

	p.updater.ChangeTitle(doc, p.title)

	p.updater.ChangeBody(doc, string(p.body))

	html, err := p.htmlService.RenderHTML(doc) // Use HTMLService to render HTML
	if err != nil {
		p.logger.Error("Error rendering updated HTML: ", err)
		return nil, err
	}

	return html, nil
}

func (p *Page) SetTitle(title string) {
	p.logger.Debug("SetTitle called with title: " + title)
	p.title = title
	p.HTML = []byte{} // Clear cached HTML
	p.logger.Debug("Title set successfully, cached HTML cleared")
	p.notifyUpdate()
	p.logger.Debug("notifyUpdate called after setting title")
}

func (p *Page) SetBody(body string) {
	p.logger.Debug("SetBody called with body: " + body)
	p.body = template.HTML(body)
	p.HTML = []byte{} // Clear cached HTML
	p.logger.Debug("Body set successfully, cached HTML cleared")
	p.notifyUpdate()
	p.logger.Debug("notifyUpdate called after setting body")
}

func (p *Page) GetTitle() string {
	p.logger.Debug("GetTitle called")
	return p.title
}

func (p *Page) GetBody() string {
	p.logger.Debug("GetBody called")
	return string(p.body)
}

func (p *Page) GetHTML() string {
	p.logger.Debug("GetHTML called")
	if len(p.HTML) == 0 {
		p.logger.Debug("Cached HTML is empty, generating new HTML")
		html, err := p.updater.UpdatePage(p.title, string(p.body), p.templatePath)
		if err != nil {
			p.logger.Error("Error updating website: ", err)
			return ""
		}
		p.HTML = []byte(html)
		p.logger.Debug("New HTML generated and cached successfully")
	} else {
		p.logger.Debug("Returning cached HTML")
	}
	return string(p.HTML)
}

func (p *Page) notifyUpdate() {
	p.logger.Debug("notifyUpdate called")
	p.updateChan <- struct{}{} // Block until the message can be sent
	p.logger.Debug("Update notification sent successfully")
}
