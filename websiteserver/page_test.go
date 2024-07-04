package websiteserver

import (
	"bytes"
	"jonesrussell/gocreate/debug"
	"jonesrussell/gocreate/utils"
	"testing"
)

func TestPage_NewPage(t *testing.T) {
	type args struct {
		title    string
		body     string
		fr       utils.FileReader
		updater  *WebsiteUpdater
		filename string
	}
	tests := []struct {
		name string
		args args
		want *Page
	}{
		{
			name: "Test with valid title and body",
			args: args{
				title:    "Test Title",
				body:     "Test Body",
				fr:       utils.MockFileReader{},                    // using the MockFileReader struct
				updater:  NewWebsiteUpdater(debug.NewLogDebugger()), // assuming you have a constructor
				filename: "../static/index.html",
			},
			want: &Page{
				Title:   "Test Title",
				Body:    "Test Body",
				HTML:    []byte("<html><head><title>Mock Title</title></head><body>Mock Body</body></html>"), // this should match the content returned by MockFileReader.ReadFile
				updater: NewWebsiteUpdater(debug.NewLogDebugger()),                                           // assuming you have a constructor
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewPage(tt.args.title, tt.args.body, tt.args.fr, tt.args.updater, tt.args.filename)
			if got.Title != tt.want.Title || got.Body != tt.want.Body || !bytes.Equal(got.HTML, tt.want.HTML) {
				t.Errorf("NewPage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPage_Render(t *testing.T) {
	type fields struct {
		Title   string
		Body    string
		HTML    []byte
		updater *WebsiteUpdater
	}
	tests := []struct {
		name    string
		fields  fields
		want    []byte
		wantErr bool
	}{
		{
			name: "Test with valid title and body",
			fields: fields{
				Title:   "Test Title",
				Body:    "Test Body",
				HTML:    []byte("<html><head><title>Mock Title</title></head><body>Mock Body</body></html>"), // this should match the content returned by MockFileReader.ReadFile
				updater: NewWebsiteUpdater(debug.NewLogDebugger()),                                           // assuming you have a constructor
			},
			want:    []byte("<html><head><title>Test Title</title></head><body>Test Body</body></html>"),
			wantErr: false,
		},
		{
			name: "Test with empty title",
			fields: fields{
				Title:   "",
				Body:    "Test Body",
				HTML:    []byte("<html><head><title>Mock Title</title></head><body>Mock Body</body></html>"), // this should match the content returned by MockFileReader.ReadFile
				updater: NewWebsiteUpdater(debug.NewLogDebugger()),                                           // assuming you have a constructor
			},
			want:    []byte("<html><head><title></title></head><body>Test Body</body></html>"),
			wantErr: false,
		},
		{
			name: "Test with empty body",
			fields: fields{
				Title:   "Test Title",
				Body:    "",
				HTML:    []byte("<html><head><title>Mock Title</title></head><body>Mock Body</body></html>"), // this should match the content returned by MockFileReader.ReadFile
				updater: NewWebsiteUpdater(debug.NewLogDebugger()),                                           // assuming you have a constructor
			},
			want:    []byte("<html><head><title>Test Title</title></head><body></body></html>"),
			wantErr: false,
		},
		{
			name: "Test with empty title and body",
			fields: fields{
				Title:   "",
				Body:    "",
				HTML:    []byte("<html><head><title>Mock Title</title></head><body>Mock Body</body></html>"), // this should match the content returned by MockFileReader.ReadFile
				updater: NewWebsiteUpdater(debug.NewLogDebugger()),                                           // assuming you have a constructor
			},
			want:    []byte("<html><head><title></title></head><body></body></html>"),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Page{
				Title:   tt.fields.Title,
				Body:    tt.fields.Body,
				HTML:    tt.fields.HTML,
				updater: tt.fields.updater,
			}
			got, err := p.Render()
			if (err != nil) != tt.wantErr {
				t.Errorf("Page.Render() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !bytes.Equal(got, tt.want) {
				t.Errorf("Page.Render() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPage_GetHTML(t *testing.T) {
	type fields struct {
		Title   string
		Body    string
		HTML    []byte
		updater *WebsiteUpdater
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Test with valid title and body",
			fields: fields{
				Title:   "Test Title",
				Body:    "Test Body",
				HTML:    []byte("<html><head><title>Mock Title</title></head><body>Mock Body</body></html>"), // this should match the content returned by MockFileReader.ReadFile
				updater: NewWebsiteUpdater(debug.NewLogDebugger()),                                           // assuming you have a constructor
			},
			want: "<html><head><title>Test Title</title></head><body>Test Body</body></html>",
		},
		// Add more test cases as needed.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Page{
				Title:   tt.fields.Title,
				Body:    tt.fields.Body,
				HTML:    tt.fields.HTML,
				updater: tt.fields.updater,
			}
			if got := p.GetHTML(); got != tt.want {
				t.Errorf("Page.GetHTML() = %v, want %v", got, tt.want)
			}
		})
	}
}
