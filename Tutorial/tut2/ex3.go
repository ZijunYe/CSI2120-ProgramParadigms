package main

import "fmt"

type Person struct {
	last_name  string
	first_name string
}

func (p *Person) Hello() {
	fmt.Printf("Hello, I am %s %s\n", p.first_name, p.last_name)
}

type Employee struct {
	Person
	salary float32
}

func main() {
	p1 := Person{"Cesar", "Jules"}
	p1.Hello()
	p2 := &Employee{Person{"Gates", "Bill"}, 15.25}
	p2.Hello()
}
