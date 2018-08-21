package dynamo

import (
	"incognitorecord/config"
	"time"

	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// Client is an implementation of DatabaseClient for DynamoDB.
// TODO: Think of a better name
type Client struct {
	client *dynamodb.DynamoDB
}

// Post is a DynamoDB item
// TODO: Where to put this...
type Post struct {
	CreationDate time.Time `json:"creationDate"`
	Title        string    `json:"title"`
	Post         string    `json:"post"`
}

// New generates and returns a client that is connected to DynamoDB. This client will be used to handle database interactions.
func New(region string) (*Client, error) {
	session, err := createDynamoSession(region)
	if err != nil {
		return nil, err
	}
	client := createDynamoClient(session)

	return &Client{client: client}, err
}

// CreateTable creates a table in DynamoDB.
// TODO: Move this to config, so we can inject this; meaning we can reuse this in the future to create other tables.
func (client Client) CreateTable() (interface{}, error) {
	input := &dynamodb.CreateTableInput{
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("creationDate"),
				AttributeType: aws.String("S"),
			},
			{
				AttributeName: aws.String("title"),
				AttributeType: aws.String("S"),
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("title"),
				// Primary key
				KeyType: aws.String("HASH"),
			},
			{
				AttributeName: aws.String("creationDate"),
				// Sort key
				KeyType: aws.String("RANGE"),
			},
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(5),
			WriteCapacityUnits: aws.Int64(5),
		},
		TableName: aws.String("IncognitoRecord"),
	}

	res, err := client.client.CreateTable(input)

	return res, err
}

// CreatePost creates a post item in DynamoDB.
func (client Client) CreatePost(creationDate time.Time, title string, post string) error {
	newPost := Post{creationDate, title, post}

	av, err := dynamodbattribute.MarshalMap(newPost)

	if err != nil {
		return err
	}
	postInput := &dynamodb.PutItemInput{Item: av, TableName: aws.String("IncognitoRecord")}

	_, errItem := client.client.PutItem(postInput)

	return errItem
}

func createDynamoSession(region string) (*session.Session, error) {
	session, err := session.NewSession(&aws.Config{Region: aws.String(config.Region)})

	return session, err
}

func createDynamoClient(session *session.Session) *dynamodb.DynamoDB {
	return dynamodb.New(session)
}
