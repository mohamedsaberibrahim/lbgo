package internals

import (
	"net/http"
)

type ServerInterface interface {
	// Get the server's name
	GetName() string
	// Get the server's address
	GetAddress() string
	// Check the server's health
	CheckHealth() bool
	// Serve the request
	ServeRequest(http.ResponseWriter, *http.Request) bool
}
