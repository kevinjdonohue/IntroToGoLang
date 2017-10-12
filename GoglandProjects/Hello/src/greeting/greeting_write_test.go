package greeting_test

import (
	"fmt"
	"testing"

	"../greeting"
)

func TestSalutation_Write(t *testing.T) {

	//arrange
	salutation := greeting.Salutation{"Joe", "Hello"}

	//act
	fmt.Fprintf(&salutation, "The count is %d", 10)
}
