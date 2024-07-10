package webserver

import (
	"bytes"
	"html/template"
	"io"
	"jonesrussell/gocreate/htmlservice"
	"jonesrussell/gocreate/logger"
	"os"

	"github.com/yosssi/gohtml"
	"golang.org/x/net/html"
)

type Template struct {
	Title        string
	Body         template.HTML
	HTML         []byte
	TemplatePath string
}

type Page struct {
	Template    Template
	Updater     PageUpdaterInterface
	UpdateChan  chan struct{}
	Logger      logger.LoggerInterface
	HTMLService htmlservice.HTMLServiceInterface
}

func NewPage(
	title string,
	body template.HTML,
	updater PageUpdaterInterface,
	templatePath string,
	logger logger.LoggerInterface,
	htmlService htmlservice.HTMLServiceInterface,
) (*Page, error) {
	page := &Page{
		Template: Template{
			Title:        title,
			Body:         body,
			TemplatePath: templatePath,
		},
		Updater:     updater,
		UpdateChan:  make(chan struct{}, 1),
		Logger:      logger,
		HTMLService: htmlService,
	}

	tmpl, err := os.ReadFile(page.Template.TemplatePath)
	if err != nil {
		page.Logger.Error("Error reading template:", err)
		return nil, err
	}

	doc, err := page.HTMLService.ParseHTML(tmpl) // Use HTMLService to parse HTML
	if err != nil {
		page.Logger.Error("Error parsing template:", err)
		return nil, err
	}

	// If title is empty, use the default title from the template
	if page.Template.Title == "" {
		var f func(*html.Node)
		f = func(n *html.Node) {
			if n.Type == html.ElementNode && n.Data == "title" {
				page.Template.Title = n.FirstChild.Data
			}
			for c := n.FirstChild; c != nil; c = c.NextSibling {
				f(c)
			}
		}
		f(doc)
	}

	// If body is empty, use the default body from the template
	if string(page.Template.Body) == "" {
		var f func(*html.Node)
		f = func(n *html.Node) {
			if n.Type == html.ElementNode && n.Data == "body" {
				var buf bytes.Buffer
				w := io.Writer(&buf)
				html.Render(w, n)
				page.Template.Body = template.HTML(gohtml.Format(buf.String()))
			}
			for c := n.FirstChild; c != nil; c = c.NextSibling {
				f(c)
			}
		}
		f(doc)
	}

	// Update the website using the updater
	docHTML, err := page.Updater.UpdatePage(page.Template.Title, string(page.Template.Body), page.Template.TemplatePath)
	if err != nil {
		page.Logger.Error("Error updating website:", err)
		return nil, err
	}

	page.Template.HTML = []byte(docHTML)

	return page, nil
}

func (p *Page) Render() ([]byte, error) {
	doc, err := p.HTMLService.ParseHTML(p.Template.HTML) // Use HTMLService to parse HTML
	if err != nil {
		p.Logger.Error("Error parsing HTML: ", err)
		return nil, err
	}

	p.Updater.ChangeTitle(doc, p.Template.Title)

	p.Updater.ChangeBody(doc, string(p.Template.Body))

	docHTML, err := p.HTMLService.RenderHTML(doc) // Use HTMLService to render HTML
	if err != nil {
		p.Logger.Error("Error rendering updated HTML: ", err)
		return nil, err
	}

	return docHTML, nil
}

func (t *Template) SetTitle(title string) {
	t.Title = title
	t.HTML = []byte{} // Clear cached HTML
}

func (t *Template) SetBody(body string) {
	t.Body = template.HTML(body)
	t.HTML = []byte{} // Clear cached HTML
}

func (p *Page) UpdateTitle(title string) {
	p.Logger.Debug("UpdateTitle called with title: " + title)
	p.Template.SetTitle(title)
	p.Logger.Debug("Title set successfully, cached HTML cleared")
	p.notifyUpdate()
	p.Logger.Debug("notifyUpdate called after setting title")
}

func (p *Page) UpdateBody(body string) {
	p.Logger.Debug("UpdateBody called with body: " + body)
	p.Template.SetBody(body)
	p.Logger.Debug("Body set successfully, cached HTML cleared")
	p.notifyUpdate()
	p.Logger.Debug("notifyUpdate called after setting body")
}

func (t *Template) GetTitle() string {
	return t.Title
}

func (t *Template) GetBody() string {
	return string(t.Body)
}

func (p *Page) GetHTML() string {
	p.Logger.Debug("GetHTML called")
	if len(p.Template.HTML) == 0 {
		p.Logger.Debug("Cached HTML is empty, generating new HTML")
		docHTML, err := p.Updater.UpdatePage(p.Template.Title, string(p.Template.Body), p.Template.TemplatePath)
		if err != nil {
			p.Logger.Error("Error updating website: ", err)
			return ""
		}
		p.Template.HTML = []byte(docHTML)
		p.Logger.Debug("New HTML generated and cached successfully")
	} else {
		p.Logger.Debug("Returning cached HTML")
	}
	return string(p.Template.HTML)
}

func (p *Page) notifyUpdate() {
	p.Logger.Debug("notifyUpdate called")
	p.UpdateChan <- struct{}{} // Block until the message can be sent
	p.Logger.Debug("Update notification sent successfully")
}
