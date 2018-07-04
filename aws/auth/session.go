package auth

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"log"
)

func CreateSession() {
	log.Println("Creating new AWS Session...")
	sess, err := session.NewSession()

	dynamodb.New(sess, aws.NewConfig().WithLogLevel(aws.LogDebugWithHTTPBody))

	if err != nil {
		log.Fatal("Failed to authenticate new Session with AWS")
	} else {
		log.Println("Successfully created AWS Session")
	}

	creds, errCreds := sess.Config.Credentials.Get()

	if errCreds != nil {
		log.Printf("Failed to find credentials")
	}

	log.Printf("Session created with credentials (%s)", creds.ProviderName)

}
