package posts

import (
	"fmt"
	"net/http"
)

// HandleRoute handles route
func HandleRoute(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "rgrhr")

	// thing := strings.TrimPrefix(request.URL.Path, "/posts/")

	id, err := request.URL.Query()["key"]

	if !err || len(id[0]) < 1 {
		fmt.Fprintf(writer, "Can't find key")
		return
	}

		fmt.Fprintf(writer, id[0])
}
