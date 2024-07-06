package webserver

import (
	"html/template"
	"jonesrussell/gocreate/logger"
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
	mockFileReader := &MockFileReader{}
	mockUpdater := &MockPageUpdater{}
	mockLogger := new(logger.MockLogger)

	title := "Test Title"
	body := template.HTML("Test Body")
	templatePath := "template.html"

	expectedHTML := "<html><head><title>Test Title</title></head><body>Test Body</body></html>"

	// Set up the mock updater to return the expected HTML
	// This needs to be done BEFORE calling NewPage
	mockUpdater.On("UpdatePage", title, string(body), templatePath).Return(expectedHTML, nil)

	page := NewPage(title, body, mockFileReader, mockUpdater, templatePath, mockLogger)

	assert.Equal(t, title, page.title)
	assert.Equal(t, body, page.body)
	assert.Equal(t, expectedHTML, string(page.HTML))

	mockUpdater.AssertExpectations(t)
	mockLogger.AssertExpectations(t)
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
				updater: NewPageUpdater(logInstance),
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
		mockUpdater := new(MockPageUpdater)
		mockUpdater.On("UpdatePage", mock.Anything, mock.Anything, mock.Anything).Return("<html><head><title>Test Title</title></head><body><p>Test Body</p></body></html>", nil)

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
		mockUpdater.AssertCalled(t, "UpdatePage", mock.Anything, mock.Anything, mock.Anything)
		mockLogger.AssertCalled(t, "Debug", mock.Anything, mock.Anything)
	})
}
