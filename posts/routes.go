package posts

import (
	"log"
	"net/http"
)

func handleFunc() {
	http.HandleFunc("/", HandleRoute)

	log.Fatal(http.ListenAndServe(":3000", nil))
}

// StartServer starts the server
func StartServer() {
	handleFunc()
}
