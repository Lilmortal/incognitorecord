package posts

import (
	"fmt"
	"log"
	"net/http"
)

var portNumber = 3000

// func getApiV1Mux() http.Handler {
// 	db, err := dynamo.New(config.Region)
// 	if err != nil {
// 		// TODO: How do most Golang applications handle errors
// 		panic("Failed to create session: " + err.Error())
// 	}

// 	// mux := http.NewServeMux()

// 	h := DynamoPostCreator{db}
// 	http.Handle("/posts/", h)

// 	return h
// 	// return mux
// }

// func getNewMux() *http.ServeMux {
// 	// mux := http.NewServeMux()

// 	http.Handle("/api/v1/", http.StripPrefix("/api/v1", getApiV1Mux()))
// 	return http.DefaultServeMux
// }

// StartServer starts the server
func StartServer() {

	hg := HandlerGroup{
		// 	{
		// 	"/users/", myUserHandler,
		// },
		{
			"/posts", DynamoPostCreator{},
		}}

	mux := http.NewServeMux()
	hg.Register("/api/v1", mux)

	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s%d", ":", portNumber), mux))
}

type HandlerGroup []struct {
	Path    string
	Handler http.Handler
}

func (group HandlerGroup) Register(prefix string, mux *http.ServeMux) {
	for _, h := range group {
		mux.Handle(prefix+h.Path, h.Handler)
	}
}
