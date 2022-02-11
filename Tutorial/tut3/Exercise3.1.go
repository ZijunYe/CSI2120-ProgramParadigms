package main

// Send the sequence 2, 3, 4, ... to channel 'ch'.
func Generate(ch chan<- int) {
	for i := 2; ; i++ {
		ch <- i // Send i to channel 'ch'.
	}
}

// Copy values from channel 'in' to channel 'out' removing those divisible by 'p'
func Filter(in <-chan int, out chan<- int, p int) {
	for {
		i := <-in // Get the values from 'in'.
		if i%p != 0 {
			out <- i // Send i to channel 'ch'.
		}
	}
}
func main() {
	ch := make(chan int) // Create a new channel.
	go Generate(ch)      // We run the Generate goroutine
	for i := 0; i < 20; i++ {
		p := <-ch
		print(p, "\n")
		ch1 := make(chan int)
		go Filter(ch, ch1, p) // all filters are chained
		ch = ch1
	}
}
