package main

import (
	"fmt"
	"time"
)

func tr(c chan int, w time.Duration) {
	{
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("Channel data sending stopped")
			}
		}() //lamdba function

	}
	i := 0
	for { //FOREVER LOOP
		time.Sleep(w)
		c <- i
		i++
	}
}
func main() {
	m1 := make(chan int)
	m2 := make(chan int)

	//two go routine run simotanously
	go tr(m1, 1*time.Second)
	go tr(m2, 1*time.Second)
	a := time.After(5 * time.Second) //wait for the duration time
P1:
	for {
		select {
		//alternating from m1 and m2
		case msg1 := <-m1:
			fmt.Println("received from m1 message", msg1)
		case msg2 := <-m2:
			fmt.Println("received from m2 message", msg2)
		case <-a: //run after 5 second
			fmt.Println("Time out")
			break P1
		}
		time.Sleep(time.Second)
	}
	close(m1)
	close(m2)
	time.Sleep(3 * time.Second)
}
