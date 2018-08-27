package server

import (
	"fmt"
	"incognitorecord/route"
	"log"
	"net/http"
)

var portNumber = 3000

// StartServer starts the server.
// This server will mostly handle the CRUD operations of blog posts.
func StartServer() {
	mux := http.NewServeMux()
	route.HandleRoutes(mux)

	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s%d", ":", portNumber), mux))
}
