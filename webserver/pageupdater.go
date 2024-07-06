package webserver

import (
	"bytes"
	"os"
	"strings"

	"jonesrussell/gocreate/logger"

	"golang.org/x/net/html"
)

type PageUpdaterInterface interface {
	ChangeTitle(doc *html.Node, title string)
	ChangeBody(doc *html.Node, body string)
	UpdatePage(title string, body string, templatePath string) (string, error)
}

type PageUpdater struct {
	logger logger.LoggerInterface
}

func NewPageUpdater(logger logger.LoggerInterface) *PageUpdater {
	return &PageUpdater{
		logger: logger,
	}
}

func (p *PageUpdater) ChangeTitle(n *html.Node, newTitle string) {
	if n.Type == html.ElementNode && n.Data == "title" {
		if n.FirstChild != nil {
			n.FirstChild.Data = newTitle
		}
		return
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		p.ChangeTitle(c, newTitle)
	}
}

func (p *PageUpdater) ChangeBody(doc *html.Node, newBody string) {
	var body *html.Node
	var findBody func(*html.Node)
	findBody = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "body" {
			body = n
			return
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			findBody(c)
		}
	}
	findBody(doc)

	if body != nil {
		// Clear existing body content
		for body.FirstChild != nil {
			body.RemoveChild(body.FirstChild)
		}

		// Parse the new body content
		nodes, err := html.ParseFragment(strings.NewReader(newBody), body)
		if err != nil {
			p.logger.Error("Failed to parse new body content", err)
			return
		}

		// Append new nodes to the body
		for _, n := range nodes {
			body.AppendChild(n)
		}
	}
}

func (p *PageUpdater) UpdatePage(title, body, templatePath string) (string, error) {
	p.logger.Debug("UpdatePage called with title: " + title + " and body: " + body)

	// Read the template file
	content, err := p.readTemplateFile(templatePath)
	if err != nil {
		return "", err
	}

	// Parse the HTML
	doc, err := p.parseHTML(content)
	if err != nil {
		return "", err
	}

	// Update the title and body
	p.updateTitleAndBody(doc, title, body)

	// Render the updated HTML
	htmlString, err := p.renderHTML(doc)
	if err != nil {
		return "", err
	}

	p.logger.Debug("Page updated successfully")
	return htmlString, nil
}

func (p *PageUpdater) readTemplateFile(templatePath string) ([]byte, error) {
	content, err := os.ReadFile(templatePath)
	if err != nil {
		p.logger.Error("Error reading template file: ", err)
	}
	return content, err
}

func (p *PageUpdater) parseHTML(content []byte) (*html.Node, error) {
	doc, err := html.Parse(bytes.NewReader(content))
	if err != nil {
		p.logger.Error("Error parsing HTML: ", err)
	}
	return doc, err
}

func (p *PageUpdater) updateTitleAndBody(doc *html.Node, title, body string) {
	p.ChangeTitle(doc, title)
	p.ChangeBody(doc, body)
}

func (p *PageUpdater) renderHTML(doc *html.Node) (string, error) {
	var buf bytes.Buffer
	err := html.Render(&buf, doc)
	if err != nil {
		p.logger.Error("Error rendering HTML: ", err)
		return "", err
	}
	return buf.String(), nil
}
