package greeting

import "testing"

func TestGenerateGreeting(t *testing.T) {
	expected := "Hello, World!"

	if got := GenerateGreeting(); got != expected {
		t.Errorf("Expected: %s, Got: %s", expected, got)
	}
}
