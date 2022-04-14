//winter2020
package main

import "fmt"

func main() {
	numbers := []int{216, 218, 221, 260}
	ch := make(chan int)

	//your solution here
	//using lambda function here

	go func() {
		for _, v := range numbers {
			ch <- v
		}
		close(ch)
	}()

	for {
		if num, ok := <-ch; !ok {
			fmt.Println("Channel closed")
			break
		} else {
			fmt.Println(num)
		}
	}
}
