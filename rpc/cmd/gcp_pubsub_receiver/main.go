package main

import (
	"flag"
	"log"

	"github.com/cybertec-postgresql/pgwatch-contrib/rpc/sinks"
)

func main() {
	port := flag.String("port", "", "Port number for the server to listen on.")
	projectID := flag.String("project-id", "", "GCP Project Id.")
	flag.Parse()

	server, err := NewPubsubReceiver(*projectID)
	if err != nil {
		log.Fatal(err)
	}

	if err := sinks.ListenAndServe(server, *port); err != nil {
		log.Fatal(err)
	}
}