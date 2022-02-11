//How to construct error message with error interface, receiver

package main

import (
	"fmt"
	"math"
)

//the interface of two type of error
type NegError struct {
	num float64 // negative number
}
type Div0Error struct {
}

func (e *NegError) Error() string { //receiver is type of negerror, return back string
	return fmt.Sprintf("NegError: Can't divided the negative number %f!", e.num)
}

func (e *Div0Error) Error() string {
	return fmt.Sprintf("Div0Error: Can't divided number by 0! ")
}

func rootDivN(num float64, n int) (res float64, err error) {
	if num < 0.0 {
		err = &NegError{num}
	}
	if n == 0 {
		err = &Div0Error{}
	}
	res = math.Sqrt(num) / float64(n)
	return
}

func main() {
	divs := []int{2, 10, 3, 0}
	nums := []float64{511.8, 0.65, -3.0, 2.123}

	for i, num := range nums {
		fmt.Printf("%d) sqrt(%f)/%d = ", i, num, divs[i])
		res, err := rootDivN(num, divs[i])
		if err == nil {
			fmt.Printf("%f\n", res)
		} else {
			fmt.Println(err)
		}
	}
}
