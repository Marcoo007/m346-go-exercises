package main

import "fmt"

type Student struct {
	FirstName string
	LastName  string
}

type Class struct {
	Students []Student
}

func main() {
	// Declare students
	student1 := Student{FirstName: "John", LastName: "Doe"}
	student2 := Student{FirstName: "Jane", LastName: "Smith"}

	// Declare classes
	class1 := Class{Students: []Student{student1, student2}}

	// Declare modules
	modules := map[string][]Class{
		"Math":    {class1},
		"Science": {class1},
	}

	// Output everything
	fmt.Println("Students:", class1.Students)
	fmt.Println("Modules:", modules)
}
