package websiteserver

import (
	"jonesrussell/gocreate/debug"
	"net/http"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewServer(t *testing.T) {
	// Create a new mock server
	server := NewMockServer(&Page{})

	// Assert that the server was created
	assert.NotNil(t, server)
}

func Test_websiteServerImpl_GetURL(t *testing.T) {
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
			s := &websiteServerImpl{
				debugger: &debug.LogDebugger{},
				mux:      http.NewServeMux(),
				srv:      &http.Server{Addr: tt.addr},
				wg:       sync.WaitGroup{},
				page:     &Page{},
			}
			if got := s.GetURL(); got != tt.want {
				t.Errorf("websiteServerImpl.GetURL() = %v, want %v", got, tt.want)
			}
		})
	}
}
