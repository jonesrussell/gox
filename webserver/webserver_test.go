package webserver

import (
	"html/template"
	"jonesrussell/gocreate/logger"
	"net/http"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

// initLogger creates a new logger and handles any potential errors
func initLogger() logger.LoggerInterface {
	log, err := logger.NewLogger("/tmp/gocreate-tests.log")
	if err != nil {
		panic("Failed to initialize logger: " + err.Error())
	}
	return log
}

// Create a new Logger and WebsiteUpdater once for all tests
var (
	logInstance = initLogger()
)

func TestNewServer(t *testing.T) {
	// Create a new mock server
	server := NewMockServer(nil)

	// Assert that the server was created
	assert.NotNil(t, server)

	// Test specific methods
	assert.NotNil(t, server.Logger())
	assert.Equal(t, "http://127.0.0.1:3000", server.GetURL())

	// Test that the server has a non-nil page
	assert.NotNil(t, server.(*MockServer).page)

	// Test default values
	assert.Equal(t, "Mock Title", server.(*MockServer).page.Title)
	assert.Equal(t, template.HTML("<h1>Mock Body</h1>"), server.(*MockServer).page.Body)
}

func TestMockServer_StartStop(t *testing.T) {
	server := NewMockServer(nil)

	// Test Start
	err := server.Start()
	assert.NoError(t, err)

	// Test Stop
	err = server.Stop()
	assert.NoError(t, err)
}

// func TestMockServer_UpdateTitleAndBody(t *testing.T) {
// 	server := NewMockServer(nil)

// 	// Test UpdateTitle
// 	server.UpdateTitle("New Title")
// 	assert.Contains(t, server.GetHTML(), "New Title")

// 	// Test UpdateBody
// 	server.UpdateBody("<p>New Body</p>")
// 	assert.Contains(t, server.GetHTML(), "<p>New Body</p>")
// }

func Test_webServer_GetURL(t *testing.T) {
	tests := []struct {
		name string
		addr string
		want string
	}{
		{
			name: "Test GetURL with 127.0.0.1:3000",
			addr: "127.0.0.1:3000",
			want: "http://127.0.0.1:3000",
		},
		{
			name: "Test GetURL with localhost:3000",
			addr: "localhost:3000",
			want: "http://localhost:3000",
		},
		{
			name: "Test GetURL with 0.0.0.0:3000",
			addr: "0.0.0.0:3000",
			want: "http://127.0.0.1:3000",
		},
		{
			name: "Test GetURL with :3000",
			addr: ":3000",
			want: "http://127.0.0.1:3000",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &webServer{
				logger: logInstance,
				mux:    http.NewServeMux(),
				srv:    &http.Server{Addr: tt.addr},
				wg:     sync.WaitGroup{},
				page:   &Page{},
			}
			if got := s.GetURL(); got != tt.want {
				t.Errorf("webServer.GetURL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMockServer_Logger(t *testing.T) {
	server := NewMockServer(nil)
	assert.NotNil(t, server.Logger())
}
