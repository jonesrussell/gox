package htmlservice

import (
	"bytes"

	"golang.org/x/net/html"
)

type HTMLService struct{}

func NewHTMLService() *HTMLService {
	return &HTMLService{}
}

func (h *HTMLService) ParseHTML(htmlData []byte) (*html.Node, error) {
	reader := bytes.NewReader(htmlData)
	doc, err := html.Parse(reader)
	if err != nil {
		return nil, err
	}
	return doc, nil
}

func (h *HTMLService) RenderHTML(doc *html.Node) ([]byte, error) {
	var buf bytes.Buffer
	err := html.Render(&buf, doc)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
