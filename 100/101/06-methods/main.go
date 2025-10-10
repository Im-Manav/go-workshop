package main

import (
	"fmt"
)

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func (p Person) String() string {
	return fmt.Sprintf("FirstName: %s, LastName: %s, Age: %d", p.FirstName, p.LastName, p.Age)
}

func (p Person) IsYoung() bool {
	return p.Age < 30
}

type IsYounger interface {
	IsYoung() bool
}

func main() {
	thingsThatCanBeYoung := []IsYounger{
		Person{"John", "Doe", 25},
		Person{"Jane", "Smith", 35},
		Person{"Alice", "Johnson", 28},
		Person{"Bob", "Brown", 40},
	}

	// The _ underscore is used to ignore the first return value of a range.
	// The first value is the index of the item in the slice.
	for _, pyt := range thingsThatCanBeYoung {
		if !pyt.IsYoung() {
			continue
		}
		fmt.Println(pyt)
	}
}
