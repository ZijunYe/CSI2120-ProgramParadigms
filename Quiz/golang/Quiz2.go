/*synchronizing using boolean channel*/

package main

import (
	"fmt"
)

func main() {
	ch := make(chan bool, 1)
	x := []int{3, 1, 4, 1, 5, 9, 2, 6}
	var y [8]int

	go calcul2(x[:4], y[:4], ch)
	<-ch
	go calcul2(x[4:], y[4:], ch)
	<-ch
	fmt.Println(y)
}
func calcul2(in []int, out []int, ch chan bool) {
	for i, v := range in {
		out[i] = 2*v*v*v + v*v
	}
	ch <- true

}

/*package main

import (
	"fmt"
)

func main() {
	//var wg sync.WaitGroup
	x := []int{3, 1, 4, 1, 5, 9, 2, 6}
	var y [8]int

	//wg.Add(2) // number of goroutines to synchronize
	// parallel loop in 2 slices
	//go calcul2(x[:4], y[:4], &wg)
	//go calcul2(x[4:], y[4:], &wg)
	done := make(chan bool, 1)
	go calcul2(x[:4], y[:4], done)
	go calcul2(x[4:], y[4:], done)
	//wg.Wait() // waiting for the 2 goroutines
	<-done
	fmt.Println(y)
}
func calcul2(in []int, out []int, done chan bool) {
	for i, v := range in {
		out[i] = 2*v*v*v + v*v
	}
	//wg.Done() // signals that the goroutine is finished
	done <- true
}*/
