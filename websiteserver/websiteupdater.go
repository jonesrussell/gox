package websiteserver

import (
	"fmt"
	"jonesrussell/gocreate/debug"
	"strings"

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
	w.PrettyPrintNode(n, "")

	if n.Type == html.ElementNode && n.Data == "body" {
		w.debugger.Debug("Found <body> element!")

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

		w.debugger.Debug("Replaced entire content of <body> with new content.")
		return
	}

	// Recursively search through children
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		w.ChangeBody(c, newBody)
	}
}

func (w *WebsiteUpdater) PrettyPrintNode(node *html.Node, indent string) {
	var nodeTypeStr string
	switch node.Type {
	case html.ErrorNode:
		nodeTypeStr = "ErrorNode"
	case html.DocumentNode:
		nodeTypeStr = "DocumentNode"
	case html.ElementNode:
		nodeTypeStr = "ElementNode"
	case html.TextNode:
		nodeTypeStr = "TextNode"
	case html.CommentNode:
		nodeTypeStr = "CommentNode"
	case html.DoctypeNode:
		nodeTypeStr = "DoctypeNode"
	default:
		nodeTypeStr = "UnknownNode"
	}

	// Simplify the representation of text nodes to avoid printing excessive whitespace
	if node.Type == html.TextNode && strings.TrimSpace(node.Data) == "" {
		w.debugger.Debug(indent + "TextNode: Whitespace")
	} else {
		w.debugger.Debug(fmt.Sprintf("%s%s: %q", indent, nodeTypeStr, node.Data))
	}

	// Recursively print child nodes with increased indentation for readability
	if node.FirstChild != nil {
		w.PrettyPrintNode(node.FirstChild, indent+"  ")
	}
	if node.NextSibling != nil {
		w.PrettyPrintNode(node.NextSibling, indent)
	}
}
