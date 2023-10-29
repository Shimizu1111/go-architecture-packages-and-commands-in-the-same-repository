package auth

import "testing"

func TestAuthenticate(t *testing.T) {
	if !Authenticate("user", "pass") {
		t.Error("Authentication failed")
	}
}
