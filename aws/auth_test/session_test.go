package auth_test

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"testing"
)

func TestCredentials(t *testing.T) {
	sess, err := session.NewSession()

	if err != nil {
		t.Errorf("Failed to authenticate")
	}

	creds, errCreds := sess.Config.Credentials.Get()

	if errCreds != nil {
		t.Errorf("Failed to find credentials")
	}

	t.Logf("Session created with credentials (%s)", creds.ProviderName)
}
