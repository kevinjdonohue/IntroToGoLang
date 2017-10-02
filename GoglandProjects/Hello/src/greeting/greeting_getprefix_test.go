package greeting_test

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"greeting"
)

func TestGetPrefixShouldReturnMrGivenBob(t *testing.T) {
	//arrange
	var name string = "Bob"
	var expectedPrefix string = "Mr "

	// act & assert
	GetPrefixShouldReturnCorrectPrefixForAGivenName(t, name, expectedPrefix)
}

func TestGetPrefixShouldReturnDrGivenJoe(t *testing.T) {
	//arrange
	var name string = "Joe"
	var expectedPrefix string = "Dr "

	// act & assert
	GetPrefixShouldReturnCorrectPrefixForAGivenName(t, name, expectedPrefix)
}

func TestGetPrefixShouldReturnDrGivenAmy(t *testing.T) {
	//arrange
	var name string = "Amy"
	var expectedPrefix string = "Dr "

	// act & assert
	GetPrefixShouldReturnCorrectPrefixForAGivenName(t, name, expectedPrefix)
}

func TestGetPrefixShouldReturnMrsGivenMary(t *testing.T) {
	//arrange
	var name string = "Mary"
	var expectedPrefix string = "Mrs "

	// act & assert
	GetPrefixShouldReturnCorrectPrefixForAGivenName(t, name, expectedPrefix)
}

func TestGetPrefixShouldReturnDudeGivenKevin(t *testing.T) {
	//arrange
	var name string = "Kevin"
	var expectedPrefix string = "Dude "

	//act & assert
	GetPrefixShouldReturnCorrectPrefixForAGivenName(t, name, expectedPrefix)
}

func TestGetPrefixShouldReturnDudeGivenEmptyString(t *testing.T) {
	//arrange
	var name string = ""
	var expectedPrefix string = "Dude "

	//act & assert
	GetPrefixShouldReturnCorrectPrefixForAGivenName(t, name, expectedPrefix)
}

func GetPrefixShouldReturnCorrectPrefixForAGivenName(t *testing.T, name string, expectedPrefix string) {
	//act
	var actualPrefix string = greeting.GetPrefix(name)

	//assert
	assert.Equal(t, expectedPrefix, actualPrefix, "they should be equal")
}
