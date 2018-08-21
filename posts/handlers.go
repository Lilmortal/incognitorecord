package posts

import (
	"fmt"
	"incognitorecord/db"
	"log"
	"net/http"
	"time"

	"github.com/aws/aws-sdk-go/aws/awserr"
)

// HandlePostsV1 is a handler for posts in v1
// TODO: Think of a good name
type HandlePostsV1 struct {
	db  db.DatabaseClient
	mux http.Handler
}

// TODO: After creating table, it should create post after that

func (handler HandlePostsV1) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	handler.mux.ServeHTTP(writer, request)

	log.Println("API")

}

type DynamoPostCreator struct {
	db1 db.DatabaseClient
	db  db.PostCreator
}

func (handler DynamoPostCreator) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodPost {
		err := handler.db.CreatePost(time.Now(), "Title", "Post")
		if err != nil {
			// TODO: Understand what this means...
			if awsErr, ok := err.(awserr.Error); ok {
				switch awsErr.Code() {
				case "ResourceNotFoundException":
					_, err = handler.db1.CreateTable()
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
