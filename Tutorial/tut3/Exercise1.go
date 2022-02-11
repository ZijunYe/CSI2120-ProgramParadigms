package main

import "fmt"

func numberGen(start, count int, out chan<- int) {
	for i := 0; i < count; i++ {
		out <- (start + i)
	}
	close(out)
}
func printNumbers(in <-chan int, done chan<- bool) {
	for {
		value, ok := <-in
		if !ok {
			done <- true
			break
		}
		fmt.Println(value)
	}
}

func main() {
	channel := make(chan int)
	done := make(chan bool)
	go numberGen(7, 7, channel)
	go printNumbers(channel, done)
	<-done
}
