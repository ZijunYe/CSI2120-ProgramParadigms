package main

import "fmt"

func mean(tab [5]int) (meanVal float64) {
	//for index,value:= range of collection
	for _, value := range tab { //what js _ means?
		meanVal += (float64)(value)
	}

	meanVal /= (float64)(len(tab))
	return
}

func main() {
	var table = [5]int{3, 4, 8, 9, 2}
	//or table := [...]int{3,4,8,9,2}
	m := mean(table) //pass by value
	fmt.Printf("result = %f\n", m)
}
