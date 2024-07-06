package webserver

import (
	"html/template"
	"jonesrussell/gocreate/logger"
	"jonesrussell/gocreate/utils"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
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
		logger   logger.LoggerInterface
	}{
		{
			name:     "Test with valid title and body",
			title:    testTitle,
			body:     template.HTML(testBody),
			filename: testFilename,
			wantPage: &Page{
				title:   testTitle,
				body:    testBody,
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
				title:   "",
				body:    testBody,
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
				title:   testTitle,
				body:    "",
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
				title:   "",
				body:    "",
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
				title:   testTitle,
				body:    "<p>This is a <strong>test</strong> paragraph.</p>",
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
				title:   "Test & Title <with> Special \"Characters\"",
				body:    testBody,
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
				title:   testTitle,
				body:    testBody,
				HTML:    []byte(mockHTML), // Assuming MockFileReader always returns the same content
				updater: NewWebsiteUpdater(logInstance),
				logger:  logInstance,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fr := utils.MockFileReader{}
			updater := NewWebsiteUpdater(logInstance)
			got := NewPage(tt.title, tt.body, fr, updater, tt.filename, tt.logger)

			assert.Equal(t, tt.wantPage.title, got.title)
			assert.Equal(t, tt.wantPage.body, got.body)
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
				title:   tt.title,
				body:    template.HTML(tt.body),
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
	t.Run("Test with valid title and body", func(t *testing.T) {
		// Create a mock updater
		mockUpdater := new(MockWebsiteUpdater)
		mockUpdater.On("UpdateWebsite", mock.Anything, mock.Anything, mock.Anything).Return("<html><head><title>Test Title</title></head><body><p>Test Body</p></body></html>", nil)

		// Create a mock file reader
		mockFileReader := new(MockFileReader)
		mockFileReader.On("ReadFile", mock.Anything).Return([]byte("mock file content"), nil)

		// Create a mock logger
		mockLogger := new(logger.MockLogger)
		mockLogger.On("Debug", mock.Anything, mock.Anything).Return()
		mockLogger.On("Error", mock.Anything, mock.AnythingOfType("error"), mock.Anything).Return()

		// Initialize the Page struct
		page := NewPage(
			"Test Title",
			template.HTML("<p>Test Body</p>"),
			mockFileReader,
			mockUpdater,
			"test_template.html",
			mockLogger,
		)

		// Call GetHTML
		html := page.GetHTML()

		// Assert the result
		assert.Contains(t, html, "Test Title")
		assert.Contains(t, html, "<p>Test Body</p>")

		// Verify that methods were called
		mockUpdater.AssertCalled(t, "UpdateWebsite", mock.Anything, mock.Anything, mock.Anything)
		mockLogger.AssertCalled(t, "Debug", mock.Anything, mock.Anything)
	})
}
