package posts

import (
	"fmt"
	"log"
	"net/http"
)

var portNumber = 3000

func getNewMux() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/posts", HandleRoute)
	return mux
}

// StartServer starts the server
func StartServer() {
	mux := getNewMux()

	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s%d", ":", portNumber), mux))
}
