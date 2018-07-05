package aws

import (
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/farrout/goaway/aws/auth"
	"testing"
)

func TestListInstances(t *testing.T) {

sess := auth.CreateSession()

	// Create a EC2 client from just a session.
	svc := ec2.New(sess)


}
