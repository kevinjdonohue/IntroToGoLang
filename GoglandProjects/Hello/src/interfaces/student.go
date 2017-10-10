package interfaces

type Student struct {
	FirstName string
	LastName string
}

func (student *Student) Rename(newFirstName string, newLastName string) {
	student.FirstName = newFirstName
	student.LastName = newLastName
}


