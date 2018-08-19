package posts

import (
	"fmt"
	"incognitorecord/config"
	"incognitorecord/db/dynamo"
	"log"
	"net/http"
)

var portNumber = 3000

func getNewMux() *http.ServeMux {
	mux := http.NewServeMux()

	db, err := dynamo.New(config.Region)
	if err != nil {
		panic("Failed to create session: " + err.Error())
	}

	mux.Handle("/api/v1/posts", HandlePostsV1{db})
	return mux
}

// StartServer starts the server
func StartServer() {
	mux := getNewMux()

	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s%d", ":", portNumber), mux))
}
