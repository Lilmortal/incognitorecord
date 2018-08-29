package post

import (
	"fmt"
	"incognitorecord/db"
	"incognitorecord/db/dynamo"
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

// NewHandler creates a new Handler struct
func NewHandler(DB db.PostClient) *Handler {
	return &Handler{DB}
}

func (handler Handler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {
		err := createPost(handler.DB)
		if err != nil {
			// TODO: Understand what this means...
			if awsErr, ok := err.(awserr.Error); ok {
				switch awsErr.Code() {
				case "ResourceNotFoundException":
					err := createTable(handler.DB, dynamo.PostConfig)
					if err != nil {
						if awsErr, ok := err.(awserr.Error); ok {
							logging.AwsPrintln("Failed to create table.", awsErr)
						}
					}
					postErr := createPost(handler.DB)
					if postErr != nil {
						if awsErr, ok := postErr.(awserr.Error); ok {
							logging.AwsPrintln("Failed to create a post even though a table has been created.", awsErr)
						}
					}
				default:
					logging.AwsPrintln("Failed to create post.", awsErr)
				}
			}
		}

	}

	fmt.Fprintf(writer, "Post created.")
	log.Println("Post created.")
}

func createPost(db db.PostWriter) error {
	err := db.CreatePost(time.Now(), "Title", "Post")
	return err
}

func createTable(db db.PostClient, config interface{}) error {
	err := db.CreateTable(config)
	return err
}
