package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

// listenAddr is the listening address for the web server we will run.
var listenAddr string

// NewServerCommand creates a cobra command. A command represents an action,
// the action in our case is to run our web server.
func NewServerCommand() *cobra.Command {
	// Before going forward, please take a look at this page:
	// https://godoc.org/github.com/spf13/cobra#Command, here you will find a
	// reference to how a cobra command is structured.
	cmd := &cobra.Command{
		Use:   "server",
		Short: "web server",
		Long: `The server is an example application to mimick the organization of the
Kubernetes API server.`,

		// Print usage when the command errors.
		SilenceUsage: false,
		Version:      "v1.0.0",

		// This is the function that will be executed when we actually run our CLI.
		// This function will return an error.
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Printf("version: %+v", cmd.Version)
			log.Printf("args: %#v", args)
			return Run(listenAddr)
		},
	}

	// Define flags.
	cmd.PersistentFlags().StringVarP(&listenAddr, "addr", "a", "0.0.0.0:8080", "server's address")

	return cmd
}

// homeHandler processes an incoming HTTP request by creating a response to
// greet the HTTP client with the host name of the computer running this app.
func homeHandler(w http.ResponseWriter, r *http.Request) {
	host, err := os.Hostname()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		errMsg := fmt.Sprintf("we saw an error: %v\n", err)
		fmt.Fprintf(w, errMsg)
		return
	}

	greeting := fmt.Sprintf("Greeting from %s!\n", host)
	fmt.Fprintf(w, greeting)
}

// Run registers http handlers to endpoints and starts the web server.
func Run(addr string) error {
	mux := http.NewServeMux()
	mux.HandleFunc("/", homeHandler)

	log.Printf("starting server at %s\n", addr)
	return http.ListenAndServe(addr, mux)
}
