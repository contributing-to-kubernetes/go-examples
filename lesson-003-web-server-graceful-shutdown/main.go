package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
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
	time.Sleep(5 * time.Second)
}

func main() {
	done := make(chan bool, 1)
	sigint := make(chan os.Signal, 1)

	// Create a request multiplexer. This will match an incoming request to a
	// route.
	mux := http.NewServeMux()
	server := http.Server{Addr: listenAddr, Handler: mux}

	// Register homeHandler with the router "/". This means that a request to
	// http://0.0.0.0:8080/ will be handled by the 'homeHandler' function.
	mux.HandleFunc("/", homeHandler)

	signal.Notify(sigint,
		os.Interrupt,    // Sent from the terminal, e.g., Ctrl+C.
		syscall.SIGTERM, // Sent to signal pod to shutdown or something.
	)

	go func() {
		<-sigint
		log.Printf("Shutting down server")

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		server.SetKeepAlivesEnabled(false)
		if err := server.Shutdown(ctx); err != nil {
			log.Fatalf("Unable to do graceful shutdown: %v", err)
		}

		close(done)
	}()

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}

	<-done
	log.Printf("Gracefully stopped")
}
