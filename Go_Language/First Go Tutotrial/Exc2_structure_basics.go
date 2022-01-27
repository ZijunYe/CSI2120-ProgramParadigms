/*
Some excersises and code snippets are taken from “The Little Go Book”

License
The Little Go Book is licensed under the Attribution-NonCommercial-ShareAlike 4.0 International license. You should not have paid for this book.
You are free to copy, distribute, modify or display the book. However, I ask that you always attribute the book to me,
Karl Seguin, and do not use it for commercial purposes.
You can see the full text of the license at:
https://creativecommons.org/licenses/by-nc-sa/4.0/

*/
//define a strucut for a rectangle and write a function to calculate the area and the length of the perimeter
package main

import "fmt"

type rect struct {
	width, height int
}

func (r *rect) area() int {
	return r.width * r.height
}

func (r *rect) perimeter() int {
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{width: 10, height: 5}
	fmt.Printf("area: %d\n", r.area())
	fmt.Printf("perim: %d", r.perimeter())
}
