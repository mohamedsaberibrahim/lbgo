package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/mohamedsaberibrahim/lbgo/internals"
)

const (
	// The port to listen on
	LOAD_BALANCER_PORT = ":8000"
	SERVER_A_PORT      = ":8080"
	SERVER_B_PORT      = ":8081"
	SERVER_C_PORT      = ":8082"
	LOCAL_HOST         = "http://localhost"
)

func serverAHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Server A is serving this request")
}

func serverBHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Server B is serving this request")
}

func serverCHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Server C is serving this request")
}

func runningServerAsRoutine() {
	serverMuxA := http.NewServeMux()
	serverMuxA.HandleFunc("/", serverAHandler)

	serverMuxB := http.NewServeMux()
	serverMuxB.HandleFunc("/", serverBHandler)

	serverMuxC := http.NewServeMux()
	serverMuxC.HandleFunc("/", serverCHandler)

	go func() {
		log.Printf("Server A is listing for requests at %s%s\n", LOCAL_HOST, SERVER_A_PORT)
		log.Fatal(http.ListenAndServe(SERVER_A_PORT, serverMuxA))
	}()
	go func() {
		log.Printf("Server B is listing for requests at %s%s\n", LOCAL_HOST, SERVER_B_PORT)
		log.Fatal(http.ListenAndServe(SERVER_B_PORT, serverMuxB))
	}()
	go func() {
		log.Printf("Server C is listing for requests at %s%s\n", LOCAL_HOST, SERVER_C_PORT)
		log.Fatal(http.ListenAndServe(SERVER_C_PORT, serverMuxC))
	}()
}

func main() {
	// Running the servers as go routines
	runningServerAsRoutine()
	// Initialize the load balancer
	lb := internals.LoadBalancer{}
	lb.New(LOAD_BALANCER_PORT)
	// Initialize the servers
	serverA := internals.Server{}
	serverB := internals.Server{}
	serverC := internals.Server{}
	// Add the servers to the load balancer
	serverA.New("Server A", LOCAL_HOST+SERVER_A_PORT)
	serverB.New("Server B", LOCAL_HOST+SERVER_B_PORT)
	serverC.New("Server C", LOCAL_HOST+SERVER_C_PORT)

	lb.AddServer(&serverA)
	lb.AddServer(&serverB)
	lb.AddServer(&serverC)

	// Redirect the requests to the load balancer
	redirectRequest := func(w http.ResponseWriter, req *http.Request) {
		lb.ServeRequest(w, req)
	}

	http.HandleFunc("/", redirectRequest)
	log.Printf("Load balancer is listing for requests at %s%s\n", LOCAL_HOST, lb.GetPort())
	log.Fatal(http.ListenAndServe(lb.GetPort(), nil))
}
