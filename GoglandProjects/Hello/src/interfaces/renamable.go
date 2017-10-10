package interfaces

type Renamable interface {
	Rename(newFirstName string, newLastName string)
}

func UpdateNames(r Renamable, newFirstName string, newLastName string) {
	r.Rename(newFirstName, newLastName)
}