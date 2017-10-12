package greeting_test

import (
	"testing"

	"../greeting"
	"github.com/stretchr/testify/assert"
)

func TestSalutationRename(t *testing.T) {

	//arrange
	s := greeting.Salutation{"Joe", "Hello"}

	//act
	s.Rename("Fred")

	//assert
	assert.Equal(t, "Fred", s.Name, "because we called rename with Fred")
}

func TestRenameToFred(t *testing.T) {

	//arrange
	s := greeting.Salutation{"Joe", "Hello"}

	//act
	greeting.RenameToFred(&s)

	//assert
	assert.Equal(t, "Fred", s.Name)
}
