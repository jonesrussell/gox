package webserver

import (
	"bytes"
	"jonesrussell/gocreate/logger"
	"os"
	"strings"

	"golang.org/x/net/html"
)

type WebsiteUpdaterInterface interface {
	ChangeTitle(doc *html.Node, title string)
	ChangeBody(doc *html.Node, body string)
	UpdateWebsite(title string, body string, templatePath string) (string, error)
}

type WebsiteUpdater struct {
	logger logger.LoggerInterface
}

func NewWebsiteUpdater(logger logger.LoggerInterface) *WebsiteUpdater {
	return &WebsiteUpdater{
		logger: logger,
	}
}

func (w *WebsiteUpdater) ChangeTitle(n *html.Node, newTitle string) {
	if n.Type == html.ElementNode && n.Data == "title" {
		if n.FirstChild != nil {
			n.FirstChild.Data = newTitle
		}
		return
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		w.ChangeTitle(c, newTitle)
	}
}

func (wu *WebsiteUpdater) ChangeBody(doc *html.Node, newBody string) {
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
			wu.logger.Error("Failed to parse new body content", err)
			return
		}

		// Append new nodes to the body
		for _, n := range nodes {
			body.AppendChild(n)
		}
	}
}

func (w *WebsiteUpdater) UpdateWebsite(title, body, templatePath string) (string, error) {
	w.logger.Debug("UpdateWebsite called with title: " + title + " and body: " + body)

	// Read the template file
	content, err := os.ReadFile(templatePath)
	if err != nil {
		w.logger.Error("Error reading template file: ", err)
		return "", err
	}

	// Parse the HTML
	doc, err := html.Parse(bytes.NewReader(content))
	if err != nil {
		w.logger.Error("Error parsing HTML: ", err)
		return "", err
	}

	// Update the title and body
	w.ChangeTitle(doc, title)
	w.ChangeBody(doc, body)

	// Render the updated HTML
	var buf bytes.Buffer
	err = html.Render(&buf, doc)
	if err != nil {
		w.logger.Error("Error rendering HTML: ", err)
		return "", err
	}

	w.logger.Debug("Website updated successfully")
	return buf.String(), nil
}
