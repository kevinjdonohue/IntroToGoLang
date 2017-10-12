package interfaces_test

import (
	"testing"
	"../interfaces"
	"github.com/stretchr/testify/assert"
)

func TestUpdateNamesForAnEmployee(t *testing.T) {
	//arrange
	employee := interfaces.Employee{"FirstName", "LastName"}

	//act
	interfaces.UpdateNames(&employee, "Kevin", "Donohue")

	//assert
	assert.Equal(t, "Kevin", employee.FirstName)
	assert.Equal(t, "Donohue", employee.LastName)
}

func TestUpdateNamesForAStudent(t *testing.T) {
	//arrange
	student := interfaces.Student{"First", "Last"}

	//act
	interfaces.UpdateNames(&student, "Kim", "Stantus")

	//assert
	assert.Equal(t, "Kim", student.FirstName)
	assert.Equal(t, "Stantus", student.LastName)
}
