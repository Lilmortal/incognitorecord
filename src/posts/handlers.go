package posts

import (
	"fmt"
	"incognitorecord/src/db"
	"log"
	"net/http"
	"time"

	"github.com/aws/aws-sdk-go/aws/awserr"
)

// PostV1Handler is a handler for posts in v1
type PostV1Handler struct {
	db db.DatabaseClient
}

func (handler PostV1Handler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodPost {
		err := handler.db.CreatePost(time.Now(), "Title", "Post")
		if err != nil {
			// TODO: Understand what this means...
			if awsErr, ok := err.(awserr.Error); ok {
				switch awsErr.Code() {
				case "ResourceNotFoundException":
					_, err = handler.db.CreateTable()
					if err != nil {
						if awsErr, ok := err.(awserr.Error); ok {
							log.Println("Failed to create table", awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
						}
					}
				default:
					log.Println("Failed to create post: ", awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
				}
			}
		}
	}

	fmt.Fprintf(writer, "Post created.")
	log.Println("Post created.")
}
