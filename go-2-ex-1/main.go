package main

import "fmt"

type FullName struct {
	FirstName    string
	LastName     string
}

// TODO: declare a structure for birth date
type BirthDate struct {
	DayOfBirth   int
	MonthOfBirth int
	YearOfBirth  int
}
type Profile struct {
	// TODO: embed full name and birth date information
	Name FullName
	BirthDate BirthDate
	NumberOfSiblings byte
	ZodiacSign       rune
}

func main() {
	var me = Profile{
		// TODO: set name and birth date information
		Name: FullName{
			FirstName: "Marco",
			LastName: "Wismer",
		},
		BirthDate: BirthDate{
			DayOfBirth: 27,
			MonthOfBirth: 7,
			YearOfBirth: 2007,
		},
		NumberOfSiblings: 1,   // TODO: adjust
		ZodiacSign:       ' ', // TODO: adjust
	}
	fmt.Println(me)

	fmt.Println("Siblings Before:", me.NumberOfSiblings)
	// TODO: imagine, you get a little brother or sister
	fmt.Println("Siblings After:", me.NumberOfSiblings)
}
