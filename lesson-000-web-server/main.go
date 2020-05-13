package main

import (
	"fmt"
	"net/http"
	"os"
)

// listenAddr is the address where our web server will be listening for
// requests.
const listenAddr = "0.0.0.0:8080"

// homeHandler takes a response writer to build a response for the given
// request.
// This http handler will greet you with the hostname of the machine where this
// app is running.
func homeHandler(w http.ResponseWriter, r *http.Request) {
	// We begin by looking up the hostname.
	host, err := os.Hostname()
	if err != nil {
		// If we see an error then we return an http code 500 and we tell the
		// client what the error was.
		w.WriteHeader(http.StatusInternalServerError)
		errMsg := fmt.Sprintf("we saw an error: %v\n", err)
		fmt.Fprintf(w, errMsg)
		return
	}

	// Build a string with the hostname.
	greeting := fmt.Sprintf("Greeting from %s!\n", host)
	fmt.Fprintf(w, greeting)
}

func main() {
	// Create a request multiplexer. This will match an incoming request to a
	// route.
	mux := http.NewServeMux()

	// Register homeHandler with the router "/". This means that a request to
	// http://0.0.0.0:8080/ will be handled by the 'homeHandler' function.
	mux.HandleFunc("/", homeHandler)

	// Start the webserver by registering our multiplexer and listening to
	// 'listenAddr'.
	fmt.Printf("starting web server at %s\n", listenAddr)
	http.ListenAndServe(listenAddr, mux)
}
