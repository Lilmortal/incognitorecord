package post

import (
	"fmt"
	"incognitorecord/db"
	"log"
	"net/http"
	"time"

	"github.com/aws/aws-sdk-go/aws/awserr"
)

type PostCreator struct {
	DB db.PostDB
}

// TODO: After creating table, it should create post after that
func (handler PostCreator) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {
		err := handler.DB.CreatePost(time.Now(), "Title", "Post")
		if err != nil {
			// TODO: Understand what this means...
			if awsErr, ok := err.(awserr.Error); ok {
				switch awsErr.Code() {
				case "ResourceNotFoundException":
					err = handler.DB.CreateTable()
					if err != nil {
						if awsErr, ok := err.(awserr.Error); ok {
							log.Println("Failed to create table.")
							log.Println("Code: ", awsErr.Code())
							log.Println("Message: ", awsErr.Message())
						}
					}
				default:
					log.Println("Failed to create post.")
					log.Println("Code: ", awsErr.Code())
					log.Println("Message: ", awsErr.Message())
				}
			}
		}
	}

	fmt.Fprintf(writer, "Post created.")
	log.Println("Post created.")
}
