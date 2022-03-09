package main

import (
	"fmt"
	"sync"
)

func main() {

	x := []int{3, 1, 4, 1, 5, 9, 2, 6}

	var y [8]int

	numThreads := len(x)
	var wg sync.WaitGroup
	wg.Add(numThreads)

	// parallel loop
	for i, v := range x {

		go func(i int, v int) {
			y[i] = calcul(i, v, &wg)
		}(i, v) // call to goroutine

	}

	// add synchronization
	wg.Wait()
	//time.Sleep(1 * time.Second)

	fmt.Println(y)

}

// you can add a channel to the list of parameters
func calcul(index int, v int, wg *sync.WaitGroup) int {
	i := 2*v*v*v + v*v
	wg.Done() //decrease 1 threads
	return i
}
