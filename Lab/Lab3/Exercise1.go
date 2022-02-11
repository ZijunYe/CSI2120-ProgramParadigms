//channel, how to check channel's value and is or not empty channel (bool)
package main

import (
	"fmt"
)

func sendInt() (ch chan int) {
	defer fmt.Println("Sender ready!")
	ch = make(chan int)
	go func() {
		for i := 0; i < 20; i++ {
			ch <- i
		}
		close(ch) //channel contain list of integer array from o to 20
	}()
	return
}

//the parameter is receiver that receive the value from the passing channel
func filterInt(src <-chan int) (dst chan int) {
	defer fmt.Println("Filter ready!")
	dst = make(chan int) //making the integer channel
	go func() {
		defer close(dst)
		open := true
		for open {
			var i int
			i, open = <-src //TWO value get from the channel
			//1. i received the value from src
			//2. open get the message whether there have value in the channel or not

			if open {
				i *= 2
				dst <- i
			}
		}
	}()
	return
}

func main() {
	src := sendInt()      //return back the channel
	dst := filterInt(src) //passing the channel, return back the new channel
	open := true

	for open {
		i := 0
		i, open = <-dst //read from filterInt
		//if there have value in channel, it will be true
		//if there don't have value in channel, it will be false
		if open {
			fmt.Printf("%d ", i)
		}

	}
	fmt.Println()
}
