package internals

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
)

type Server struct {
	name    string
	address string
	proxy   *httputil.ReverseProxy
}

func (s *Server) New(name string, address string) {
	serverUrl, err := url.Parse(address)
	if err != nil {
		fmt.Println("Error while parsing server url", err)
	}

	s.name = name
	s.address = address
	s.proxy = httputil.NewSingleHostReverseProxy(serverUrl)
}

func (s *Server) GetName() string {
	return s.name
}

func (s *Server) GetAddress() string {
	return s.address
}

func (s *Server) CheckHealth() bool {
	return true
}

func (s *Server) ServeRequest(rw http.ResponseWriter, req *http.Request) bool {
	s.proxy.ServeHTTP(rw, req)
	return true
}
