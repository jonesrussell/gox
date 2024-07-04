package webserver

import (
	"jonesrussell/gocreate/debug"

	"golang.org/x/net/html"
)

type WebsiteUpdater struct {
	debugger debug.Debugger
}

func NewWebsiteUpdater(debugger debug.Debugger) *WebsiteUpdater {
	return &WebsiteUpdater{
		debugger: debugger,
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
