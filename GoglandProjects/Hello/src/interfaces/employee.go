package interfaces

type Employee struct {
	FirstName string
	LastName string
}

func (employee *Employee) Rename(newFirstName string, newLastName string) {
	employee.FirstName = newFirstName
	employee.LastName = newLastName
}
