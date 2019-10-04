package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"

	"fmt"
)

// Create struct to hold info about new item
type User struct {
	Name  string
	Age   int
	Email string
}

func main() {
	// Initialize a session that the SDK will use to load
	// credentials from the shared credentials file ~/.aws/credentials
	// and region from the shared configuration file ~/.aws/config.
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Create DynamoDB client
	dbClient := dynamodb.New(sess)

	tableName := "users"

	/*------------------------------------------------
		Create User
	------------------------------------------------*/

	/*------------------------------------------------
		Find User
	------------------------------------------------*/
	result, err := dbClient.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"Name": {
				S: aws.String("Ben"),
			},
		},
	})
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	user := User{}

	err = dynamodbattribute.UnmarshalMap(result.Item, &user)
	if err != nil {
		panic(fmt.Sprintf("Failed to unmarshal Record, %v", err))
	}

	if user.Name == "" {
		fmt.Println("Could not find ")
		return
	}

	fmt.Println("Found user:")
	fmt.Println("Name:  ", user.Name)
	fmt.Println("Age: ", user.Age)
	fmt.Println("Email:  ", user.Email)
}
