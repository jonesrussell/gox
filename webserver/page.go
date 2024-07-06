package webserver

import (
	"bytes"
	"html/template"
	"jonesrussell/gocreate/logger"
	"jonesrussell/gocreate/utils"

	"golang.org/x/net/html"
)

type Page struct {
	title        string
	body         template.HTML
	HTML         []byte
	fileReader   utils.FileReaderInterface
	updater      PageUpdaterInterface
	templatePath string
	updateChan   chan struct{}
	logger       logger.LoggerInterface
}

func NewPage(
	title string,
	body template.HTML,
	fileReader utils.FileReaderInterface,
	updater PageUpdaterInterface,
	templatePath string,
	logger logger.LoggerInterface,
) *Page {
	page := &Page{
		title:        title,
		body:         body,
		fileReader:   fileReader,
		updater:      updater,
		templatePath: templatePath,
		updateChan:   make(chan struct{}, 1),
		logger:       logger,
	}

	// Update the website using the updater
	html, err := updater.UpdatePage(title, string(body), templatePath)
	if err != nil {
		logger.Error("Error updating website:", err)
		return page
	}

	page.HTML = []byte(html)

	return page
}

func (p *Page) Render() ([]byte, error) {
	doc, err := html.Parse(bytes.NewReader(p.HTML))
	if err != nil {
		return nil, err
	}

	p.updater.ChangeTitle(doc, p.title)

	p.updater.ChangeBody(doc, string(p.body))

	var buf bytes.Buffer
	err = html.Render(&buf, doc)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
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
	select {
	case p.updateChan <- struct{}{}:
		p.logger.Debug("Update notification sent successfully")
	default:
		p.logger.Debug("Update channel is full, notification not sent")
	}
}
