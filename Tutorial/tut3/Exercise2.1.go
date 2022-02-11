package main

import (
	"fmt"
	"time"
)

func tr(done chan bool, w time.Duration) <-chan int {
	intChan := make(chan int)
	go func() {
		defer close(intChan)
		defer fmt.Println("Channel data sending stopped")
		i := 0
		for {
			time.Sleep(w)
			select {
			case intChan <- i:
				i++
			case <-done: // we can always read from a closed channel return
			}
		}
	}()

	return intChan
}

func main() {
	done := make(chan bool)

	m1 := tr(done, 1*time.Second)
	m2 := tr(done, 3*time.Second)
	a := time.After(5 * time.Second)
P1:
	for {
		select {
		case msg1 := <-m1:
			fmt.Println("received from m1 message", msg1)
		case msg2 := <-m2:
			fmt.Println("received from m2 message", msg2)
		case <-a:
			fmt.Println("Time out")
			break P1
		}
		time.Sleep(time.Second)
	}
	close(done)
	time.Sleep(3 * time.Second)
}
