package webserver

import (
	"html/template"
	"jonesrussell/gocreate/utils"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	testTitle    = "Test Title"
	testBody     = "Test Body"
	testFilename = "../static/index.html"
	mockHTML     = "<html><head><title>Mock Title</title></head><body>Mock Body</body></html>"
)

func TestPage_NewPage(t *testing.T) {
	tests := []struct {
		name     string
		title    string
		body     template.HTML
		filename string
		wantPage *Page
	}{
		{
			name:     "Test with valid title and body",
			title:    testTitle,
			body:     template.HTML(testBody),
			filename: testFilename,
			wantPage: &Page{
				Title:   testTitle,
				Body:    testBody,
				HTML:    []byte(mockHTML),
				updater: NewWebsiteUpdater(logInstance),
			},
		},
		{
			name:     "Test with empty title",
			title:    "",
			body:     template.HTML(testBody),
			filename: testFilename,
			wantPage: &Page{
				Title:   "",
				Body:    testBody,
				HTML:    []byte(mockHTML),
				updater: NewWebsiteUpdater(logInstance),
			},
		},
		{
			name:     "Test with empty body",
			title:    testTitle,
			body:     template.HTML(""),
			filename: testFilename,
			wantPage: &Page{
				Title:   testTitle,
				Body:    "",
				HTML:    []byte(mockHTML),
				updater: NewWebsiteUpdater(logInstance),
			},
		},
		{
			name:     "Test with empty title and body",
			title:    "",
			body:     template.HTML(""),
			filename: testFilename,
			wantPage: &Page{
				Title:   "",
				Body:    "",
				HTML:    []byte(mockHTML),
				updater: NewWebsiteUpdater(logInstance),
			},
		},
		{
			name:     "Test with HTML in body",
			title:    testTitle,
			body:     template.HTML("<p>This is a <strong>test</strong> paragraph.</p>"),
			filename: testFilename,
			wantPage: &Page{
				Title:   testTitle,
				Body:    "<p>This is a <strong>test</strong> paragraph.</p>",
				HTML:    []byte(mockHTML),
				updater: NewWebsiteUpdater(logInstance),
			},
		},
		{
			name:     "Test with special characters in title",
			title:    "Test & Title <with> Special \"Characters\"",
			body:     template.HTML(testBody),
			filename: testFilename,
			wantPage: &Page{
				Title:   "Test & Title <with> Special \"Characters\"",
				Body:    testBody,
				HTML:    []byte(mockHTML),
				updater: NewWebsiteUpdater(logInstance),
			},
		},
		{
			name:     "Test with different filename",
			title:    testTitle,
			body:     template.HTML(testBody),
			filename: "../static/different.html",
			wantPage: &Page{
				Title:   testTitle,
				Body:    testBody,
				HTML:    []byte(mockHTML), // Assuming MockFileReader always returns the same content
				updater: NewWebsiteUpdater(logInstance),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fr := utils.MockFileReader{}
			updater := NewWebsiteUpdater(logInstance)
			got := NewPage(tt.title, tt.body, fr, updater, tt.filename)

			assert.Equal(t, tt.wantPage.Title, got.Title)
			assert.Equal(t, tt.wantPage.Body, got.Body)
			assert.Equal(t, tt.wantPage.HTML, got.HTML)
			assert.Equal(t, tt.wantPage.updater, got.updater)
		})
	}
}

func TestPage_Render(t *testing.T) {
	tests := []struct {
		name    string
		title   string
		body    string
		want    []byte
		wantErr bool
	}{
		{
			name:    "Test with valid title and body",
			title:   testTitle,
			body:    testBody,
			want:    []byte("<html><head><title>Test Title</title></head><body>Test Body</body></html>"),
			wantErr: false,
		},
		{
			name:    "Test with empty title",
			title:   "",
			body:    testBody,
			want:    []byte("<html><head><title></title></head><body>Test Body</body></html>"),
			wantErr: false,
		},
		{
			name:    "Test with empty body",
			title:   testTitle,
			body:    "",
			want:    []byte("<html><head><title>Test Title</title></head><body></body></html>"),
			wantErr: false,
		},
		{
			name:    "Test with empty title and body",
			title:   "",
			body:    "",
			want:    []byte("<html><head><title></title></head><body></body></html>"),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Page{
				Title:   tt.title,
				Body:    template.HTML(tt.body),
				HTML:    []byte(mockHTML),
				updater: NewWebsiteUpdater(logInstance),
			}
			got, err := p.Render()
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func TestPage_GetHTML(t *testing.T) {
	tests := []struct {
		name  string
		title string
		body  string
		html  []byte
		want  string
	}{
		{
			name:  "Test with valid title and body",
			title: testTitle,
			body:  testBody,
			html:  []byte(mockHTML),
			want:  "<html><head><title>Test Title</title></head><body>Test Body</body></html>",
		},
		{
			name:  "Test with empty title and body",
			title: "",
			body:  "",
			html:  []byte("<html><head><title></title></head><body></body></html>"),
			want:  "<html><head><title></title></head><body></body></html>",
		},
		// Add more test cases as needed.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Page{
				Title:   tt.title,
				Body:    template.HTML(tt.body),
				HTML:    tt.html,
				updater: NewWebsiteUpdater(logInstance),
			}
			got := p.GetHTML()
			assert.Equal(t, tt.want, got)
		})
	}
}
