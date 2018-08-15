package posts

import (
	"log"
	"net/http"
)

var PORT_NUMBER = 3000

func getNewMux() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", HandleRoute)
	return mux
}

// StartServer starts the server
func StartServer() {
	mux := getNewMux()

	log.Fatal(http.ListenAndServe(":{{PORT_NUMBER}}", mux))
}
