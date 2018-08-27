package route

import (
	"log"
	"net/http"
)

// HandlerGroup handles linking URL paths to a handler
type HandlerGroup []struct {
	Path    string
	Handler http.Handler
}

func recoverRoutes() {
	if r := recover(); r != nil {
		log.Println("Recovered from ", r)
	}
}

// Register function handles linking the prefix and the paths to a handler
func (group HandlerGroup) Register(prefix string, mux *http.ServeMux) {
	for _, h := range group {
		defer recoverRoutes()

		mux.Handle(prefix+h.Path, h.Handler)
	}
}
