package main

import "fmt"

type Saiyan struct {
	Name  string
	Power int
}

func main() {

	//1. passing by value, the goku power will not be change
	/*goku := Saiyan{
		Name: "Goku",
		Power: 9000,
	}

	Super(goku) */

	//2. passing by pointer
	goku := &Saiyan{"Goku", 9000}
	Super_Pointer(goku)
	//what ever change made in function, power will be change

	fmt.Printf("The power is %d\n", goku.Power)
}

func Super(s Saiyan) {
	s.Power += 1000
}

func Super_Pointer(s *Saiyan) {
	s.Power += 1000
}
