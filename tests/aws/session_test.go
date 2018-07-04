package aws

import (
	"github.com/farrout/goaway/aws/auth"
	"testing"
)

func TestCredentials(t *testing.T) {
	sess := auth.CreateSession()

	creds, errCreds := sess.Config.Credentials.Get()

	if errCreds != nil {
		t.Errorf("Failed to find credentials")
	}

	t.Logf("Session created with credentials (%s)", creds.ProviderName)
}
