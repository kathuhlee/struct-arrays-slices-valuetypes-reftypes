package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
	// contactInfo
	// the above declares the field name of contactInfo and it has the type contactInfo
}

func main() {
	// this is not recommended, i.e. if you change the order of fname and lname, itll be tedious to update
	// kathleen := person{"Kathleen", "Wong"}

	// kathleen := person{firstName: "Kathleen", lastName: "Wong"}
	// fmt.Println(kathleen)

	// var kathleen person
	// kathleen.firstName = "Kathleen"
	// kathleen.lastName = "Wong"

	// fmt.Println(kathleen)
	// fmt.Printf("%+v", kathleen)
	// %v makes it print out fields and value

	kathleen := person{
		firstName: "Kathleen",
		lastName:  "Wong",
		contact: contactInfo{
			email:   "kw@gmail.com",
			zipCode: 99999,
		},
		// contactInfo: contactInfo{
		// 	email:   "kw@gmail.com",
		// 	zipCode: 99999,
		// },
	}
	// &variableName & is an operator
	// look at this variable and give me the memory address of the value this variable is pointing at
	kathleenPointer := &kathleen
	kathleenPointer.updateName("Mochi")
	kathleen.print()
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
	// give me the value this memory address is pointing at

	// *person is a type description, it means we're working with a pointer to a person
	// updateName can only be called with the receiver of a pointer to a person
	// *pointerToPerson - an operator, it means we want to manipulate the value of the pointer we are referencing
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

// arrays
//  primitive data structure
//  cannot be resized
//  rarely used directly

//  slices
// 	can grow and shrink
// 	used 99% of the time for lists of elements
//  when we create a slice, go will automatically create an array and structure that records the length of the slice, the capacity of the slice and a reference to the underlying array

// value types - use pointers to change these things into a function
// 	int
// 	float
// 	string
// 	bool
// 	structs

// reference types - do not worry about pointers with these
// 	slices
// 	maps
// 	channels
// 	pointers
// 	functions

Map (Go) Hash (Ruby)  Object (Javascript)  Dict (Python)
all similar to one another across languages

a map in go carries key value pairs
all keys and all values in a map have to be of the same type
