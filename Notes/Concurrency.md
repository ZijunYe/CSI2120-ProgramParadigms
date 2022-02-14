## Concurrency and Parallelism 

- **Concurrent programming** 
  -   an application is a process running on a machine(a process is an independently executing entity that runs in its own address space) 
  -   a process is composed of one or more operating system threads(light weighted process) --> threads are concurrently executing entities that share the same address space  
- **Execution thread** 
  -   an execution thread is a sequence of executable statements and it may or may not interact with other threads(threads often share some varibles with other threads) 
  -   a thread can be blocked; access a shared variable or resource used currently by another thread. waiting for the result or completion of another thread 
  -   an application program often be divided into multiple processes and potentially many threads 

| Parallel programming       | Concurrent programming  |
| -------------------------- | ------------------------------------------ |
| process running on different processor or cores simultaneously  | concurrent means expressing program structure such that is organized into independently executing actions|
| /                          |it also target a single processor, the process run in turn over time according to a schedule|


- **Concurrent Programming language must support**
  1. creation and execution of processes and threads 
  2. data communication between threads and processes(may use mechanism using inter-process communication defined by the operating system) 
  3. synchronization operations (there are two ways to do synchronization
          - cooperative synchronization
              - a process waits for the execution of another before continuting its own execution(like conversation) 
          - Competitive synchronization ?? 
              -  multiple processes use the same resource with some form of locking mechanism for mutual exclusion(non-deterministic concurrency) 

- **level of concurrency** ?? 
   1. at the statement level - set of statement are executed concurrently while the main processs suspends its execution. all threads share the same data. 
   2. at the sub-program level - new process is created for running subroutine. once the new process starts, the calling process continues its execution.requires a synchronization mechanism 
   3. at the object level - each objet instance of a class competes for resources and method on different object run concurrently. class variable ar not shared 
   4. at the program level- parent process run one or more child processes. child process ID must known by parent. data may be shared 

- **type of concurrency**
  - physical
        - multile processes/ threads share different processor or cores 
  - logical 
        -   multiple processes/ threads hsare execution time on a single processor 
  - distributed
        -   processes/threads of an application share several machines over a network 


- **concurrency in go** 
  -  two machinsm are provided : non-deterministic concurrency and deterministic concurrency (gorountine and channel) 
  -  CSP(communication sequential processes) 
          -   based on the idea that avoding data sharing will avoid the biggest problem 
          -   threads in CSP do not communicate by sharing data
          -   timing of threads is based on messaging between threads 
      
## Goroutines 
- run concurrently, run in several threads 
- run in the same address space 
- goroutine exits when its code is complete or it is not complete, but main complete
- gorountine is invoked using the keyword ```go```
```go 
import "runtime"
func main(){
  //change max number of OS threads to 3 
  runtime.GOMAXPROCS(3) //how many processor 
  
  go fool() //this is thread 
}
```

***EXAMPLE: CALLING goroutine***
```go
 func main(){
     runtime.GOMAXPROCS(3)
     
    sTime := time.Now(); // time it
    
    fmt.Println("Starting")
    
    go letters() // goroutine A
    go numbers() // goroutine B
    
    fmt.Println("Waiting …")
    
    time.Sleep(2*time.Second) // for main thread to sleep, them alternating executing threadA and threadB 
    fmt.Println("\nDone\n")
    eTime := time.Now();
    fmt.Printf("Run time: %s", eTime.Sub(sTime))
}

func numbers() {
    for number := 1; number < 27; number++ {
    // pause before every print
    time.Sleep(10*time.Millisecond)
    fmt.Printf("%d ", number)
    }
}
func letters() {
    for char := 'a'; char < 'a'+26; char++ {
    time.Sleep(10)
    fmt.Printf("%c ", char)
    }
}
```

-  goroutine are design to effectively communicate by message passing for 
  -  exchange of information 
  -  synchronizing execution  


| Deterministic       | Non-Deterministic  |
| ------------------- | ------------------ |
| Channel             | waitGroup,semaphore,Mutex and Lock| 

## Channels 

**How to Declare channels? Why need to declaration?** 
- declaring channel only creates a reference; it use to allocate space for the channel
- only one goroutine has access to a channel 
- it will have waiting time; reading and writing in a channel are blocking statement 
```
var ch chan string 
ch = make(chan string) //declare the channel for string 

ch:= make(chan string)//shorter with initializer string 
```
- by default channel has capacity of 1 element 
- to send  or receive, using arrow operator 
```
ch <- str //send 
str = <- ch //receive 
``` 

**Want Channel to receive more than 1 Element?**
- it is called *Buffered channels* 
- a buffered channel will not block a sender until the buffer is full 
```
ch = make(chan string,2) //allows the insertion of 2 
```

**Example of channel**
```go
func main(){
	ch:= make(chan string) //create channel 
	go sendString(ch)
	go receiveString(ch)
	time.Sleep(1*time.Second)
}

func sendString(ch chan string){
	ch <- "Ottawa" //after channel receive Ottawa, then it will block until receiver receive the info from the channel 
	ch <- "Toronto"
	ch <- "Gatineau"
	ch <- "Casselman"
}

func recieveString(ch chan string){
	var str string
	for{//forever for loop
		str = <- ch //read channel, empty the channel 
		fmt.Printf("%s",str)
	}
}

```

**Range loop applied to a channel**
- we can loop over the elements in channel
```GO
func receiveString (ch chan string){
	go func()
	for str := range ch{
		fmt.printf("%s",str)
	}
}
```


**How to closing channel**
- Explicitly closed 
```go 
func sendString(strArr []string) chan string{
	ch:= make(chan string)
	go func(){//start a lambda in a go routine 
		for _,s := range starArr{
			ch <- s //send a string array to channel as whole 
		}
		close(ch) //explicit close channel 

	}()

	return ch
}
```
- How to test whether the channel is closed or not 
```go 
for {
	str,ok := <- ch //ok is represent boolean to check whether the channel closed or not 
	if !ok {
		break; 
	}
	fmt.Printf("%s",str)
}
```
- for multiple channel 
```go 
Forever: 
for {
	select{ //similar to switch 
		case str,ok := <- ch1: //If ch1 is blocked, then move on to next one 
			if !ok{
				break Forever 
			}
			fmt.Printf("%s \n", str)
		case str, ok := <- ch2:
			if !ok{
				break Forever
			}
			fmy.Print("%s \n", str)
	}
}
```

**Sync main with channel**
```go 
package main
import{
	"fmt"
	"time"
}

func worker (done chan bool){
	fmt.Print("Starting...")
	time.Sleep(time.Second)
	fmt.Print("end")
	done <- true //give the value to the channel
}


func main(){
	done := make(chan bool,1)

	go worker(done)

	<- done //synchronization point, here we don't need variable to be as receiver 
}
```

**Timer**
- represent a single event in the future 
- tell the timer how long you want to wait
```go
package main

import (
    "fmt"
    "time"
)

func main() {

    timer1 := time.NewTimer(2 * time.Second)

    <-timer1.C //blocks on the timer's channel C until it sends value 
    fmt.Println("Timer 1 fired")

    timer2 := time.NewTimer(time.Second)
    go func() { //threads 
        <-timer2.C
        fmt.Println("Timer 2 fired")
    }()
    stop2 := timer2.Stop() //check timer is stopped or not 
    if stop2 {
        fmt.Println("Timer 2 stopped")
    }

    time.Sleep(2 * time.Second)
}
```

## Semaphore 
- general mechanism that synchronization and resource sharing between process 
- it has a gobal channel 
- type of semaphore 
1. WaitGroup 
2. using channel 
3. Mutex and lock 

1. **WaitGroup**  
- wait for multiple goroutines to finish 
- keyword ```sync.WaitGroup```

	- ```WaitGroup.Add(int)``` counting up the # of things wait for
	- ```WaitGroup.Done()``` one less thing to wait for 
	- ```WaitGroup.Wait()``` wait until all things done 

```go 
func main(){
	var wg sync.WaitGroup
	x:=[]int{3,1,4,1,5,9,2,6}
	var y [8]int 

	wg.Add(2) 
	go calcul2(x[:4], y[:4],&wg) //first half 
	go calcul2(x[4:],y[4:],&wg)//second half 

	wg.Wait() // wait for the 2 goroutine 
	fmt.Println(y)
}


func calcul2(int []int, out []int, wg *sync.WaitGroup){
	for i,v:= range in{
		out[i] = 2*v*v*v + v*v
	}

	wg.Done() //sognals that the goroutine is over, decrease the waiting number 
}
```

2. **channel** version
```go 
// semaphore_channel
package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Print("starting...")
	time.Sleep(time.Second)
	fmt.Println("end")
	done <- true // signal for end
}

func main() {

	// synchronization channel
	done := make(chan bool, 1)
	// launching the goroutine
	go worker(done)

	//here we continue the main

	// synchronization point (rendezvous)
	<-done
}
```

3. **Mutex & Lock**
- can't have more than 1 item 
- threads that do not need communication 

```go
// semaphoreGo

package main

import (
	"fmt"
	"time"
)

type semaphore chan bool  //type of channel 

// to acquire a resource
func (s semaphore) acquire(n int) {

	for i := 0; i < n; i++ {
		s <- true
	}
}

// to free resources
func (s semaphore) free(n int) {

	for i := 0; i < n; i++ {
		<-s//update 
	}
}

// MUTEX
// to lock
func (s semaphore) Lock() {
	s.acquire(1)
}

// to unlock
func (s semaphore) Unlock() {
	s.free(1)
}

func main() {

	x := []int{3, 1, 4, 1, 5, 9, 2, 6}

	// creating a lock
	sem := make(semaphore, 1)

	go modify(x, 2, sem)
	go modify(x, 5, sem)

	time.Sleep(time.Second)
	fmt.Println(x)
}

// function modifying the values ​​of a table
func modify(table []int, inc int, sem semaphore) {

	// lock access to the array
	sem.Lock()
	for i, v := range table {
		table[i] = v + inc
	}
	// free access to the table

	sem.Unlock()

}
```


```go
//build in mutex
package main

import (
    "fmt"
    "sync"
    "time"
)

var sem sync.Mutex 
func main() {

	x := []int{3, 1, 4, 1, 5, 9, 2, 6}

	// creating a lock

	go modify(x, 2)
	go modify(x, 5)

	time.Sleep(time.Second)
	fmt.Println(x)
}

// function modifying the values ​​of a table
func modify(table []int, inc int) {

	// lock access to the array
	sem.Lock()
	for i, v := range table {
		table[i] = v + inc
	}
	// free access to the table

	sem.Unlock()

}

```
























