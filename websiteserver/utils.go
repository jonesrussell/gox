package websiteserver

import (
	"os"

	"jonesrussell/gocreate/utils"

	"golang.org/x/net/html"
)

var ReadFile utils.FileReader = os.ReadFile

func changeTitle(n *html.Node, newTitle string) {
	if n.Type == html.ElementNode && n.Data == "title" {
		if n.FirstChild != nil {
			n.FirstChild.Data = newTitle
		}
		return
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		changeTitle(c, newTitle)
	}
}
