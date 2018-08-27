package logging

import (
	"log"

	"github.com/aws/aws-sdk-go/aws/awserr"
)

// AwsPrintln prints out the Aws error code and message, as well as an optional message.
func AwsPrintln(message string, awsErr awserr.Error) {
	log.Println(message)
	log.Println("Code: ", awsErr.Code())
	log.Println("Message: ", awsErr.Message())
}
