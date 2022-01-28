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

## Channels 
