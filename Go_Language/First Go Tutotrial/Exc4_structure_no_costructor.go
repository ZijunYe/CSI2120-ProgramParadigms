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

func NewSaiyan(name string, power int) *Saiyan {
	return &Saiyan{
		Name:  name,
		Power: power,
	}
}

//Or return by value:
/*
func NewSaiyan(name string, power int) Saiyan {
	return Saiyan{
		Name:  name,
		Power: power,
	}
}
*/

func main() {
	//Warning: there is a new but it just reserves memory (no constructor is run)
	goku1 := new(Saiyan)
	fmt.Println(goku1.Name)
	// same as
	goku2 := &Saiyan{}
	fmt.Println(goku2.Name)
}
