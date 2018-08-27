package route

import (
	"log"
	"net/http"
)

type HandlerGroup []struct {
	Path    string
	Handler http.Handler
}

func recoverRoutes() {
	if r := recover(); r != nil {
		log.Println("Recovered from ", r)
	}
}

func (group HandlerGroup) Register(prefix string, mux *http.ServeMux) {
	for _, h := range group {
		defer recoverRoutes()

		mux.Handle(prefix+h.Path, h.Handler)
	}
}
