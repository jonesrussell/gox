package webserver

import (
	"jonesrussell/gocreate/logger"

	"golang.org/x/net/html"
)

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

func (w *WebsiteUpdater) ChangeBody(n *html.Node, newBody string) {
	if n.Type == html.ElementNode && n.Data == "body" {
		// Clear existing children
		for c := n.FirstChild; c != nil; {
			next := c.NextSibling
			n.RemoveChild(c) // RemoveChild will not remove nodes that are not direct children of n
			c = next
		}

		// Append new content as a text node
		newNode := &html.Node{
			Type: html.TextNode,
			Data: newBody,
		}
		n.AppendChild(newNode)

		return
	}

	// Recursively search through children
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		w.ChangeBody(c, newBody)
	}
}
