package greeting_test

import (
	"testing"
	"greeting"
)

func TestGetPrefixShouldReturnMrGivenBob(t *testing.T) {
	//arrange
	var expectedPrefix string = "Mr "

	//act
	var actualPrefix string = greeting.GetPrefix("Bob")

	//assert
	if actualPrefix != expectedPrefix {
		t.Errorf("Expected: %s Got: %s", "'"+expectedPrefix+"'", "'"+actualPrefix+"'")
	}
}
