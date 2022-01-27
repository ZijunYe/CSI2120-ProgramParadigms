//Go Basic: Functions

/*
Write a function in Go that takes as input a float variable and returns two integer values.
One integer value which is the floor of the float value and the second integer value which is the ceiling of the float value. Print the result to console.
*/

package main

import (
	"fmt"
	"math"
)

func main() {
	x := 32.6
	f, c := floatVar(x)
	fmt.Printf("The floor number of %f is %f\n", x, f)
	fmt.Printf("The ceiling number of %f is %f\n", x, c)
}

func floatVar(num float64) (floor float64, ceiling float64) {
	floor = math.Floor(num)
	ceiling = math.Ceil(num)
	return
}

/*
func limits(x float64)(l,u int){
	l = int(x) //convert to the integer
	if x > 0{
		u = l + 1
	}else{
		u = l
		l = u - 1
	}

}

*/
