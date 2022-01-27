package main

import "fmt"

func main() {
	x, y := 32, 64
	func(a int) {
		x = 128 //x is defined outside, so it will refer the global
		fmt.Println("a = ", a)
		fmt.Println("x - ", x)
		fmt.Println("y = ", y)

	}(256) //define outside of function call for the function

	fmt.Println("x = ", x)
	fmt.Println("y = ", y)
}

//lambda function will be called directly or assigned to a variable (direct call)
