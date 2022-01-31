package main

import "fmt"

func main() {

	a := [5]int{12, -6, 9, 87, -11}
	b := [7]int{1, 2, 3, -4, -5, -6, 9}

	min, max := getMinMax(a[:])
	fmt.Println("a:")
	fmt.Println("Min= ", min)
	fmt.Println("Max= ", max)

	min, max = getMinMax(b[:])
	fmt.Println("b:")
	fmt.Println("Min= ", min)
	fmt.Println("Max= ", max)
}

//input array, output 2 value
func getMinMax(s []int) (min int, max int) {
	//looping through the slice
	//using for-range loop
	for i := range s {
		if i == 0 {
			min = s[i]
			max = s[i]
		} else { //i != 0
			if s[i] > max {
				max = s[i]
			}

			if s[i] < min {
				min = s[i]
			}
		}
	}
	return
}
