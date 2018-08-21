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
		// TODO: How do most Golang applications handle errors
		panic("Failed to create session: " + err.Error())
	}

	mux1 := http.NewServeMux()
	mux1.Handle("/posts", DynamoPostCreator{db})

	mux.Handle("/api/v1", HandlePostsV1{db, mux: mux1})
	return mux
}

// StartServer starts the server
func StartServer() {
	mux := getNewMux()

	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s%d", ":", portNumber), mux))
}
