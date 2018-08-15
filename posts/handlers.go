package posts

import (
	"fmt"
	"net/http"
)

// HandleRoute handles route
func HandleRoute(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "rgrhr")
	fmt.Println(request.URL)
}
