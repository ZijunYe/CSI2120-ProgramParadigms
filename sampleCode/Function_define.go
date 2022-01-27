package main

import "fmt"

const pi = 3.1416

var x int = 5 //global variable

func main() {
	var ( //grouping or factoring the keyword
		a float64 = 8.8
		b float64
	)

	b = foo(a)
	fmt.Printf("Result: %f", b)
}

func foo(z float64) float64 {
	u := 3.3 //initializing declaration
	return u * z
}

func plusminus(a int, b int) (sum int, difference int) {
	sum = a + b
	difference = a - b
	return // knows what to return based on the same variable name
}
