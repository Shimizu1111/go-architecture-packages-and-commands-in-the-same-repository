package validation

import "testing"

func TestValidateInput(t *testing.T) {
	_, err := ValidateInput("10")
	if err != nil {
		t.Errorf("expected nil error, got %v", err)
	}
}
