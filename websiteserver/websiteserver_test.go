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

func Test_websiteServerImpl_GetAddress(t *testing.T) {
	tests := []struct {
		name string
		addr string
		want string
	}{
		{
			name: "Test GetAddress",
			addr: "127.0.0.1:3000",
			want: "127.0.0.1:3000",
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
			if got := s.GetAddress(); got != tt.want {
				t.Errorf("websiteServerImpl.GetAddress() = %v, want %v", got, tt.want)
			}
		})
	}
}
