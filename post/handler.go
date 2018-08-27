package post

import (
	"fmt"
	"incognitorecord/db"
	"incognitorecord/logging"
	"log"
	"net/http"
	"time"

	"github.com/aws/aws-sdk-go/aws/awserr"
)

// Handler for posts related services.
type Handler struct {
	DB db.PostClient
}

// TODO: After creating table, it should create post after that
func (handler Handler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {
		createPost(handler.DB)
	}

	fmt.Fprintf(writer, "Post created.")
	log.Println("Post created.")
}

func createPost(db db.PostClient) {
	err := db.CreatePost(time.Now(), "Title", "Post")
	if err != nil {
		// TODO: Understand what this means...
		if awsErr, ok := err.(awserr.Error); ok {
			switch awsErr.Code() {
			case "ResourceNotFoundException":
				err = db.CreateTable()
				if err != nil {
					if awsErr, ok := err.(awserr.Error); ok {
						logging.AwsPrintln("Failed to create table.", awsErr)
					}
				}
			default:
				logging.AwsPrintln("Failed to create post.", awsErr)
			}
		}
	}
}
