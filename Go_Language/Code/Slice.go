//How to create slice?
/*
	var slice []int = array[start:end] //end is exlusive
	slice := make([]int, 10,100)//size and capacity
*/

package main

import "fmt"

func mean(tab []int) (meanVal float64) {
	for _, value := range tab {
		meanVal += (float64)(value)
	}

	meanVal /= (float64)(len(tab))
	return
}

func main() {
	var table = [5]int{3, 4, 8, 9, 2}

	m := mean(table[:]) //all element
	fmt.Printf("result1 = %f\n", m)

	m2 := mean(table[2:]) //elements 2 to the end
	fmt.Printf("result2 = %f\n", m2)

	m3 := mean(table[1:3]) //2 elements from 1 up to 3,not include index of 3
	fmt.Printf("result3 = %f\n", m3)

}
