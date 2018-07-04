package auth

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"log"
)

func CreateSession() *session.Session {
	log.Println("Creating new AWS Session...")
	sess := session.Must(session.NewSession())

	dynamodb.New(sess, aws.NewConfig().WithLogLevel(aws.LogDebugWithHTTPBody))

	creds, errCreds := sess.Config.Credentials.Get()

	if errCreds != nil {
		log.Printf("Failed to find credentials")
	}

	log.Printf("Session created with credentials (%s)", creds.ProviderName)

	return sess
}
