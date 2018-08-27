package route

import (
	"incognitorecord/config"
	"incognitorecord/db/dynamo"
	"incognitorecord/post"
	"net/http"
)

// HandleRoutes will handle which routes will link to which handlers.
func HandleRoutes(mux *http.ServeMux) {
	db, err := dynamo.New(config.Region)
	if err != nil {
		panic("Failed to create session: " + err.Error())
	}

	v1HandlerGroup := HandlerGroup{
		{
			"/posts", post.Handler{DB: db},
		}}

	v1HandlerGroup.Register("/api/v1", mux)

}
