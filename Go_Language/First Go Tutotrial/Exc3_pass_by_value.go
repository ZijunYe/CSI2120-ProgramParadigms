/*
Some excersises and code snippets are taken from “The Little Go Book”

License
The Little Go Book is licensed under the Attribution-NonCommercial-ShareAlike 4.0 International license. You should not have paid for this book.
You are free to copy, distribute, modify or display the book. However, I ask that you always attribute the book to me,
Karl Seguin, and do not use it for commercial purposes.
You can see the full text of the license at:
https://creativecommons.org/licenses/by-nc-sa/4.0/

*/
package main

import "fmt"

type Saiyan struct {
	Name  string
	Power int
}

func main() {

	goku := Saiyan{
		Name:  "Goku",
		Power: 9000,
	}
	Super(goku)
	fmt.Printf("The power is %d\n", goku.Power)
}

func Super(s Saiyan) {
	s.Power += 10000
}
