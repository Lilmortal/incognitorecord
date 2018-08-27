package server

import (
	"fmt"
	"incognitorecord/route"
	"log"
	"net/http"
)

var portNumber = 3000

// StartServer starts the server
func StartServer() {

	mux := http.NewServeMux()
	route.HandleRoutes(mux)

	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s%d", ":", portNumber), mux))
}
