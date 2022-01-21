package main

import "fmt"

func main() {
	var ptr *int //define a pointer that points to the int value
	var i int
	i = 7

	ptr = &i //ptr have address of of i

	*ptr = i + 5 //change the value inside address of i

	fmt.Printf("result = %d\n", i) //%d means base 10
}
