package webserver

import (
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

func TestMockServer_StartStop(t *testing.T) {
	server := NewMockServer(nil)

	// Test Start
	err := server.Start()
	assert.NoError(t, err)

	// Test Stop
	err = server.Stop()
	assert.NoError(t, err)
}

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
			want: "http://localhost:3000",
		},
		{
			name: "Test GetURL with :3000",
			addr: ":3000",
			want: "http://localhost:3000",
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
