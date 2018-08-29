package dynamo

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// Config is the minimum config needed to create a DynamoDB table.
type Config struct {
	AttributeDefinitions  []*dynamodb.AttributeDefinition
	KeySchema             []*dynamodb.KeySchemaElement
	ProvisionedThroughput *dynamodb.ProvisionedThroughput
	TableName             *string
}

// PostConfig defines the keys and attributes needed to create a post table.
var PostConfig = Config{postAttributeDefinitions, postKeySchema, postProvisionedThroughput, postTableName}

var postAttributeDefinitions = []*dynamodb.AttributeDefinition{
	{
		AttributeName: aws.String("creationDate"),
		AttributeType: aws.String("S"),
	},
	{
		AttributeName: aws.String("title"),
		AttributeType: aws.String("S"),
	}}

var postKeySchema = []*dynamodb.KeySchemaElement{
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
}

var postProvisionedThroughput = &dynamodb.ProvisionedThroughput{
	ReadCapacityUnits:  aws.Int64(5),
	WriteCapacityUnits: aws.Int64(5),
}

var postTableName = aws.String("IncognitoRecord")
