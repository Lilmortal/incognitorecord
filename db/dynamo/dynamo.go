package dynamo

import (
	"incognitorecord/config"
	"time"

	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// PostClient is an implementation of a struct handling all posts related events for DynamoDB.
// TODO: Think of a better name
type PostClient struct {
	DB *dynamodb.DynamoDB
}

// Post is a DynamoDB item
type Post struct {
	CreationDate time.Time `json:"creationDate"`
	Title        string    `json:"title"`
	Post         string    `json:"post"`
}

// New generates and returns a client that is connected to DynamoDB. This client will be used to handle database interactions.
func New(region string) (*PostClient, error) {
	session, err := createDynamoSession(region)
	if err != nil {
		return nil, err
	}
	client := createDynamoClient(session)

	return &PostClient{client}, err
}

// CreateTable creates a table in DynamoDB.
// TODO: Move this to config, so we can inject this; meaning we can reuse this in the future to create other tables.
func (client PostClient) CreateTable() error {
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

	_, err := client.DB.CreateTable(input)

	return err
}

// CreatePost creates a post item in DynamoDB.
func (client PostClient) CreatePost(creationDate time.Time, title string, post string) error {
	newPost := Post{creationDate, title, post}

	av, err := dynamodbattribute.MarshalMap(newPost)

	if err != nil {
		return err
	}
	postInput := &dynamodb.PutItemInput{Item: av, TableName: aws.String("IncognitoRecord")}

	_, errItem := client.DB.PutItem(postInput)

	return errItem
}

// GetPost gets a post in DynamoDB.
func (client PostClient) GetPost(title string) error {
	return nil
}

// DeletePost deletes a post in DynamoDB.
func (client PostClient) DeletePost(title string) error {
	return nil
}

func createDynamoSession(region string) (*session.Session, error) {
	session, err := session.NewSession(&aws.Config{Region: aws.String(config.Region)})

	return session, err
}

func createDynamoClient(session *session.Session) *dynamodb.DynamoDB {
	return dynamodb.New(session)
}
