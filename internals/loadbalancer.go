package internals

import (
	"fmt"
	"net/http"
)

type LoadBalancer struct {
	// The list of servers
	servers []ServerInterface
	// The current server index
	currentServerIndex int
	// Port to listen on
	port string
}

func (lb *LoadBalancer) New(port string) {
	lb.port = port
	lb.currentServerIndex = 0
}

func (lb *LoadBalancer) AddServer(server ServerInterface) {
	lb.servers = append(lb.servers, server)
}

func (lb *LoadBalancer) ServeRequest(rw http.ResponseWriter, req *http.Request) bool {
	// Get the next server
	server := lb.get_next_server()
	// Logging the server name that serving the request
	fmt.Println("Serving request from server: ", server.GetName())
	// Serve the request
	return server.ServeRequest(rw, req)
}

func (lb *LoadBalancer) GetPort() string {
	return lb.port
}

func (lb *LoadBalancer) get_next_server() ServerInterface {
	// Get the next server
	server := lb.servers[lb.currentServerIndex]
	// Iterate on servers till we find a healthy one
	for !server.CheckHealth() {
		lb.currentServerIndex = (lb.currentServerIndex + 1) % len(lb.servers)
		server = lb.servers[lb.currentServerIndex]
	}
	// Increment the current server index
	lb.currentServerIndex = (lb.currentServerIndex + 1) % len(lb.servers)
	// Return the server
	return server
}
