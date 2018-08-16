package posts

import (
	"fmt"
	"log"
	"net/http"
)

var portNumber = 3000

type handleV1 struct {
	mux *http.ServeMux
}

func (v handleV1) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("is it working")
	v.mux.HandleFunc("/posts", HandleRoute)
}

func getNewMux() *http.ServeMux {
	mux := http.NewServeMux()

	mux.Handle("/api/v1", handleV1{mux})
	return mux
}

// StartServer starts the server
func StartServer() {
	mux := getNewMux()

	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s%d", ":", portNumber), mux))
}
