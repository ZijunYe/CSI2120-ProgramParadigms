//Buffered Channels

package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func numbers(sz int) (res chan float64) {
	res = make(chan float64)
	go func() {
		defer close(res)
		num := 0.0
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Microsecond)
		for i := 0; i < sz; i++ {
			num += math.Sqrt(math.Abs(rand.Float64()))
		}
		num /= float64(sz)
		res <- num //only 1 value
		return
	}()
	return //return back the channel
}

func main() {
	var nGo int
	rand.Seed(42)                        //generate random number
	fmt.Print("Number of Go routines: ") //get value from user
	fmt.Scanf("%d \n", &nGo)             //number of go routine
	res := make([]chan float64, nGo)     //array of channel

	//each channel can receive nGo thread
	//res := make(chan float64, nGo)
	//when create an array of channel, it has to have capacity like buffer channel


	for i := 0; i < nGo; i++ { //filled all the channel
		res[i] = numbers(1000) //for each channel

	}

	threads := nGo
	for threads > 0 {
		for i := range res {
			select {
			case num, check := <-res[i]:
				if check {
					fmt.Printf("Result %d: %f\n", i, num)
				} else {
					threads -= 1
				}
			}

		}
	}
}
