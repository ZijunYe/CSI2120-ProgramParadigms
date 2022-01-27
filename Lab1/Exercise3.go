//Go basic: structs and pointers

/*
A Go program is to read a person
(last and first name) from console (use an input array if using the sandbox) and assign an Id counting up.
The program must use the structure Person.
*/

package main

import "fmt"

type Person struct {
	lastName  string
	firstName string
	iD        int
}

func main() {
	nextId := 101
	for {
		var (
			p   Person
			err error
		)
		nextId, err = inPerson(&p, nextId)
		if err != nil {
			fmt.Println("Invalid entry ... exiting")
			break
		}
		printPerson(p)
	}
}

func inPerson(p *Person, id int) (next int, err error) {
	//get input from user and construct person object

	//fist name, last name
	fmt.Print("Please enter your last name: ")
	_, err = fmt.Scanf("%s", &p.lastName)
	if err != nil {
		return
	}

	fmt.Print("Please enter your firstname: ")
	_, err = fmt.Scanf("%s", &p.firstName)
	if err != nil {
		return
	}

	next = id + 1
	p.iD = id
	return
}

func printPerson(p Person) {
	fmt.Printf("%10.d\t%s\t%s\n", p.iD, p.firstName, p.lastName)
}
