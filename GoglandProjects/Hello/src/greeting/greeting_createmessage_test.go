package greeting_test

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"greeting"
)

func TestCreateMessageShouldReturnHelloAndYoMessagesGivenJoeAndHello(t *testing.T) {
	//arrange
	var name string = "Joe"
	var messagePrefix string = "Hello"
	var expectedPrimaryMessage string = messagePrefix + " " + name
	var expectedAlternativeMessage string = "Yo " + " " + name

	//act

	actualPrimaryMessage, actualAlternativeMessage := greeting.CreateMessage(name, messagePrefix)

	//assert
	assert.Equal(t, expectedPrimaryMessage, actualPrimaryMessage, "they should be equal")
	assert.Equal(t, expectedAlternativeMessage, actualAlternativeMessage, "they should be equal")
}

func TestCreateMessageShouldReturnEmptyHelloAndYoMessagesGivenEmptyName(t *testing.T) {
	//arrange
	var name string = ""
	var messagePrefix string = "Hello"
	var expectedPrimaryMessage string = messagePrefix + " " + name
	var expectedAlternativeMessage string = "Yo " + " " + name

	//act

	actualPrimaryMessage, actualAlternativeMessage := greeting.CreateMessage(name, messagePrefix)

	//assert
	assert.Equal(t, expectedPrimaryMessage, actualPrimaryMessage, "they should be equal")
	assert.Equal(t, expectedAlternativeMessage, actualAlternativeMessage, "they should be equal")
}

func TestCreateMessageShouldReturnEmptyMessagesGivenEmptyNameAndMessagePrefix(t *testing.T) {
	//arrange
	var name string = ""
	var messagePrefix string = ""
	var expectedPrimaryMessage string = messagePrefix + " " + name
	var expectedAlternativeMessage string = "Yo " + " " + name

	//act

	actualPrimaryMessage, actualAlternativeMessage := greeting.CreateMessage(name, messagePrefix)

	//assert
	assert.Equal(t, expectedPrimaryMessage, actualPrimaryMessage, "they should be equal")
	assert.Equal(t, expectedAlternativeMessage, actualAlternativeMessage, "they should be equal")
}