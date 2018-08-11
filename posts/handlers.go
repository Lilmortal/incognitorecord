package posts

import (
	"fmt"
	"net/http"
)

// HandleRoute handles route
func HandleRoute(writer http.ResponseWriter, request *http.Request) {
	// writer.Write([]byte("aaa"))
	fmt.Fprintf(writer, "rgrhr")
	fmt.Println("tesaaat")
}
