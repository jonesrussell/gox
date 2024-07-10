package webserver

import (
	"html/template"
	"jonesrussell/gocreate/htmlservice"
	"jonesrussell/gocreate/logger"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/net/html"
)

const (
	testTitle    = "Test Title"
	testBody     = "Test Body"
	testFilename = "../static/index.html"
	mockHTML     = "<html><head><title>Mock Title</title></head><body>Mock Body</body></html>"
)

func TestPage_NewPage(t *testing.T) {
	tests := []struct {
		name         string
		title        string
		body         string
		expectedHTML string
	}{
		{
			name:         "Test with valid title and body",
			title:        testTitle,
			body:         testBody,
			expectedHTML: "<html><head><title>Test Title</title></head><body>Test Body</body></html>",
		},
		// Add more test cases as needed
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockUpdater := &MockPageUpdater{}
			mockLogger := new(logger.MockLogger)
			mockHTMLService := new(htmlservice.MockHTMLService) // Create a mock HTMLService

			// Set up the mock updater to return the expected HTML
			mockUpdater.On("UpdatePage", tt.title, tt.body, testFilename).Return(tt.expectedHTML, nil)

			// Set up the mock HTMLService to return a nil doc and no error
			mockHTMLService.On("ParseHTML", mock.Anything).Return(&html.Node{}, nil)

			page, err := NewPage(tt.title, template.HTML(tt.body), mockUpdater, testFilename, mockLogger, mockHTMLService) // Pass the mock HTMLService to NewPage
			if err != nil {
				t.Fatal("Error creating page:", err)
			}

			assert.Equal(t, tt.title, page.Template.Title)
			assert.Equal(t, template.HTML(tt.body), page.Template.Body)
			assert.Equal(t, tt.expectedHTML, string(page.Template.HTML))

			mockUpdater.AssertExpectations(t)
			mockLogger.AssertExpectations(t)
			mockHTMLService.AssertExpectations(t) // Verify that the HTMLService methods were called as expected
		})
	}
}
